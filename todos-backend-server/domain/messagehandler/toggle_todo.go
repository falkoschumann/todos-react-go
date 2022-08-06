package messagehandler

import (
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/port"
)

func NewToggleTodo(repo port.TodosRepository) message.ToggleTodoCommandHandler {
	toggleAll := func(todos []data.Todo, id int) []data.Todo {
		var result []data.Todo
		for _, t := range todos {
			if t.Id == id {
				t.Completed = !t.Completed
			}
			result = append(result, t)
		}
		return result
	}

	return func(c message.ToggleTodoCommand) message.CommandStatus {
		todos := repo.Load()
		todos = toggleAll(todos, c.Id)
		repo.Store(todos)
		return message.MakeSuccess()
	}
}
