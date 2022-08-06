package port

import "todos_backend_server/domain/data"

type TodosRepository interface {
	Load() []data.Todo
	Store(todos []data.Todo)
}
