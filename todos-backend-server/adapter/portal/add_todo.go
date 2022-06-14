package portal

import (
	"net/http"

	"todos_backend_server/domain"
)

func AddTodo(h domain.AddTodoCommandHandler) http.Handler {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) *httpError {
		if err := isJSON(r); err != nil {
			return err
		}

		var command domain.AddTodoCommand
		if err := decodeJSON(r, &command); err != nil {
			return err
		}

		status := h(command)
		if err := encodeJSON(w, status); err != nil {
			return err
		}

		return nil
	})
}
