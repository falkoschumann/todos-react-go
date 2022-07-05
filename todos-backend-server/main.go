package main

import (
	"fmt"
	"log"
	"net/http"

	"todos_backend_server/adapter/portal"
	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain"
	"todos_backend_server/messagehandler"
)

func main() {
	host, port, data := portal.ParseFlags()
	repo := provider.NewJsonTodosRepository(data)
	createBackendAPIRouter(repo)
	createFrontendRouter()
	runServer(host, port)
}

func createBackendAPIRouter(r domain.TodosRepository) {
	http.Handle("/api/todos/add-todo", portal.AddTodo(messagehandler.AddTodo(r)))
	http.Handle("/api/todos/clear-completed", portal.ClearCompleted(messagehandler.ClearCompleted(r)))
	http.Handle("/api/todos/destroy-todo", portal.DestroyTodo(messagehandler.DestroyTodo(r)))
	http.Handle("/api/todos/save-todo", portal.SaveTodo(messagehandler.SaveTodo(r)))
	http.Handle("/api/todos/select-todos", portal.SelectTodos(messagehandler.SelectTodos(r)))
	http.Handle("/api/todos/toggle-all", portal.ToggleAll(messagehandler.ToggleAll(r)))
	http.Handle("/api/todos/toggle-todo", portal.ToggleTodo(messagehandler.ToggleTodo(r)))
}

func createFrontendRouter() {
	handler := portal.NewSpaHandler()
	handler.StaticPath = "/static"
	http.Handle("/", handler)
}

func runServer(host string, port uint) {
	log.Println("Server listening on port", port)
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
