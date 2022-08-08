package message

type AddTodoCommand struct {
	Title string `json:"title"`
}

type AddTodoCommandHandler func(c AddTodoCommand) CommandStatus
