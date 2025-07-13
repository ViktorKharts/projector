package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// 1. remove project
// 2. remove task
// 3. list projects

var rootCmd = &cobra.Command{
	Use:   "projector",
	Short: "Projector is a todo cli app",
	Long:  "projector is a an enhanced todo list application that allows you to set, track, and complete the tasks you are working on.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute Projector '%s'\n", err)
		os.Exit(1)
	}
}
