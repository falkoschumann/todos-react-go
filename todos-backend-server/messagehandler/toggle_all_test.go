package messagehandler_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain"
	"todos_backend_server/messagehandler"
)

func TestToggleAll(t *testing.T) {
	t.Run("set all todos completed.", func(t *testing.T) {
		repo := provider.NewMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		toggleAll := messagehandler.ToggleAll(repo)

		status := toggleAll(domain.ToggleAllCommand{Checked: true})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("ToggleAll() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: true},
		}
		stored := repo.Load()
		if diff := cmp.Diff(want, stored); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})

	t.Run("set all todos active.", func(t *testing.T) {
		repo := provider.NewMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		toggleAll := messagehandler.ToggleAll(repo)

		status := toggleAll(domain.ToggleAllCommand{Checked: false})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("ToggleAll() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: false},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		}
		stored := repo.Load()
		if diff := cmp.Diff(want, stored); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})
}
