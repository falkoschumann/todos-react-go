package messagehandler

import (
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/port"
)

func NewAddTodo(repo port.TodosRepository) message.AddTodoCommandHandler {
	getNextId := func(todos []data.Todo) int {
		var id = 0
		for _, t := range todos {
			if t.Id > id {
				id = t.Id
			}
		}
		id++
		return id
	}

	addTodo := func(todos []data.Todo, title string) []data.Todo {
		id := getNextId(todos)
		t := data.Todo{Id: id, Title: title, Completed: false}
		return append(todos, t)
	}

	return func(c message.AddTodoCommand) message.CommandStatus {
		var title = sanitizedTitle(c.Title)
		if title == "" {
			return message.MakeSuccess()
		}

		todos := repo.Load()
		todos = addTodo(todos, title)
		repo.Store(todos)
		return message.MakeSuccess()
	}
}
