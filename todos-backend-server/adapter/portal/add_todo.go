package portal

import (
	"encoding/json"
	"log"
	"net/http"
	"todos_backend_server/domain"
)

func GetAddTodo(h domain.AddTodoCommandHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var command domain.AddTodoCommand
		err := json.NewDecoder(r.Body).Decode(&command)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		status := h(command)
		log.Println("Add todo", command, status)
		json.NewEncoder(w).Encode(status)
	}
}
