package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "parser_client/app/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	//TODO:обработать в замыкании ошибку функции close()
	defer conn.Close()
	client := pb.NewParseServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//TODO:обработать в замыкании ошибку функции cancel()
	defer cancel()
	response, err := client.GetTable(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User: %v", response)
}
