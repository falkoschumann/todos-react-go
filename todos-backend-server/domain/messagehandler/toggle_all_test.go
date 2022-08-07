package messagehandler_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain/data"
	"todos_backend_server/domain/message"
	"todos_backend_server/domain/messagehandler"
)

var toggleAllTests = []struct {
	name             string
	givenStoredTodos []data.Todo
	whenCommand      message.ToggleAllCommand
	thenStatus       message.CommandStatus
	thenStoredTodos  []data.Todo
}{
	{
		name: "set all todos completed.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
		whenCommand: message.ToggleAllCommand{Checked: true},
		thenStatus:  message.MakeSuccess(),
		thenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: true},
		},
	},
	{
		name: "set all todos active.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
		whenCommand: message.ToggleAllCommand{Checked: false},
		thenStatus:  message.MakeSuccess(),
		thenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: false},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
	},
}

func TestToggleAll(t *testing.T) {
	for _, tt := range toggleAllTests {
		t.Run(tt.name, func(t *testing.T) {
			repo := provider.NewMemoryTodosRepository(tt.givenStoredTodos)
			toggleAll := messagehandler.NewToggleAll(repo)

			gotStatus := toggleAll(tt.whenCommand)
			gotStoredTodos := repo.Load()

			if !cmp.Equal(tt.thenStatus, gotStatus) {
				t.Errorf("ToggleAll(%v) = %v, want %v", tt.whenCommand, gotStatus, tt.thenStatus)
			}
			if diff := cmp.Diff(tt.thenStoredTodos, gotStoredTodos); diff != "" {
				t.Errorf("ToggleAll() stored todos mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
