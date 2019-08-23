.PHONY: all gen vendor

gen:
		@protoc --go_out=plugins=grpc:. apis/hello.proto

vendor:
		@go mod vendor

docker-build: vendor
		@docker build -t github.com/solcates/grpc-istio-example .

docker-push:
		@docker push github.com/solcates/grpc-istio-example

docker: docker-build docker-push