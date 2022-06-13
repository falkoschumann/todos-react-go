package messagehandler

import (
	"testing"
	"todos_backend_server/domain"

	"github.com/google/go-cmp/cmp"
)

func TestToggleTodo(t *testing.T) {
	t.Run("activates a todo.", func(t *testing.T) {
		repo := newMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		toggleTodo := ToggleTodo(repo)

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
		repo := newMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		toggleTodo := ToggleTodo(repo)

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
