package cmd

import (
	"cli-task-manager/db"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var listCmd = &cobra.Command {
	Use : "list",
	Short : "List all task your task-list",
	Run : func(cmd *cobra.Command , args [] string) {
		tasks , err := db.AllTasks()
		if err != nil {
			fmt.Println("something went wrong ..", err.Error())
			os.Exit(1)
		}
		if len (tasks) == 0 {
			fmt.Println("you have no tasks ")
			os.Exit(1)
		}
		fmt.Println("you have following tasks")
		for _ , task := range tasks {
			fmt.Printf("%d : %s \n",task.Key,task.Value)
		}
	},
}
func init(){
	RootCMD.AddCommand(listCmd)
}
