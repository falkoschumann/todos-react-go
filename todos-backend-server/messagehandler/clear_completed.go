package messagehandler

import "todos_backend_server/domain"

func ClearCompleted(repo domain.TodosRepository) domain.ClearCompletedCommandHandler {
	clearCompleted := func(todos []domain.Todo) []domain.Todo {
		var result []domain.Todo
		for _, t := range todos {
			if !t.Completed {
				result = append(result, t)
			}
		}
		return result
	}

	return func(c domain.ClearCompletedCommand) domain.CommandStatus {
		todos := repo.Load()
		todos = clearCompleted(todos)
		repo.Store(todos)
		return domain.Success()
	}
}
