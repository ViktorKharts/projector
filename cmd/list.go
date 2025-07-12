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
	fd, err := storage.ReadStorage()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if reflect.DeepEqual(fd, models.FileData{}) {
		fmt.Printf("Projector Info: you have no projects created.\n")
		return
	}

	if fd.SelectedProject == "" {
		fmt.Printf("Projector Info: please select a project to work on.\n")
		return
	}

	for _, proj := range fd.Projects {
		if proj.Name == fd.SelectedProject {
			project := proj

			for _, task := range project.Tasks {
				if !task.IsComplete {
					fmt.Println(task.Value)
				}
			}

			return
		}
	}
}
