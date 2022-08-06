package message

type AddTodoCommand struct {
	Title string `json:"title"`
}

type AddTodoCommandHandler func(c AddTodoCommand) CommandStatus

type ClearCompletedCommand struct{}

type ClearCompletedCommandHandler func(c ClearCompletedCommand) CommandStatus

type DestroyTodoCommand struct {
	Id int `json:"id"`
}

type DestroyTodoCommandHandler func(c DestroyTodoCommand) CommandStatus

type SaveTodoCommand struct {
	Id       int    `json:"id"`
	NewTitle string `json:"newTitle"`
}

type SaveTodoCommandHandler func(c SaveTodoCommand) CommandStatus

type ToggleAllCommand struct {
	Checked bool `json:"checked"`
}

type ToggleAllCommandHandler func(c ToggleAllCommand) CommandStatus

type ToggleTodoCommand struct {
	Id int `json:"id"`
}

type ToggleTodoCommandHandler func(c ToggleTodoCommand) CommandStatus
