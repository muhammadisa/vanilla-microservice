package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/muhammadisa/vanilla-microservice/model"
	pb "github.com/muhammadisa/vanilla-microservice/protobuf"
	uuid "github.com/satori/go.uuid"
)

func (t todoService) CreateTodo(ctx context.Context, todo *pb.Todo) (*empty.Empty, error) {
		uid, err := uuid.FromString(todo.Id)
		if err != nil {
			return &empty.Empty{}, err
		}
		err = t.writer.WriteTodo(ctx, model.Todo{
			ID:        uid,
			Name:      todo.Name,
			Completed: todo.Completed,
		})
		if err != nil {
			return &empty.Empty{}, err
		}
		return &empty.Empty{}, nil
}