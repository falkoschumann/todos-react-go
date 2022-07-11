package provider

import "todos_backend_server/domain"

type MemoryTodosRepository struct {
	todos []domain.Todo
}

func NewMemoryTodosRepository(todos []domain.Todo) *MemoryTodosRepository {
	return &MemoryTodosRepository{
		todos: todos,
	}
}

func (r *MemoryTodosRepository) Load() []domain.Todo {
	return r.todos
}

func (r *MemoryTodosRepository) Store(todos []domain.Todo) {
	r.todos = todos
}
