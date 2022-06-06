package portal

import (
	"encoding/json"
	"log"
	"net/http"
	"todos_backend_server/domain"
)

func GetToggleAll(h domain.ToggleAllCommandHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var command domain.ToggleAllCommand
		err := json.NewDecoder(r.Body).Decode(&command)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		status := h(command)
		log.Println("Toggle all", command, status)
		json.NewEncoder(w).Encode(status)
	}
}
