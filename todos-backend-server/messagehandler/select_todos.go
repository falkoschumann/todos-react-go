package messagehandler

import "todos_backend_server/domain"

func GetSelectTodos(r domain.TodosRepository) domain.SelectTodosQueryHandler {
	return func(q domain.SelectTodosQuery) domain.SelectTodosQueryResult {
		t := r.Load()
		return domain.SelectTodosQueryResult{Todos: t}
	}
}
