package messagehandler_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/messagehandler"
)

func TestDestryTodo(t *testing.T) {
	t.Run("destroys a todo.", func(t *testing.T) {
		repo := provider.NewMemoryTodosRepository([]data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		destroyTodo := messagehandler.NewDestroyTodo(repo)

		whenCommand := message.DestroyTodoCommand{Id: 2}
		status := destroyTodo(whenCommand)
		storedTodos := repo.Load()

		thenStatus := message.MakeSuccess()
		if !cmp.Equal(thenStatus, status) {
			t.Errorf("ClearCompleted(%v) = %v, want %v", whenCommand, status, thenStatus)
		}
		thenStoredTodos := []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		}
		if diff := cmp.Diff(thenStoredTodos, storedTodos); diff != "" {
			t.Errorf("ClearCompleted() stored todos mismatch (-want +got):\n%s", diff)
		}
	})
}
