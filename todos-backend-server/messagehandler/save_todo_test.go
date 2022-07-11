package messagehandler_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain"
	"todos_backend_server/messagehandler"
)

func TestSaveTodo(t *testing.T) {
	t.Run("changes todos title.", func(t *testing.T) {
		repo := provider.NewMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		saveTodo := messagehandler.SaveTodo(repo)

		status := saveTodo(domain.SaveTodoCommand{Id: 1, NewTitle: "Taste TypeScript"})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("SaveTodo() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 1, Title: "Taste TypeScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		}
		stored := repo.Load()
		if diff := cmp.Diff(want, stored); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("destroys todo if title is empty.", func(t *testing.T) {
		repo := provider.NewMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		saveTodo := messagehandler.SaveTodo(repo)

		status := saveTodo(domain.SaveTodoCommand{Id: 1, NewTitle: ""})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("SaveTodo() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		}
		stored := repo.Load()
		if diff := cmp.Diff(want, stored); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})
}
