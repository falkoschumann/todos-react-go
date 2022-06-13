package provider

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path"
	"todos_backend_server/domain"
)

type JsonTodosRepository struct {
	file string
}

func NewJsonTodosRepository(file string) *JsonTodosRepository {
	return &JsonTodosRepository{
		file: file,
	}
}

func (r *JsonTodosRepository) Load() []domain.Todo {
	todos := []domain.Todo{}
	f, err := os.Open(r.file)
	if errors.Is(err, os.ErrNotExist) {
		return todos
	}
	if err != nil {
		log.Println("Opening todos file failed:", err)
		return todos
	}

	err = json.NewDecoder(f).Decode(&todos)
	if err != nil {
		log.Println("Decoding todos file failed:", err)
		return todos
	}
	if todos == nil {
		todos = []domain.Todo{}
	}
	return todos
}

func (r *JsonTodosRepository) Store(todos []domain.Todo) {
	dir := path.Dir(r.file)
	os.MkdirAll(dir, 0750)
	f, err := os.Create(r.file)
	if err != nil {
		log.Println("Creating todos file failed:", err)
		return
	}

	err = json.NewEncoder(f).Encode(todos)
	if err != nil {
		log.Println("Encoding todos file failed:", err)
		return
	}
}
