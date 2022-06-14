package messagehandler

import "todos_backend_server/domain"

func SaveTodo(repo domain.TodosRepository) domain.SaveTodoCommandHandler {
	saveTodo := func(todos []domain.Todo, id int, newTitle string) []domain.Todo {
		var result []domain.Todo
		for _, t := range todos {
			if t.Id != id {
				result = append(result, t)
				continue
			}

			if newTitle == "" {
				continue
			}

			t.Title = newTitle
			result = append(result, t)
		}
		return result
	}

	return func(c domain.SaveTodoCommand) domain.CommandStatus {
		todos := repo.Load()
		todos = saveTodo(todos, c.Id, c.NewTitle)
		repo.Store(todos)
		return domain.Success()
	}
}
