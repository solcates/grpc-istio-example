.PHONY: all gen vendor

CI_COMMIT_REF_SLUG ?= dev
DOMAIN ?= example.com
HOST ?= example-$(CI_COMMIT_REF_SLUG)
GIT_COMMIT ?= $(shell git rev-list -1 HEAD)
TAGS ?= -ldflags "-X github.com/solcates/grpc-istio-example/pkg/greeter.Version=$(GIT_COMMIT)"
DEVTAG ?= latest

gen:
		@echo Generating Code
		@protoc \
		-I apis/ \
		-I $(GOPATH)/src/ \
		-I $(GOPATH)/src/github.com/googleapis/googleapis \
		--grpc-gateway_out=logtostderr=true:apis \
		--swagger_out=logtostderr=true:apis \
		--go_out=plugins=grpc:apis/. hello.proto

vendor:
		@echo Vendoring
		@go mod vendor

## Docker
docker-build: gen vendor
		@echo Building Docker Image
		$(eval DEVTAG := $(shell docker build --build-arg GIT_COMMIT=$(GIT_COMMIT) --quiet -t solcates/grpc-istio-example .))

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
		go run $(TAGS) ./cmd/greeter client --name LocalAlice --debug

## Running on Kubernetes
ARGS=   --set domain=$(DOMAIN) \
		--set host=$(HOST) \
		--set image.tag=$(DEVTAG)

deploy: docker
		@echo Deploying Chart
		@helm upgrade --install --namespace $(HOST) $(HOST) . \
		$(ARGS)
#		--recreate-pods

run-client-remote:
		@echo Connecting to Remote Client : $(HOST).$(DOMAIN)
		go run $(TAGS) ./cmd/greeter client --host $(HOST).$(DOMAIN) --name K8SAlice

