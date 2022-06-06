package messagehandler

import (
	"todos_backend_server/domain"
)

func GetToggleAll(repo domain.TodosRepository) domain.ToggleAllCommandHandler {
	toggleAll := func(todos []domain.Todo, checked bool) []domain.Todo {
		var r []domain.Todo
		for _, t := range todos {
			t.Completed = checked
			r = append(r, t)
		}
		return r
	}

	return func(c domain.ToggleAllCommand) domain.CommandStatus {
		todos := repo.Load()
		todos = toggleAll(todos, c.Checked)
		repo.Store(todos)
		return domain.Success()
	}
}
