package main

import (
	"context"
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/src-d/lookout"
	terraformanalyzer "github.com/src-d/lookout-terraform-analyzer"
	"github.com/src-d/lookout/util/grpchelper"
	"google.golang.org/grpc"
	"gopkg.in/src-d/go-log.v1"
)

var (
	name    = "terraform-analyzer"
	version string
	build   string
)

type config struct {
	Host           string `envconfig:"HOST" default:"0.0.0.0"`
	Port           int    `envconfig:"PORT" default:"2001"`
	DataServiceURL string `envconfig:"DATA_SERVICE_URL" default:"ipv4://localhost:10301"`
	LogLevel       string `envconfig:"LOG_LEVEL" default:"info"`
}

func main() {
	var conf config
	envconfig.MustProcess("LOOKOUT_TERRAFORM", &conf)

	log.DefaultFactory = &log.LoggerFactory{Level: conf.LogLevel}
	log.DefaultLogger = log.New(nil)

	grpcAddr, err := grpchelper.ToGoGrpcAddress(conf.DataServiceURL)
	if err != nil {
		log.Errorf(err, "failed to parse DataService addres %s", conf.DataServiceURL)
		return
	}

	conn, err := grpchelper.DialContext(
		context.Background(),
		grpcAddr,
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.FailFast(false)),
	)
	if err != nil {
		log.Errorf(err, "cannot create connection to DataService %s", grpcAddr)
		return
	}

	analyzer := terraformanalyzer.Analyzer{
		DataClient: lookout.NewDataClient(conn),
		Version:    version,
	}

	server := grpchelper.NewServer()
	lookout.RegisterAnalyzerServer(server, analyzer)

	analyzerURL := fmt.Sprintf("ipv4://%s:%d", conf.Host, conf.Port)
	lis, err := grpchelper.Listen(analyzerURL)
	if err != nil {
		log.Errorf(err, "failed to start analyzer gRPC server on %s", analyzerURL)
		return
	}

	log.Infof("server has started on '%s'", analyzerURL)
	err = server.Serve(lis)
	if err != nil {
		log.Errorf(err, "gRPC server failed listening on %v", lis)
	}
}
