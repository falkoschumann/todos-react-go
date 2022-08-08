package message

type ToggleAllCommand struct {
	Checked bool `json:"checked"`
}

type ToggleAllCommandHandler func(c ToggleAllCommand) CommandStatus
