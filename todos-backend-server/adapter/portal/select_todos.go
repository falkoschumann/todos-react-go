package portal

import (
	"net/http"

	"todos_backend_server/domain/message"
)

func NewSelectTodos(handle message.SelectTodosQueryHandler) http.Handler {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) *httpError {
		var query message.SelectTodosQuery
		result := handle(query)
		if err := encodeJSON(w, result); err != nil {
			return err
		}

		return nil
	})
}
