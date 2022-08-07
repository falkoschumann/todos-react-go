package messagehandler_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/messagehandler"
)

var saveTodoTests = []struct {
	name             string
	givenStoredTodos []data.Todo
	whenCommand      message.SaveTodoCommand
	thenStatus       message.CommandStatus
	thenStoredTodos  []data.Todo
}{
	{
		name: "changes todos title.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
		whenCommand: message.SaveTodoCommand{Id: 1, NewTitle: "Taste TypeScript"},
		thenStatus:  message.MakeSuccess(),
		thenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste TypeScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
	},
	{
		name: "trims title.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
		whenCommand: message.SaveTodoCommand{Id: 1, NewTitle: "   Taste TypeScript  "},
		thenStatus:  message.MakeSuccess(),
		thenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste TypeScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
	},
	{
		name: "destroys todo if title is empty.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
		whenCommand: message.SaveTodoCommand{Id: 1, NewTitle: ""},
		thenStatus:  message.MakeSuccess(),
		thenStoredTodos: []data.Todo{
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
	},
}

func TestSaveTodo(t *testing.T) {
	for _, tt := range saveTodoTests {
		t.Run(tt.name, func(t *testing.T) {
			repo := provider.NewMemoryTodosRepository(tt.givenStoredTodos)
			saveTodo := messagehandler.NewSaveTodo(repo)

			gotStatus := saveTodo(tt.whenCommand)
			gotStoredTodos := repo.Load()

			if !cmp.Equal(tt.thenStatus, gotStatus) {
				t.Errorf("ClearCompleted(%v) = %v, want %v", tt.whenCommand, gotStatus, tt.thenStatus)
			}
			if diff := cmp.Diff(tt.thenStoredTodos, gotStoredTodos); diff != "" {
				t.Errorf("SaveTodo() stored todos mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
