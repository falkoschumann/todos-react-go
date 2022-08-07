package messagehandler_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/messagehandler"
)

var toggleTodoTests = []struct {
	name             string
	givenStoredTodos []data.Todo
	whenCommand      message.ToggleTodoCommand
	thenStatus       message.CommandStatus
	thenStoredTodos  []data.Todo
}{
	{
		name: "activates a todo.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
		whenCommand: message.ToggleTodoCommand{Id: 1},
		thenStatus:  message.MakeSuccess(),
		thenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: false},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
	},
	{
		name: "completes a todo.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
		whenCommand: message.ToggleTodoCommand{Id: 2},
		thenStatus:  message.MakeSuccess(),
		thenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: true},
		},
	},
}

func TestToggleTodo(t *testing.T) {
	for _, tt := range toggleTodoTests {
		t.Run(tt.name, func(t *testing.T) {
			repo := provider.NewMemoryTodosRepository(tt.givenStoredTodos)
			toggleTodo := messagehandler.NewToggleTodo(repo)

			gotStatus := toggleTodo(tt.whenCommand)
			gotStoredTodos := repo.Load()

			if !cmp.Equal(tt.thenStatus, gotStatus) {
				t.Errorf("ToggleTodo(%v) = %v, want %v", tt.whenCommand, gotStatus, tt.thenStatus)
			}
			if diff := cmp.Diff(tt.thenStoredTodos, gotStoredTodos); diff != "" {
				t.Errorf("ToggleTodo() Stored todos mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
