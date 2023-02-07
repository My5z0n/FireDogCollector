package main

import (
	"log"
	"net"

	"github.com/My5z0n/FireDogCollector/OtelCollector/repository"
	"github.com/My5z0n/FireDogCollector/OtelCollector/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	//pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"github.com/My5z0n/FireDogCollector/OtelCollector/api"
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	//commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	//resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
	//tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
)

func main() {
	var config = utils.GetEnvConfig()
	netAddr := config.OtelUrl + ":" + config.OtelPort

	lis, err := net.Listen("tcp", netAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	r, err := repository.NewTraceRepository(config)
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	coltracepb.RegisterTraceServiceServer(s, &api.Server{
		TraceRepository: r,
	})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
