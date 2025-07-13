package cmd

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/spf13/cobra"

	"github.com/viktorkharts/projector/models"
	"github.com/viktorkharts/projector/storage"
)

func init() {
	rootCmd.AddCommand(addProjectCmd)
}

var addProjectCmd = &cobra.Command{
	Use:     "add-p",
	Short:   "Creates a new project.",
	Long:    "Add a new project to the list of ongoing initiatives.",
	Aliases: []string{"ap", "addp", "add-project", "np", "new-project", "newp"},
	Run:     addProject,
}

func addProject(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Printf("Projector Info: please provide a name for a new project.\n")
		return
	}

	pName := args[0]

	p := models.Project{
		Id:   uuid.NewString(),
		Name: pName,
		Tasks: []models.Task{
			{
				Id:         uuid.NewString(),
				Value:      fmt.Sprintf("Project %s initiated.", args[0]),
				IsComplete: false,
			},
		},
	}

	s, _ := storage.Read()

	if _, ok := s.Projects[pName]; ok {
		fmt.Printf("Projector Info: '%s' project already exists.\n", pName)
		return
	}

	s.SelectedProject = p.Name
	s.Projects[pName] = p

	if err := storage.Write(s); err != nil {
		fmt.Printf("%s", err.Error())
	}
}
