package message

type ToggleTodoCommand struct {
	Id int `json:"id"`
}

type ToggleTodoCommandHandler func(c ToggleTodoCommand) CommandStatus
