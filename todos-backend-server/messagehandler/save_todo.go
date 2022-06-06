package messagehandler

import (
	"todos_backend_server/domain"
)

func GetSaveTodo(repo domain.TodosRepository) domain.SaveTodoCommandHandler {
	saveTodo := func(todos []domain.Todo, id int, newTitle string) []domain.Todo {
		var r []domain.Todo
		for _, t := range todos {
			if t.Id == id {
				t.Title = newTitle
			}
			r = append(r, t)
		}
		return r
	}

	return func(c domain.SaveTodoCommand) domain.CommandStatus {
		todos := repo.Load()
		todos = saveTodo(todos, c.Id, c.NewTitle)
		repo.Store(todos)
		return domain.Success()
	}
}
