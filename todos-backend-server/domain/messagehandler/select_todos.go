package messagehandler

import (
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/port"
)

func NewSelectTodos(r port.TodosRepository) message.SelectTodosQueryHandler {
	return func(q message.SelectTodosQuery) message.SelectTodosQueryResult {
		todos := r.Load()
		return message.SelectTodosQueryResult{Todos: todos}
	}
}
