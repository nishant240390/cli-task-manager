package main

import (
	cmd2 "cli-task-manager/cmd"
	"cli-task-manager/db"
	homedir "github.com/mitchellh/go-homedir"
	"path/filepath"
)
func main(){

	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
    if err != nil {
    	panic(err)
	}
	cmd2.RootCMD.Execute()
}
