package messagehandler_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/messagehandler"
)

func TestClearCompleted(t *testing.T) {
	t.Run("removes completed todos.", func(t *testing.T) {
		repo := provider.NewMemoryTodosRepository([]data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		clearCompleted := messagehandler.NewClearCompleted(repo)

		cmd := message.ClearCompletedCommand{}
		gotStatus := clearCompleted(cmd)
		gotStoredTodos := repo.Load()

		wantStatus := message.MakeSuccess()
		if !cmp.Equal(wantStatus, gotStatus) {
			t.Errorf("ClearCompleted(%v) = %v, want %v", cmd, gotStatus, wantStatus)
		}
		wantStoredTodos := []data.Todo{
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		}
		if diff := cmp.Diff(wantStoredTodos, gotStoredTodos); diff != "" {
			t.Errorf("ClearCompleted() stored todos mismatch (-want +got):\n%s", diff)
		}
	})
}
