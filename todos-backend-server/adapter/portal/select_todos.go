package portal

import (
	"encoding/json"
	"log"
	"net/http"
	"todos_backend_server/domain"
)

func GetSelectTodos(h domain.SelectTodosQueryHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var query domain.SelectTodosQuery
		err := json.NewDecoder(r.Body).Decode(&query)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		result := h(query)
		log.Println("Select todos", query, result)
		json.NewEncoder(w).Encode(result)
	}
}
