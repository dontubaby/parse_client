package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"parser_client/app/pb"
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
	response, err := client.GetTable(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	for k, _ := range response.Table {
		fmt.Printf("key: %s\n", k)
	}

	//fmt.Printf("key: %s value: %v\n", k, v)
	fmt.Printf("key: %s\n", response.Table["Арт гекса"])

}
