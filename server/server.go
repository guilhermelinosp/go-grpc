package main

import (
	"context"
	"log"
	"net"

	pb "github.com/guilhermelinosp/go-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	peer, ok := peer.FromContext(ctx)
	if ok {
		log.Printf("Request received from %s\n", peer.Addr.String())
	}

	return &pb.HelloResponse{Message: req.Name + " Pong"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &server{})

	log.Println("gRPC server running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
