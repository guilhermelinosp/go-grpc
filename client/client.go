package main

import (
	"context"
	"log"
	"time"

	pb "github.com/guilhermelinosp/go-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect to server: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Ping"})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	log.Println("Response status:", res.Message)
}
