# grpc-istio-example

## Overview

This repo is meant to reproduce an issue with using Istio's Gateway to do TLS Termination for a gRPC service running in Pods using sidecars.

The issue is that inbound traffic is not making it to the Pod

## What is in the app
It is comprised of:

- A greeter example App fromhttps://grpc.io/docs/quickstart/go/
    - It handles gRPC and REST(via gRPC-gateway) on ports tcp/31400 and tcp/8080 in clear text
- A docker image that will launch the application and serve up both gRPC and REST on interface 0.0.0.0
- A Helm Chart that will deploy the application on a k8s cluster with Istio also deployed.
    - It assumes Istio 1.2+ is installed
    - It also expects the ingressgateway to do TLS termination with something like cert-manager
    - It contains:
        - K8S Deployment launching Pods listening on tcp/31400 and tcp/8080 in clear text (Istio's MTLS SideCar overlay)
        - K8S Service for tcp/31400 and tcp/8080 in clear text (Istio's MTLS will overlay)
        - Istio Gateway for FQDN(example-dev.example.com by default) that will terminate at the Istio Ingress Gateway LoadBalancer Listening on HTTPS/31400 and HTTPS/443
        - Istio VirtualService for FQDN that routes port 31400 to the tcp/31400 service and all else to the tcp/8080 Service
        
