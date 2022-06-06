package messagehandler

import "todos_backend_server/domain"

func GetClearCompleted(repo domain.TodosRepository) domain.ClearCompletedCommandHandler {
	clearCompleted := func(todos []domain.Todo) []domain.Todo {
		var r []domain.Todo
		for _, t := range todos {
			if !t.Completed {
				r = append(r, t)
			}
		}
		return r
	}

	return func(c domain.ClearCompletedCommand) domain.CommandStatus {
		todos := repo.Load()
		todos = clearCompleted(todos)
		repo.Store(todos)
		return domain.Success()
	}
}
