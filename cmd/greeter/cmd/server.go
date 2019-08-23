/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"github.com/solcates/grpc-istio-example/apis"
	"github.com/solcates/grpc-istio-example/pkg/greeter"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run Echo Server",

	Run: func(cmd *cobra.Command, args []string) {
		stop := make(chan error)
		// Start the GRPC Server
		go runGRPC(stop)

		// Start the REST Server
		go runREST(stop)

		// Wait till someone says to stop...
		select {
		case err := <-stop:
			if err != nil {
				logrus.Fatal(err)
			}
		}
	},
}

func runGRPC(stop chan<- error) {
	var err error
	defer func() {
		stop <- err
	}()
	s := greeter.NewHelloServer()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	// create a gRPC server object
	grpcServer := grpc.NewServer()
	// attach the Ping service to the server
	apis.RegisterGreeterServer(grpcServer, s)
	logrus.Infof("gRPC Server Listening on 0.0.0.0:%d ", grpcPort)
	if err = grpcServer.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %s", err)
	}
	return
}

func runREST(stop chan<- error) {
	var err error
	defer func() {
		stop <- err
	}()

	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err = apis.RegisterGreeterHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%v:%v", "127.0.0.1", grpcPort), opts); err != nil {
		return
	}
	logrus.Infof("REST Server listening on %v:%v ", host, restPort)

	if err = http.ListenAndServe(fmt.Sprintf("%v:%v", host, restPort), mux); err != nil {
		return
	}
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVar(&host, "host", "0.0.0.0", "Host to listen on")

}
