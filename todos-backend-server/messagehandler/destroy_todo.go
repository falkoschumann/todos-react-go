package messagehandler

import "todos_backend_server/domain"

func DestroyTodo(repo domain.TodosRepository) domain.DestroyTodoCommandHandler {
	destroyTodo := func(todos []domain.Todo, id int) []domain.Todo {
		var result []domain.Todo
		for _, t := range todos {
			if t.Id == id {
				continue
			}

			result = append(result, t)
		}
		return result
	}

	return func(c domain.DestroyTodoCommand) domain.CommandStatus {
		todos := repo.Load()
		todos = destroyTodo(todos, c.Id)
		repo.Store(todos)
		return domain.Success()
	}
}
