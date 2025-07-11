package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

type FileData struct {
	Project string
	Tasks   []string
}

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List available projects.",
	Long:    "List available projects.",
	Aliases: []string{"l"},
	Run:     list,
}

func list(cmd *cobra.Command, args []string) {
	var fd FileData

	storage := os.Getenv("HOME") + "/projector-storage"
	f, err := os.ReadFile(storage)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			_, err := os.Create(storage)
			if err != nil {
				fmt.Printf("Projector Error: failed to create a storage file.\n%s\n", err.Error())
				os.Exit(1)
			}
			list(cmd, args)
		} else {
			fmt.Printf("Projector Error: failed to read storage file.\n%s\n", err.Error())
			os.Exit(1)
		}
	}

	if err = json.Unmarshal(f, &fd); err != nil {
		fmt.Printf("Projector Error: failed to parse file byte data into json.\n%s\n", err.Error())
		os.Exit(1)
	}

	fmt.Println(fd)
}
