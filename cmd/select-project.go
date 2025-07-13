package cmd

import (
	"fmt"

	"github.com/google/uuid"
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
	if len(args) == 0 {
		fmt.Printf("Projector Info: please provide a project name you want to select.\n")
		fmt.Printf("Projector Info: please make sure you have projects set up.\n")
		return
	}

	pName := args[0]
	s, _ := storage.Read()

	if _, ok := s.Projects[pName]; !ok {
		fmt.Printf("Projector Info: project '%s' doesn't exists.\n", pName)
		return
	}

	s.Projects[pName] = models.Project{
		Id:    uuid.NewString(),
		Name:  pName,
		Tasks: []models.Task{},
	}

	if err := storage.Write(s); err != nil {
		fmt.Printf("%s", err.Error())
	}

	fmt.Printf("Projector Info: '%s' project selected.\n", s.SelectedProject)
}
