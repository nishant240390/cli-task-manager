package cmd

import (
	"github.com/spf13/cobra"
)

var RootCMD = &cobra.Command {
	Use : "task",
	Short : "A cli task manager",
}
