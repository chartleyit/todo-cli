/*
Copyright © 2025 Chris Hartley
*/
package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chartleyit/todo-cli/dataStore"
	"github.com/chartleyit/todo-cli/models"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a todo item",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("complete called")

		ds = dataStore.New(file)
		todos, err := ds.Load()
		if err != nil {
			fmt.Println(err)
		}

		completeId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Failed to convert id string to int")
		}

		for i, todo := range todos {
			if todo.Id == completeId {
				todos[i].Status = models.Done
				todos[i].Done = time.Now()
			}
		}

		ds.Save(todos)
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
