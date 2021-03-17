package service

import (
	_interfacerepo "github.com/muhammadisa/vanilla-microservice/repository/interface"
	_interfacesvc "github.com/muhammadisa/vanilla-microservice/service/interface"
)

// todoService holdings every service needed
type todoService struct {
	writer _interfacerepo.Repository
}

// NewTodoService create instance of service
func NewTodoService(w _interfacerepo.Repository) _interfacesvc.TodoService {
	return &todoService{
		writer: w,
	}
}
