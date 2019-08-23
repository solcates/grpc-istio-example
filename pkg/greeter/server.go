package greeter

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/solcates/grpc-istio-example/apis"
)

type HelloServer struct {
}

func NewHelloServer() *HelloServer {
	return &HelloServer{}
}

func (h HelloServer) SayHello(ctx context.Context, req *apis.HelloRequest) (res *apis.HelloReply, err error) {
	logrus.Debugf("req: %v", req)
	res = &apis.HelloReply{
		Message: fmt.Sprintf("Hello, %v", req.Name),
	}
	return
}

func (h HelloServer) SayHelloAgain(ctx context.Context, req *apis.HelloRequest) (res *apis.HelloReply, err error) {
	logrus.Debugf("req: %v", req)
	res = &apis.HelloReply{
		Message: fmt.Sprintf("Hello Again, %v", req.Name),
	}
	return
}
