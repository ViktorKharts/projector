package cmd

import (
	"fmt"

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
	fd := models.FileData{
		SelectedProject: "",
		Projects:        []models.Project{},
	}

	pName := args[0]
	if len(args) == 0 {
		pName = "New Project"
	}

	p := models.Project{
		Name: pName,
		Tasks: []models.Task{
			{
				Value:      fmt.Sprintf("Project %s initiated.", args[0]),
				IsComplete: false,
			},
		},
	}

	fd, _ = storage.Read()
	fd.SelectedProject = p.Name
	fd.Projects = append(fd.Projects, p)

	if err := storage.Write(fd); err != nil {
		fmt.Printf("%s", err.Error())
	}
}
