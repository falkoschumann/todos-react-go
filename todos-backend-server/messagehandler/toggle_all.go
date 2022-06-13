package messagehandler

import (
	"todos_backend_server/domain"
)

func ToggleAll(repo domain.TodosRepository) domain.ToggleAllCommandHandler {
	toggleAll := func(todos []domain.Todo, checked bool) []domain.Todo {
		var result []domain.Todo
		for _, t := range todos {
			t.Completed = checked
			result = append(result, t)
		}
		return result
	}

	return func(c domain.ToggleAllCommand) domain.CommandStatus {
		todos := repo.Load()
		todos = toggleAll(todos, c.Checked)
		repo.Store(todos)
		return domain.Success()
	}
}
