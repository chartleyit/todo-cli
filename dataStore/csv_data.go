package dataStore

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"time"

	"github.com/chartleyit/todo-cli/models"
)

type CSVData struct {
	FilePath string
}

func (c *CSVData) Load() ([]*models.TodoItem, error) {
	fmt.Println("Load CSV Data from", c.FilePath)

	f, err := readFile(c)
	if err != nil {
		return nil, err
	}

	defer closeFile(f)

	todoItems, err := decodeCSVtodo(f)
	if err != nil {
		return nil, fmt.Errorf("error decoding CSV; %w", err)
	}

	return todoItems, nil
}

func (c *CSVData) Save(items []*models.TodoItem) error {
	fmt.Println("Save CSV Data to", c)
	f, err := writeFile(c)
	if err != nil {
		return err
	}

	defer closeFile(f)

	// ! ERROR writer doesn't cleanly write new file
	writer := csv.NewWriter(f)
	err = writer.WriteAll(encodeCSVtodo(items))
	if err != nil {
		return fmt.Errorf("failed writing file; %w", err)
	}

	return nil
}

func readFile(c *CSVData) (*os.File, error) {
	f, err := os.OpenFile(c.FilePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func writeFile(c *CSVData) (*os.File, error) {
	f, err := os.OpenFile(c.FilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}

func decodeCSVtodo(f *os.File) ([]*models.TodoItem, error) {
	var todos []*models.TodoItem

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to readall from file reader; %w", err)
	}

	for i, record := range records {
		// TODO enable when header line FEAT-001
		// if i == 0 {
		// 	continue
		// }

		if len(record) < 7 {
			return nil, fmt.Errorf("incorrect record leng; %d", i)
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, fmt.Errorf("failed to convert id to int, %s", record[0])
		}

		pid, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, fmt.Errorf("failed to convert parent id to int, %s", record[1])
		}

		// TODO convert string to list of ints
		// cid, err := strconv.Atoi(record[2])
		// if err != nil {
		// 	return nil, fmt.Errorf("can not convert children ids to []int, %s", record[2])
		// }

		createdAt, err := time.Parse(time.RFC3339, record[4])
		if err != nil {
			return nil, fmt.Errorf("unable to parse created time, %s", record[4])
		}

		due, err := time.Parse(time.RFC3339, record[5])
		if err != nil {
			return nil, fmt.Errorf("unable to parse due time, %s", record[5])
		}

		done, err := time.Parse(time.RFC3339, record[6])
		if err != nil {
			return nil, fmt.Errorf("unable to parse due time, %s", record[6])
		}

		status, err := strconv.Atoi(record[7])
		if err != nil {
			fmt.Println(record)
			return nil, fmt.Errorf("failed to convert status to int, %s", record[7])
		}

		todo := &models.TodoItem{
			Id:        id,
			ParentId:  pid,
			Task:      record[3],
			CreatedAt: createdAt,
			Due:       due,
			Status:    models.Status(status),
			Done:      done,
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func encodeCSVtodo(todos []*models.TodoItem) [][]string {
	var records [][]string
	// TODO put a header line in [][]string FEAT-001

	for _, todo := range todos {
		record := []string{
			strconv.Itoa(todo.Id),
			strconv.Itoa(todo.ParentId),
			// TODO convert list to string
			"[]", // place holder for children string
			todo.Task,
			todo.CreatedAt.Format(time.RFC3339),
			todo.Due.Format(time.RFC3339),
			todo.Done.Format(time.RFC3339),
			strconv.Itoa(int(todo.Status)),
		}
		records = append(records, record)
	}

	return records
}
