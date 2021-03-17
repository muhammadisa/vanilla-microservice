package _interface

import (
	"context"
	"github.com/muhammadisa/vanilla-microservice/model"
)

type Repository interface {
	WriteTodo(ctx context.Context, todo model.Todo) error
	ReadTodos(ctx context.Context) model.Todos
}
