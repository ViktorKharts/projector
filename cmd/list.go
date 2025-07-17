package cmd

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/viktorkharts/projector/models"
	"github.com/viktorkharts/projector/storage"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List available projects.",
	Long:    "List available projects.",
	Aliases: []string{"l", "li", "ls", "lis", "lsi", "lsit", "lits", "list"},
	Run:     list,
}

func list(cmd *cobra.Command, args []string) {
	s, err := storage.Read()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if reflect.DeepEqual(s, models.Storage{}) {
		fmt.Printf("Projector Info: you have no projects created.\n")
		return
	}

	if s.SelectedProject == "" {
		fmt.Printf("Projector Info: please select a project to work on.\n")
		return
	}

	p, _ := s.Projects[s.SelectedProject]

	if len(p.Tasks) == 0 {
		fmt.Printf("Projector Info: current project contains no tasks.\n")
		return
	}

	for _, t := range p.Tasks {
		if !t.IsComplete {
			fmt.Println(t.Value)
		}
	}
}
