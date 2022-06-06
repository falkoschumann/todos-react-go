package messagehandler

import (
	"testing"
	"todos_backend_server/domain"

	"github.com/google/go-cmp/cmp"
)

func TestSelectTodos(t *testing.T) {
	repo := newMemoryTodosRepository([]domain.Todo{
		{Id: 1, Title: "Taste JavaScript", Completed: true},
		{Id: 2, Title: "Buy Unicorn", Completed: false},
	})
	selectTodo := GetSelectTodos(repo)

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
}
