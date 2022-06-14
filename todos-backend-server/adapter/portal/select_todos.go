package portal

import (
	"net/http"

	"todos_backend_server/domain"
)

func SelectTodos(h domain.SelectTodosQueryHandler) http.Handler {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) *httpError {
		var query domain.SelectTodosQuery
		result := h(query)
		if err := encodeJSON(w, result); err != nil {
			return err
		}

		return nil
	})
}
