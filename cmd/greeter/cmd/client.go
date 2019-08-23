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
	"github.com/sirupsen/logrus"
	"github.com/solcates/grpc-istio-example/apis"
	"google.golang.org/grpc"
	"log"

	"github.com/spf13/cobra"
)

var name string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Run Echo Client",

	Run: func(cmd *cobra.Command, args []string) {

		var conn *grpc.ClientConn
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, grpcPort), grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()

		gc := apis.NewGreeterClient(conn)
		resp, err := gc.SayHello(context.Background(), &apis.HelloRequest{
			Name: name,
		})
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Info(resp)
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().StringVar(&host, "host", "localhost", "Host to connect to")
	clientCmd.Flags().StringVar(&name, "name", "Alice", "Host to connect to")
}
