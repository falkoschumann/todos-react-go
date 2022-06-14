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
	filepath string
}

func NewJsonTodosRepository(filepath string) *JsonTodosRepository {
	return &JsonTodosRepository{filepath: filepath}
}

func (r *JsonTodosRepository) Load() []domain.Todo {
	todos := []domain.Todo{}

	file, err := os.Open(r.filepath)
	if errors.Is(err, os.ErrNotExist) {
		return todos
	}
	if err != nil {
		log.Println("Opening todos file failed:", err)
		return todos
	}

	err = json.NewDecoder(file).Decode(&todos)
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
	dir := path.Dir(r.filepath)
	err := os.MkdirAll(dir, 0750)
	if err != nil {
		log.Println("Creating todos directory failed:", err)
		return
	}

	file, err := os.Create(r.filepath)
	if err != nil {
		log.Println("Creating todos file failed:", err)
		return
	}

	err = json.NewEncoder(file).Encode(todos)
	if err != nil {
		log.Println("Encoding todos file failed:", err)
		return
	}
}
