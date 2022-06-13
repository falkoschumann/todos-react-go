package portal

import (
	"encoding/json"
	"log"
	"net/http"
)

type httpError struct {
	Message string
	Code    int
	Error   error
}

type httpHandler func(http.ResponseWriter, *http.Request) *httpError

func (h httpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		log.Println("http:", err)
		http.Error(w, err.Message, err.Code)
	}
}

func isJSON(r *http.Request) *httpError {
	ct := r.Header.Get("Content-Type")
	if ct != "application/json" {
		return &httpError{
			Message: "Content type must be application/json.",
			Code:    http.StatusUnsupportedMediaType,
		}
	}

	return nil
}

func decodeJSON(r *http.Request, v any) *httpError {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return &httpError{
			Message: "Can not decode JSON from request body.",
			Code:    http.StatusBadRequest,
			Error:   err,
		}
	}

	return nil
}

func encodeJSON(w http.ResponseWriter, v any) *httpError {
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return &httpError{
			Message: "Can not encode JSON for response body.",
			Code:    http.StatusInternalServerError,
			Error:   err,
		}
	}

	return nil
}
