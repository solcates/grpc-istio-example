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
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/solcates/grpc-istio-example/apis"
	"github.com/solcates/grpc-istio-example/pkg/greeter"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run Echo Server",

	Run: func(cmd *cobra.Command, args []string) {

		s := greeter.NewHelloServer()

		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 31400))
		if err != nil {
			logrus.Fatalf("failed to listen: %v", err)
		}

		// create a gRPC server object
		grpcServer := grpc.NewServer()
		// attach the Ping service to the server
		apis.RegisterGreeterServer(grpcServer, s)
		logrus.Info("Starting Greeter Server")
		logrus.Infof("Listening on 0.0.0.0:%d ", 31400)
		if err := grpcServer.Serve(lis); err != nil {
			logrus.Fatalf("failed to serve: %s", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
