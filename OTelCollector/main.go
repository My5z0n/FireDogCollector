package main

import (
	"github.com/My5z0n/FireDogCollector/OtelCollector/repository"
	"log"
	"net"

	"google.golang.org/grpc"
	//pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"github.com/My5z0n/FireDogCollector/OtelCollector/api"
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	//commonpb "go.opentelemetry.io/proto/otlp/common/v1"
	//resourcepb "go.opentelemetry.io/proto/otlp/resource/v1"
	//tracepb "go.opentelemetry.io/proto/otlp/trace/v1"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:4320")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	r, err := repository.NewTraceRepository("9001", "helloworld")

	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}

	s := grpc.NewServer()
	coltracepb.RegisterTraceServiceServer(s, &api.Server{
		Abba:            "hello",
		TraceRepository: r,
	})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
