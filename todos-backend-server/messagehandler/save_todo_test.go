package messagehandler

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/domain"
)

func TestSaveTodo(t *testing.T) {
	t.Run("changes todos title.", func(t *testing.T) {
		repo := newMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		saveTodo := SaveTodo(repo)

		status := saveTodo(domain.SaveTodoCommand{Id: 1, NewTitle: "Taste TypeScript"})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("SaveTodo() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 1, Title: "Taste TypeScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		}
		if diff := cmp.Diff(want, repo.Load()); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("destroys todo if title is empty.", func(t *testing.T) {
		repo := newMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		saveTodo := SaveTodo(repo)

		status := saveTodo(domain.SaveTodoCommand{Id: 1, NewTitle: ""})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("SaveTodo() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		}
		if diff := cmp.Diff(want, repo.Load()); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})
}
