package greeter

import (
	"context"
	"github.com/solcates/grpc-istio-example/apis"
	"reflect"
	"testing"
)

func TestHelloServer_SayHello(t *testing.T) {
	type args struct {
		ctx context.Context
		req *apis.HelloRequest
	}
	tests := []struct {
		name    string
		args    args
		wantRes *apis.HelloReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				req: &apis.HelloRequest{
					Name: "Alice",
				},
			},
			wantRes: &apis.HelloReply{
				Message: "Hello, Alice",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HelloServer{}
			gotRes, err := h.SayHello(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("SayHello() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestHelloServer_SayHelloAgain(t *testing.T) {
	type args struct {
		ctx context.Context
		req *apis.HelloRequest
	}
	tests := []struct {
		name    string
		args    args
		wantRes *apis.HelloReply
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ctx: context.Background(),
				req: &apis.HelloRequest{
					Name: "Alice",
				},
			},
			wantRes: &apis.HelloReply{
				Message: "Hello Again, Alice",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := HelloServer{}
			gotRes, err := h.SayHelloAgain(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SayHelloAgain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("SayHelloAgain() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestNewHelloServer(t *testing.T) {
	tests := []struct {
		name string
		want *HelloServer
	}{
		{
			name: "OK",
			want: &HelloServer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHelloServer(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHelloServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
