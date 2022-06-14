package messagehandler

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/domain"
)

func TestSelectTodos(t *testing.T) {
	t.Run("returns all todos.", func(t *testing.T) {
		repo := newMemoryTodosRepository([]domain.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		selectTodo := SelectTodos(repo)

		result := selectTodo(domain.SelectTodosQuery{})

		want := domain.SelectTodosQueryResult{
			Todos: []domain.Todo{
				{Id: 1, Title: "Taste JavaScript", Completed: true},
				{Id: 2, Title: "Buy Unicorn", Completed: false},
			},
		}
		if diff := cmp.Diff(want, result); diff != "" {
			t.Errorf("SelectTodos() mismatch (-want +got):\n%s", diff)
		}
	})
}
