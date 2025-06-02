/*
Copyright © 2025 Chris Hartley
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/chartleyit/todo-cli/dataStore"
	"github.com/chartleyit/todo-cli/models"
	"github.com/spf13/cobra"
)

var ds dataStore.Data
var parentId int

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add subcommand adds a new item to your todo list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		// TODO look at viper config accessing

		var todo models.TodoItem

		ds = &dataStore.CSVData{FilePath: file}
		todos, err := ds.Load()
		if err != nil {
			fmt.Println(err)
		}

		todo.Id = len(todos) + 1
		todo.ParentId = parentId
		todo.Task = args[0]
		todo.CreatedAt = time.Now()
		todo.Status = models.Status(0)

		todos = append(todos, &todo)

		ds.Save(todos)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addCmd.Flags().IntVarP(&parentId, "parent", "p", 0, "Parent task id")

}

func initDataStore() {
	switch format {
	case "json":
		fmt.Errorf("json not implemented yet")
	default:
		ds = &dataStore.CSVData{FilePath: file}
	}
}
