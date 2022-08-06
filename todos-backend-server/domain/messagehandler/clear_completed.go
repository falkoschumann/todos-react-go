package messagehandler

import (
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/port"
)

func NewClearCompleted(repo port.TodosRepository) message.ClearCompletedCommandHandler {
	clearCompleted := func(todos []data.Todo) []data.Todo {
		var result []data.Todo
		for _, t := range todos {
			if !t.Completed {
				result = append(result, t)
			}
		}
		return result
	}

	return func(c message.ClearCompletedCommand) message.CommandStatus {
		todos := repo.Load()
		todos = clearCompleted(todos)
		repo.Store(todos)
		return message.MakeSuccess()
	}
}
