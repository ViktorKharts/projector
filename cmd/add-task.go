package cmd

import (
	"fmt"

	"github.com/google/uuid"
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
	s, _ := storage.Read()

	if s.SelectedProject == "" {
		fmt.Printf("Projector Info: please select a project to create a task for.\n")
		fmt.Printf("Projector Info: please make sure you have projects set up.\n")
		return
	}

	t := models.Task{
		Id:         uuid.NewString(),
		Value:      args[0],
		IsComplete: false,
	}

	p, _ := s.Projects[s.SelectedProject]
	p.Tasks = append(p.Tasks, t)
	s.Projects[s.SelectedProject] = p

	if err := storage.Write(s); err != nil {
		fmt.Printf("%s", err.Error())
	}
}
