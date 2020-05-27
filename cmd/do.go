package cmd

import (
	"cli-task-manager/db"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var doCmd = &cobra.Command {
	Use : "do",
	Short : "Do a task from  your task-list",
	Run : func(cmd *cobra.Command , args [] string) {
		var ids []int
		for _,i := range args {
			 id,err := strconv.Atoi(i)
			 if err != nil {
				panic("error while marking task complete")
			}
			ids = append(ids, id)
		}
		for _,id := range ids {
		   err := db.DeleteTask(id)
		   if err != nil {
				fmt.Printf("failed to complete task %d \n",id)
		   } else  {
			    fmt.Printf("successfully marked compelted task %d \n",id)
		   }
		}
	},
}
func init(){
	RootCMD.AddCommand(doCmd)
}