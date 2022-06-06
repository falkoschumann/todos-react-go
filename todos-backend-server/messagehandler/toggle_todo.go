package messagehandler

import (
	"todos_backend_server/domain"
)

func GetToggleTodo(repo domain.TodosRepository) domain.ToggleTodoCommandHandler {
	toggleAll := func(todos []domain.Todo, id int) []domain.Todo {
		var r []domain.Todo
		for _, t := range todos {
			if t.Id == id {
				t.Completed = !t.Completed
			}
			r = append(r, t)
		}
		return r
	}

	return func(c domain.ToggleTodoCommand) domain.CommandStatus {
		todos := repo.Load()
		todos = toggleAll(todos, c.Id)
		repo.Store(todos)
		return domain.Success()
	}
}
