package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"todos_backend_server/adapter/portal"
	"todos_backend_server/adapter/provider"
	"todos_backend_server/messagehandler"
)

func main() {
	host, port := parseFlags()
	createRouter()
	runServer(host, port)
}

func parseFlags() (host string, port uint) {
	flag.StringVar(&host, "host", "", "the server listen on this `host` (default all)")
	flag.UintVar(&port, "port", 8080, "the server listen on this `port`")
	flag.Parse()
	return
}

func createRouter() {
	repo := provider.NewJsonTodosRepository("./data/todos.json")
	http.Handle("/", portal.NewSpaHandler())
	http.HandleFunc("/api/todos/add-todo", portal.GetAddTodo(messagehandler.GetAddTodo(repo)))
	http.HandleFunc("/api/todos/clear-completed", portal.GetClearCompleted(messagehandler.GetClearCompleted(repo)))
	http.HandleFunc("/api/todos/destroy-todo", portal.GetDestroyTodo(messagehandler.GetDestroyTodo(repo)))
	http.HandleFunc("/api/todos/save-todo", portal.GetSaveTodo(messagehandler.GetSaveTodo(repo)))
	http.HandleFunc("/api/todos/select-todos", portal.GetSelectTodos(messagehandler.GetSelectTodos(repo)))
	http.HandleFunc("/api/todos/toggle-all", portal.GetToggleAll(messagehandler.GetToggleAll(repo)))
	http.HandleFunc("/api/todos/toggle-todo", portal.GetToggleTodo(messagehandler.GetToggleTodo(repo)))
}

func runServer(host string, port uint) {
	log.Println("Server listening on port", port)
	addr := fmt.Sprintf("%s:%d", host, port)
	log.Fatal(http.ListenAndServe(addr, nil))
}
