.PHONY: all gen vendor

CI_COMMIT_REF_SLUG ?= dev
DOMAIN ?= example.com
HOST ?= example-$(CI_COMMIT_REF_SLUG)

gen:
		@protoc \
		-I apis/ \
		-I $(GOPATH)/src/ \
		-I $(GOPATH)/src/github.com/googleapis/googleapis \
		--grpc-gateway_out=logtostderr=true:apis \
		--swagger_out=logtostderr=true:apis \
		--go_out=plugins=grpc:apis/. hello.proto

vendor:
		@go mod vendor

## Docker
docker-build: gen vendor
		@docker build -t solcates/grpc-istio-example .

docker-push:
		@docker push solcates/grpc-istio-example

docker: docker-build docker-push

## Building
build:
		@go build --mod=vendor ./cmd/greeter


## Running local

run-server:
		@go run ./cmd/greeter server
run-client:
		go run ./cmd/greeter client --name LocalAlice

## Running Kubernetes
ARGS=   --set domain=$(DOMAIN) \
		--set host=$(HOST)
deploy: docker
		@helm upgrade --install --namespace $(HOST) $(HOST) . \
		$(ARGS) \
		--recreate-pods
run-client-remote:
		go run ./cmd/greeter client --host $(HOST).$(DOMAIN) --name K8SAlice

