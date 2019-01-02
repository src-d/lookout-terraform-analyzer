package main

import (
	"context"
	"fmt"
	"net"
	"os"

	terraformanalyzer "github.com/src-d/lookout-terraform-analyzer"

	"gopkg.in/src-d/lookout-sdk.v0/pb"

	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	log "gopkg.in/src-d/go-log.v1"
)

var (
	name    = "terraform-analyzer"
	version string
	build   string
)

const maxMessageSize = 100 * 1024 * 1024 //100mb

type config struct {
	Host           string `envconfig:"HOST" default:"0.0.0.0"`
	Port           int    `envconfig:"PORT" default:"9930"`
	DataServiceURL string `envconfig:"DATA_SERVICE_URL" default:"ipv4://localhost:10301"`
	LogLevel       string `envconfig:"LOG_LEVEL" default:"info"`
}

func main() {
	var conf config
	envconfig.MustProcess("LOOKOUT_TERRAFORM", &conf)

	log.DefaultFactory = &log.LoggerFactory{Level: conf.LogLevel}
	log.DefaultLogger = log.New(nil)

	grpcAddr, err := pb.ToGoGrpcAddress(conf.DataServiceURL)
	if err != nil {
		log.Errorf(err, "failed to parse DataService address %s", conf.DataServiceURL)
		return
	}

	conn, err := pb.DialContext(
		context.Background(),
		grpcAddr,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.FailFast(false)),
	)
	if err != nil {
		log.Errorf(err, "cannot create connection to DataService %s", grpcAddr)
		os.Exit(1)
	}

	defer conn.Close()

	analyzer := &terraformanalyzer.Analyzer{
		DataClient: pb.NewDataClient(conn),
		Version:    version,
	}

	opts := []grpc.ServerOption{
		grpc.MaxRecvMsgSize(maxMessageSize),
		grpc.MaxSendMsgSize(maxMessageSize),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterAnalyzerServer(server, analyzer)

	analyzerURL := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	lis, err := net.Listen("tcp", analyzerURL)
	if err != nil {
		log.Errorf(err, "failed to start analyzer gRPC server on %s", analyzerURL)
		os.Exit(1)
	}

	log.Infof("server has started on '%s'", analyzerURL)
	err = server.Serve(lis)
	if err != nil {
		log.Errorf(err, "gRPC server failed listening on %v", lis)
		os.Exit(1)
	}
}
