package messagehandler

import (
	"testing"
	"todos_backend_server/domain"

	"github.com/google/go-cmp/cmp"
)

func TestSaveTodo(t *testing.T) {
	repo := newMemoryTodosRepository([]domain.Todo{
		{Id: 1, Title: "Taste JavaScript", Completed: true},
		{Id: 2, Title: "Buy Unicorn", Completed: false},
	})
	saveTodo := GetSaveTodo(repo)

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
}
