package domain

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type TodosRepository interface {
	Load() []Todo
	Store(todos []Todo)
}
