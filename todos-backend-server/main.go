package main

import (
	"fmt"
	"log"
	"net/http"

	"todos_backend_server/adapter/portal"
	"todos_backend_server/adapter/provider"
	"todos_backend_server/domain/messagehandler"
)

func main() {
	host, port, dataDir := portal.ParseFlags()
	initAPI(dataDir)
	initClient()
	runServer(host, port)
}

func initAPI(dataDir string) {
	repo := provider.NewJsonTodosRepository(dataDir)
	http.Handle("/api/todos/add-todo", portal.NewAddTodo(messagehandler.NewAddTodo(repo)))
	http.Handle("/api/todos/clear-completed", portal.NewClearCompleted(messagehandler.NewClearCompleted(repo)))
	http.Handle("/api/todos/destroy-todo", portal.NewDestroyTodo(messagehandler.NewDestroyTodo(repo)))
	http.Handle("/api/todos/save-todo", portal.NewSaveTodo(messagehandler.NewSaveTodo(repo)))
	http.Handle("/api/todos/select-todos", portal.NewSelectTodos(messagehandler.NewSelectTodos(repo)))
	http.Handle("/api/todos/toggle-all", portal.NewToggleAll(messagehandler.NewToggleAll(repo)))
	http.Handle("/api/todos/toggle-todo", portal.NewToggleTodo(messagehandler.NewToggleTodo(repo)))
}

func initClient() {
	handler := portal.NewSPA()
	http.Handle("/", handler)
}

func runServer(host string, port uint) {
	log.Println("Server listening on port", port)
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
