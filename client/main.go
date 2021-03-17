package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/muhammadisa/vanilla-microservice/protobuf"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := pb.NewTodoServiceClient(conn)
	ctx := context.Background()
	todos, err := client.RetrieveTodos(ctx, &empty.Empty{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("Todos list : %+v", todos))
}
