package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viktorkharts/projector/models"
	"github.com/viktorkharts/projector/storage"
)

func init() {
	rootCmd.AddCommand(addSelectProjectCmd)
}

var addSelectProjectCmd = &cobra.Command{
	Use:     "select-p",
	Short:   "Selects a project to work on.",
	Long:    "Selects a project as an acitve one to create/create/update/delete tasks in.",
	Aliases: []string{"s", "select", "sel", "sp", "selp"},
	Run:     selectProject,
}

func selectProject(cmd *cobra.Command, args []string) {
	fd := models.FileData{
		SelectedProject: "",
		Projects:        []models.Project{},
	}

	if len(args) == 0 {
		fmt.Printf("Projector Info: please provide a project name you want to select.\n")
		fmt.Printf("Projector Info: please make sure you have projects set up.\n")
		return
	}

	fd, _ = storage.Read()

	// TODO: prevent selecting non-existent project
	fd.SelectedProject = args[0]

	if err := storage.Write(fd); err != nil {
		fmt.Printf("%s", err.Error())
	}

	fmt.Printf("Projector Info: '%s' project selected.\n", fd.SelectedProject)
}
