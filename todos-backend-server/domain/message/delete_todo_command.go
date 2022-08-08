package message

type DestroyTodoCommand struct {
	Id int `json:"id"`
}

type DestroyTodoCommandHandler func(c DestroyTodoCommand) CommandStatus
