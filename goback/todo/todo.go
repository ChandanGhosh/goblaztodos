package todo

import (
	"errors"
	"sync"
	"github.com/rs/xid"
)

// Todo ...
type Todo struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	Done    bool   `json:"done"`
}

var (
	todoList []Todo
	mtx      sync.RWMutex
	once     sync.Once
)

func init() {

	once.Do(initialize)
}

func initialize() {

	todoList = []Todo{}
}

// GetAllTodo ...
func GetAllTodo() []Todo {
	return todoList
}

// AddTodo ...
func AddTodo(message string) string {

	t := Todo{
		ID:      xid.New().String(),
		Message: message,
		Done:    false,
	}
	mtx.Lock()
	todoList = append(todoList, t)
	mtx.Unlock()
	return t.ID
}

// UpdateToDo ...
func UpdateToDo(todo Todo) error {
	for i, t := range todoList {
		if t.ID == todo.ID {
			mtx.Lock()
			todoList[i] = todo

			mtx.Unlock()
			return nil
		}
	}
	return errors.New("Could not find any todo with the id: " + todo.ID)

}
