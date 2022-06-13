package messagehandler

import "todos_backend_server/domain"

func SelectTodos(r domain.TodosRepository) domain.SelectTodosQueryHandler {
	return func(q domain.SelectTodosQuery) domain.SelectTodosQueryResult {
		todos := r.Load()
		return domain.SelectTodosQueryResult{Todos: todos}
	}
}
