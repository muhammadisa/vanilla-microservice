package repository

import (
	"context"
	"github.com/muhammadisa/vanilla-microservice/model"
	uuid "github.com/satori/go.uuid"
)

func (t *todoRepository) ReadTodos(ctx context.Context) model.Todos {
	return model.Todos{
		{
			ID:        uuid.NewV4(),
			Name:      "Bath",
			Completed: false,
		},
	}
}
