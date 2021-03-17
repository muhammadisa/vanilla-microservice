package repository

import (
	"database/sql"
	_interface "github.com/muhammadisa/vanilla-microservice/repository/interface"
)

// todoRepository holdings every repo needed
type todoRepository struct {
	db *sql.DB
}

// NewTodoRepository create instance of repository
func NewTodoRepository(database *sql.DB) _interface.Repository {
	return &todoRepository{
		db: database,
	}
}