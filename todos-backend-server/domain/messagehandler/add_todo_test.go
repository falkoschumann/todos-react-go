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
	name             string
	givenStoredTodos []data.Todo
	whenCommand      message.AddTodoCommand
	thenStatus       message.CommandStatus
	thenStoredTodos  []data.Todo
}{
	{
		name: "saves new todo.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		},
		whenCommand: message.AddTodoCommand{Title: "Buy Unicorn"},
		thenStatus:  message.MakeSuccess(),
		thenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
	},
	{
		name: "saves trimmed title.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		},
		whenCommand: message.AddTodoCommand{Title: "  Buy Unicorn   "},
		thenStatus:  message.MakeSuccess(),
		thenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
	},
	{
		name: "does nothing if title is empty.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		},
		whenCommand: message.AddTodoCommand{Title: "  "},
		thenStatus:  message.MakeSuccess(),
		thenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
		},
	},
}

func TestAddTodo(t *testing.T) {
	for _, tt := range addTodoTests {
		t.Run(tt.name, func(t *testing.T) {
			repo := provider.NewMemoryTodosRepository(tt.givenStoredTodos)
			addTodo := messagehandler.NewAddTodo(repo)

			gotStatus := addTodo(tt.whenCommand)
			gotTodos := repo.Load()

			if !cmp.Equal(tt.thenStatus, gotStatus) {
				t.Errorf("AddTodo(%v) = %v, want %v", tt.whenCommand, gotStatus, tt.thenStatus)
			}
			if diff := cmp.Diff(tt.thenStoredTodos, gotTodos); diff != "" {
				t.Errorf("AddTodo() stored todos mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
