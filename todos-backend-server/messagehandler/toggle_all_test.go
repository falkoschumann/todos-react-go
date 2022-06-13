package messagehandler

import (
	"testing"
	"todos_backend_server/domain"

	"github.com/google/go-cmp/cmp"
)

func TestToggleAll(t *testing.T) {
	t.Run("set all todos completed.", func(t *testing.T) {
		repo := newMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		toggleAll := ToggleAll(repo)

		status := toggleAll(domain.ToggleAllCommand{Checked: true})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("ToggleAll() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: true},
		}
		if diff := cmp.Diff(want, repo.Load()); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("set all todos active.", func(t *testing.T) {
		repo := newMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		toggleAll := ToggleAll(repo)

		status := toggleAll(domain.ToggleAllCommand{Checked: false})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("ToggleAll() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: false},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		}
		if diff := cmp.Diff(want, repo.Load()); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})
}
