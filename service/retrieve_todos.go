package service

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/muhammadisa/vanilla-microservice/protobuf"
)

func (t todoService) RetrieveTodos(ctx context.Context, blank *empty.Empty) (*pb.Todos, error) {
	pbTodo := make([]*pb.Todo, 0)
	todos := t.writer.ReadTodos(ctx)
	for _, v := range todos {
		pbTodo = append(pbTodo, &pb.Todo{
			Id:        v.ID.String(),
			Name:      v.Name,
			Completed: v.Completed,
		})
	}
	return &pb.Todos{
		Todos: pbTodo,
	}, nil
}
