package message

type ClearCompletedCommand struct{}

type ClearCompletedCommandHandler func(c ClearCompletedCommand) CommandStatus
