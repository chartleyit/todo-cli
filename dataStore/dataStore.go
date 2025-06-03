package dataStore

import (
	"fmt"
	"strings"

	"github.com/chartleyit/todo-cli/models"
)

type DataHandler interface {
	Load() ([]*models.TodoItem, error)
	Save([]*models.TodoItem) error
}

func New(path string) DataHandler {
	var format string

	// parse file name and determine proper handler
	e := strings.Split(path, ".")
	format = e[len(e)-1]

	switch format {
	case "json":
		fmt.Println("Not yet implemented")
		return nil
	case "csv":
		return &CSVData{FilePath: path}
	case "sqlite":
		fmt.Println("Not yet implemented")
		return nil
	default:
		fmt.Printf("unknown format, %s\n", format)
		return nil
	}
}
