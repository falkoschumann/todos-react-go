package messagehandler

import "todos_backend_server/domain"

type memoryTodosRepository struct {
	todos []domain.Todo
}

func newMemoryTodosRepository(todos []domain.Todo) *memoryTodosRepository {
	return &memoryTodosRepository{
		todos: todos,
	}
}

func (r *memoryTodosRepository) Load() []domain.Todo {
	return r.todos
}

func (r *memoryTodosRepository) Store(todos []domain.Todo) {
	r.todos = todos
}
