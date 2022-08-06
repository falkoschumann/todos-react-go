package messagehandler

import (
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/port"
)

func NewSaveTodo(repo port.TodosRepository) message.SaveTodoCommandHandler {
	saveTodo := func(todos []data.Todo, id int, newTitle string) []data.Todo {
		var result []data.Todo
		for _, t := range todos {
			if t.Id != id {
				result = append(result, t)
				continue
			}

			if newTitle == "" {
				continue
			}

			t.Title = newTitle
			result = append(result, t)
		}
		return result
	}

	return func(c message.SaveTodoCommand) message.CommandStatus {
		var newTitle = sanitizedTitle(c.NewTitle)
		todos := repo.Load()
		todos = saveTodo(todos, c.Id, newTitle)
		repo.Store(todos)
		return message.MakeSuccess()
	}
}
