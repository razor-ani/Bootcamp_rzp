package cmd

import (
	"fmt"
	"os"

	"../db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("TODO list is empty")
			return
		}
		fmt.Printf("Total you have following %d tasks on your TODO list: \n", len(tasks))
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() { //without this cobra doen't know about any command you add
	// one can have as many init() functions as they wish all will run before main.go
	RootCmd.AddCommand(listCmd)
}
