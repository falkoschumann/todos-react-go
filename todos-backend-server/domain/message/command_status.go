package message

type CommandStatus struct {
	Success      bool   `json:"success"`
	ErrorMessage string `json:"errorMessage"`
}

func MakeSuccess() CommandStatus {
	return CommandStatus{Success: true}
}

func MakeFailure(errorMessage string) CommandStatus {
	return CommandStatus{Success: false, ErrorMessage: errorMessage}
}
