package message

import "todos_backend_server/domain/data"

type SelectTodosQuery struct{}

type SelectTodosQueryResult struct {
	Todos []data.Todo `json:"todos"`
}

type SelectTodosQueryHandler func(q SelectTodosQuery) SelectTodosQueryResult
