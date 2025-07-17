package cmd

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/viktorkharts/projector/models"
	"github.com/viktorkharts/projector/storage"
)

func init() {
	rootCmd.AddCommand(listTasksCmd)
}

var listTasksCmd = &cobra.Command{
	Use:     "list-t",
	Short:   "List tasks in the current project.",
	Long:    "List all tasks in the currently selected project.",
	Aliases: []string{"lt", "lit", "lst", "lsit-t", "lits-t", "list-t"},
	Run:     listTasks,
}

func listTasks(cmd *cobra.Command, args []string) {
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

	fmt.Printf("%s tasks:\n", p.Name)

	for i, t := range p.Tasks {
		if !t.IsComplete {
			fmt.Printf("%d. %s\n", i+1, t.Value)
		}
	}
}
