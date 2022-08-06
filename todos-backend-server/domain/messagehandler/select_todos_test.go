package messagehandler_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/messagehandler"
)

func TestSelectTodos(t *testing.T) {
	t.Run("returns all todos.", func(t *testing.T) {
		repo := provider.NewMemoryTodosRepository([]data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		})
		selectTodo := messagehandler.NewSelectTodos(repo)

		result := selectTodo(message.SelectTodosQuery{})

		want := message.SelectTodosQueryResult{
			Todos: []data.Todo{
				{Id: 1, Title: "Taste JavaScript", Completed: true},
				{Id: 2, Title: "Buy Unicorn", Completed: false},
			},
		}
		if diff := cmp.Diff(want, result); diff != "" {
			t.Errorf("SelectTodos() mismatch (-want +got):\n%s", diff)
		}
	})
}
