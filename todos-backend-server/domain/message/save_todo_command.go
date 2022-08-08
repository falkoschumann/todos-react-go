package message

type SaveTodoCommand struct {
	Id       int    `json:"id"`
	NewTitle string `json:"newTitle"`
}

type SaveTodoCommandHandler func(c SaveTodoCommand) CommandStatus
