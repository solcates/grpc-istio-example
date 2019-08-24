.PHONY: all gen vendor e2e

CI_COMMIT_REF_SLUG ?= test1
BASE_NAME ?= example
DOMAIN ?= example.com
HOST ?= $(BASE_NAME)-$(CI_COMMIT_REF_SLUG)
RELEASE_NAME ?= $(BASE_NAME)-$(CI_COMMIT_REF_SLUG)
GIT_COMMIT ?= $(shell git rev-list -1 HEAD)
TAGS ?= -ldflags "-X github.com/solcates/grpc-istio-example/pkg/greeter.Version=$(GIT_COMMIT)"
DEVTAG ?= latest
hash = $(or $(word 2,$(subst :, ,$1)),$(value 2))

gen-grpc:
		@protoc \
		-I apis/ \
		-I $(GOPATH)/src/ \
		-I $(GOPATH)/src/github.com/googleapis/googleapis \
		--grpc-gateway_out=logtostderr=true:apis \
		--swagger_out=logtostderr=true:apis \
		--go_out=plugins=grpc:apis/. hello.proto
gen-rest:
		@swagger generate client -f apis/hello.swagger.json -c apis/rest/client -m apis/rest/models -q
gen-echo:
		@echo Generating

gen: gen-grpc gen-rest

vendor:
		@echo Vendoring
		@go mod vendor

## Docker
docker-build: gen vendor
		@echo Building Docker Image
		$(eval DEVTAG := $(shell docker build --build-arg GIT_COMMIT=$(GIT_COMMIT) --quiet -t solcates/grpc-istio-example . ))

docker-push:
		@echo Pushing Docker Image to Registry
		@docker push solcates/grpc-istio-example

docker: docker-build docker-push

## Building
build:
		@go build $(TAGS) --mod=vendor ./cmd/greeter

## Running local

run-server:
		go run $(TAGS) ./cmd/greeter server --debug
run-client:
		go run $(TAGS) ./cmd/greeter client --name LocalGRPC --debug
run-client-rest:
		go run $(TAGS) ./cmd/greeter client --name LocalREST --debug --rest --host 127.0.0.1

## Running on Kubernetes
ARGS=   --set domain=$(DOMAIN) \
		--set host=$(HOST) \
		--set image.hash=$(DEVTAG)

deploy:  deploy-only
deploy-only:
		@echo Deploying Chart
		@kubectl create namespace $(RELEASE_NAME) || true
		@kubectl label namespace $(RELEASE_NAME) istio-injection=enabled --overwrite || true
		@helm upgrade --install --namespace $(RELEASE_NAME) $(RELEASE_NAME) . \
		$(ARGS)
#		--recreate-pods

run-client-remote:
		@echo Connecting to Remote Client via gRPC: grpc://$(HOST).$(DOMAIN):443
		@go run $(TAGS) ./cmd/greeter client --host $(HOST).$(DOMAIN) --grpc_port 443 --name K8SAlice

run-client-remote-rest:
		@echo Connecting to Remote Client via REST: https://$(HOST).$(DOMAIN)
		@go run $(TAGS) ./cmd/greeter client --host $(HOST).$(DOMAIN) --rest_port 443 --name K8SAlice --rest

teardown:
		@helm delete --purge $(RELEASE_NAME)

## Testing


## e2e
e2e:
		@$(MAKE) run-client-remote-rest
		sleep 1
		@$(MAKE) run-client-remote


full-run: deploy
		@echo Wait a bit till the deployment has refreshed the pods...
		$(MAKE) e2e