package portal

import (
	"net/http"

	"todos_backend_server/domain/message"
)

func NewAddTodo(handle message.AddTodoCommandHandler) http.Handler {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) *httpError {
		if err := isJSON(r); err != nil {
			return err
		}

		var command message.AddTodoCommand
		if err := decodeJSON(r, &command); err != nil {
			return err
		}

		status := handle(command)
		if err := encodeJSON(w, status); err != nil {
			return err
		}

		return nil
	})
}
