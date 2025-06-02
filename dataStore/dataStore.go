package dataStore

import (
	"github.com/chartleyit/todo-cli/models"
)

type DataHandler interface {
	Load() ([]*models.TodoItem, error)
	Save([]*models.TodoItem) error
}
