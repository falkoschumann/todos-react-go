package portal

import (
	"encoding/json"
	"log"
	"net/http"
	"todos_backend_server/domain"
)

func GetSaveTodo(h domain.SaveTodoCommandHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var command domain.SaveTodoCommand
		err := json.NewDecoder(r.Body).Decode(&command)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		status := h(command)
		log.Println("Save todo", command, status)
		json.NewEncoder(w).Encode(status)
	}
}
