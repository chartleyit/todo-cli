/*
Copyright © 2025 Chris Hartley
*/
package cmd

import (
	"fmt"

	"github.com/chartleyit/todo-cli/dataStore"
	"github.com/chartleyit/todo-cli/ui"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print todo list to terminal",
	Long:  `More detail with example`,
	Run: func(cmd *cobra.Command, args []string) {
		ds = &dataStore.CSVData{FilePath: file}
		todos, err := ds.Load()
		if err != nil {
			fmt.Println(err)
		}

		t := ui.New()
		t.AddHeader(
			"ID", "Task", "Status", "Created", "Due",
		)
		for _, todo := range todos {
			t.AddLine(
				todo.Id,
				todo.Task,
				todo.Status.String(),
				todo.CreatedAt.Format("2006-01-02 03:04"),
				todo.Due.Format("2006-01-02 03:04"),
			)
		}
		t.Print()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
