package messagehandler

import (
	"todos_backend_server/domain"
)

func AddTodo(repo domain.TodosRepository) domain.AddTodoCommandHandler {
	getNextId := func(todos []domain.Todo) int {
		var id = 0
		for _, t := range todos {
			if t.Id > id {
				id = t.Id
			}
		}
		id++
		return id
	}

	addTodo := func(todos []domain.Todo, title string) []domain.Todo {
		id := getNextId(todos)
		t := domain.Todo{Id: id, Title: title, Completed: false}
		return append(todos, t)
	}

	return func(c domain.AddTodoCommand) domain.CommandStatus {
		if c.Title == "" {
			return domain.Success()
		}

		todos := repo.Load()
		todos = addTodo(todos, c.Title)
		repo.Store(todos)
		return domain.Success()
	}
}
