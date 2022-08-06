package provider

import (
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/port"
)

type MemoryTodosRepository struct {
	todos []data.Todo
}

func NewMemoryTodosRepository(todos []data.Todo) port.TodosRepository {
	return &MemoryTodosRepository{
		todos: todos,
	}
}

func (r *MemoryTodosRepository) Load() []data.Todo {
	return r.todos
}

func (r *MemoryTodosRepository) Store(todos []data.Todo) {
	r.todos = todos
}
