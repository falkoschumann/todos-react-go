package messagehandler

import (
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/port"
)

func NewDestroyTodo(repo port.TodosRepository) message.DestroyTodoCommandHandler {
	destroyTodo := func(todos []data.Todo, id int) []data.Todo {
		var result []data.Todo
		for _, t := range todos {
			if t.Id == id {
				continue
			}

			result = append(result, t)
		}
		return result
	}

	return func(c message.DestroyTodoCommand) message.CommandStatus {
		todos := repo.Load()
		todos = destroyTodo(todos, c.Id)
		repo.Store(todos)
		return message.MakeSuccess()
	}
}
