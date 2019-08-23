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
docker-build: vendor
		@docker build -t solcates/grpc-istio-example .

docker-push:
		@docker push solcates/grpc-istio-example

docker: docker-build docker-push

## Building
build:
		@go build --mod=vendor ./cmd/greeter


## Running
run-client-remote:
		@go run ./cmd/greeter client --host $(HOST).$(DOMAIN)

run-server:
		@go run ./cmd/greeter server
## Deployments


ARGS=   --set domain=$(DOMAIN) \
		--set host=$(HOST)
deploy:
		@helm upgrade --install --namespace $(HOST) $(HOST) . \
		$(ARGS) \
		--recreate-pods