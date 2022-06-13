package portal

import (
	"net/http"
	"todos_backend_server/domain"
)

func ToggleAll(h domain.ToggleAllCommandHandler) http.Handler {
	return httpHandler(func(w http.ResponseWriter, r *http.Request) *httpError {
		if err := isJSON(r); err != nil {
			return err
		}

		var command domain.ToggleAllCommand
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
