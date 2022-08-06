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
	name                string
	givenStoredTodos    []data.Todo
	command             message.ToggleAllCommand
	status              message.CommandStatus
	expectedStoredTodos []data.Todo
}{
	{
		name: "set all todos completed.",
		givenStoredTodos: []data.Todo{
			{Id: 1, Title: "Taste JavaScript", Completed: true},
			{Id: 2, Title: "Buy Unicorn", Completed: false},
		},
		command: message.ToggleAllCommand{Checked: true},
		status:  message.MakeSuccess(),
		expectedStoredTodos: []data.Todo{
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
		command: message.ToggleAllCommand{Checked: false},
		status:  message.MakeSuccess(),
		expectedStoredTodos: []data.Todo{
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

			gotStatus := toggleAll(tt.command)
			gotStoredTodos := repo.Load()

			if !cmp.Equal(tt.status, gotStatus) {
				t.Errorf("ToggleAll(%v) = %v, want %v", tt.command, gotStatus, tt.status)
			}
			if diff := cmp.Diff(tt.expectedStoredTodos, gotStoredTodos); diff != "" {
				t.Errorf("ToggleAll() stored todos mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
