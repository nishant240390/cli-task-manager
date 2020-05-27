package cmd

import (
	"cli-task-manager/db"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var addCmd = &cobra.Command {
	Use : "add",
	Short : "Adds a task your task-list",
	Run : func(cmd *cobra.Command , args [] string) {
		task := strings.Join(args, " ");

		_,err := db.CreateTask(task)
		if err != nil {
			fmt.Println("something went wrong", err.Error())
			return
		}
		fmt.Printf("Added \"%s\" in your list\n",task)
	},
}
func init(){
	RootCMD.AddCommand(addCmd)
}