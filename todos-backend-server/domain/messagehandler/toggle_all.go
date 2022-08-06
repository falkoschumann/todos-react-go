package messagehandler

import (
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/port"
)

func NewToggleAll(repo port.TodosRepository) message.ToggleAllCommandHandler {
	toggleAll := func(todos []data.Todo, checked bool) []data.Todo {
		var result []data.Todo
		for _, t := range todos {
			t.Completed = checked
			result = append(result, t)
		}
		return result
	}

	return func(c message.ToggleAllCommand) message.CommandStatus {
		todos := repo.Load()
		todos = toggleAll(todos, c.Checked)
		repo.Store(todos)
		return message.MakeSuccess()
	}
}
