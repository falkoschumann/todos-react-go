package messagehandler_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/messagehandler"
)

var addTodoTests = []struct {
	name                string
	givenStoredTodos    []data.Todo
	command             message.AddTodoCommand
	status              message.CommandStatus
	expectedStoredTodos []data.Todo
}{
	{
		name: "saves new todo.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		},
		command: message.AddTodoCommand{Title: "Buy Unicorn"},
		status:  message.MakeSuccess(),
		expectedStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
	},
	{
		name: "saves trimmed title.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		},
		command: message.AddTodoCommand{Title: "  Buy Unicorn   "},
		status:  message.MakeSuccess(),
		expectedStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
	},
	{
		name: "does nothing if title is empty.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		},
		command: message.AddTodoCommand{Title: "  "},
		status:  message.MakeSuccess(),
		expectedStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		},
	},
}

func TestAddTodo(t *testing.T) {
	for _, tt := range addTodoTests {
		t.Run(tt.name, func(t *testing.T) {
			repo := provider.NewMemoryTodosRepository(tt.givenStoredTodos)
			addTodo := messagehandler.NewAddTodo(repo)

			gotStatus := addTodo(tt.command)
			gotTodos := repo.Load()

			if !cmp.Equal(tt.status, gotStatus) {
				t.Errorf("AddTodo(%v) = %v, want %v", tt.command, gotStatus, tt.status)
			}
			if diff := cmp.Diff(tt.expectedStoredTodos, gotTodos); diff != "" {
				t.Errorf("AddTodo() stored todos mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
