package messagehandler

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/domain"
)

func TestDestryTodo(t *testing.T) {
	t.Run("destroys a todo.", func(t *testing.T) {
		repo := newMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		destroyTodo := DestroyTodo(repo)

		status := destroyTodo(domain.DestroyTodoCommand{Id: 2})

		if diff := cmp.Diff(domain.Success(), status); diff != "" {
			t.Errorf("DestroyTodo() mismatch (-want +got):\n%s", diff)
		}
		want := []domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		}
		if diff := cmp.Diff(want, repo.Load()); diff != "" {
			t.Errorf("Todos mismatch (-want +got):\n%s", diff)
		}
	})
}
