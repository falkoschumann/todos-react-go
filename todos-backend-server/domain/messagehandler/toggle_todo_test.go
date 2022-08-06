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
	name                string
	givenStoredTodos    []data.Todo
	command             message.ToggleTodoCommand
	status              message.CommandStatus
	expectedStoredTodos []data.Todo
}{
	{
		name: "activates a todo.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
		command: message.ToggleTodoCommand{Id: 1},
		status:  message.MakeSuccess(),
		expectedStoredTodos: []data.Todo{
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
		command: message.ToggleTodoCommand{Id: 2},
		status:  message.MakeSuccess(),
		expectedStoredTodos: []data.Todo{
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

			gotStatus := toggleTodo(tt.command)
			gotStoredTodos := repo.Load()

			if !cmp.Equal(tt.status, gotStatus) {
				t.Errorf("ToggleTodo(%v) = %v, want %v", tt.command, gotStatus, tt.status)
			}
			if diff := cmp.Diff(tt.expectedStoredTodos, gotStoredTodos); diff != "" {
				t.Errorf("ToggleTodo() Stored todos mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
