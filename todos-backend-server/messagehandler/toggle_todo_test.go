package messagehandler_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain"
	"todos_backend_server/messagehandler"
)

func TestToggleTodo(t *testing.T) {
	t.Run("activates a todo.", func(t *testing.T) {
		repo := provider.NewMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		toggleTodo := messagehandler.ToggleTodo(repo)

		status := toggleTodo(domain.ToggleTodoCommand{Id: 1})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("ToggleTodo() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: false},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		}
		if diff := cmp.Diff(want, repo.Load()); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("completes a todo.", func(t *testing.T) {
		repo := provider.NewMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		toggleTodo := messagehandler.ToggleTodo(repo)

		status := toggleTodo(domain.ToggleTodoCommand{Id: 2})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("ToggleTodo() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: true},
		}
		if diff := cmp.Diff(want, repo.Load()); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})
}
