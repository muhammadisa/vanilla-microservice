package repository

import (
	"context"
	"github.com/muhammadisa/vanilla-microservice/model"
	"log"
)

func (t *todoRepository) WriteTodo(ctx context.Context, todo model.Todo) error {
	log.Println(todo)
	return nil
}
