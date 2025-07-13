package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viktorkharts/projector/models"
	"github.com/viktorkharts/projector/storage"
)

func init() {
	rootCmd.AddCommand(addTaskCmd)
}

var addTaskCmd = &cobra.Command{
	Use:     "add-t",
	Short:   "Creates a new task in the current project.",
	Long:    "Add a new task in the currently selected project.",
	Aliases: []string{"at", "addt", "add-task", "nt", "new-task", "newt"},
	Run:     addTask,
}

func addTask(cmd *cobra.Command, args []string) {
	fd, _ := storage.Read()

	if fd.SelectedProject == "" {
		fmt.Printf("Projector Info: please select a project to create a task for.\n")
		fmt.Printf("Projector Info: please make sure you have projects set up.\n")
		return
	}

	t := models.Task{
		Value:      args[0],
		IsComplete: false,
	}

	var p models.Project
	for _, sp := range fd.Projects {
		if sp.Name == fd.SelectedProject {
			p = sp
			break
		}
	}

	p.Tasks = append(p.Tasks, t)
	fd.Projects = append(fd.Projects, p)

	if err := storage.Write(fd); err != nil {
		fmt.Printf("%s", err.Error())
	}
}
