package _interface

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/muhammadisa/vanilla-microservice/protobuf"
)

type TodoService interface {
	CreateTodo(ctx context.Context, todo *pb.Todo) (*empty.Empty, error)
	RetrieveTodos(ctx context.Context, blank *empty.Empty) (*pb.Todos, error)
}
