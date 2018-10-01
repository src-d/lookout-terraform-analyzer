package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"gopkg.in/src-d/go-log.v1"
	"gopkg.in/src-d/lookout-sdk.v0/pb"
)

var portToListen = 2020
var dataSrvAddr = "localhost:10301"
var version = "alpha"
var maxMessageSize = 100 * 1024 * 1024 //100mb

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", portToListen))
	if err != nil {
		log.Errorf(err, "failed to listen on port: %d", portToListen)
	}

	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(maxMessageSize),
		grpc.MaxSendMsgSize(maxMessageSize),
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAnalyzerServer(grpcServer, &analyzer{})
	log.Infof("starting gRPC Analyzer server at port %d", portToListen)
	grpcServer.Serve(lis)
}
