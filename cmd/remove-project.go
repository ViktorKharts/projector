package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/viktorkharts/projector/storage"
)

func init() {
	rootCmd.AddCommand(addRemoveProjectCmd)
}

var addRemoveProjectCmd = &cobra.Command{
	Use:     "remove-p",
	Short:   "Removes a project from the library.",
	Long:    "Removes a project and all the associated tasks to it.",
	Aliases: []string{"rep", "remove", "rm"},
	Run:     removeProject,
}

func removeProject(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Printf("Projector Info: please provide a project you want to remove.\n")
		return
	}

	pName := args[0]
	s, _ := storage.Read()

	if _, ok := s.Projects[pName]; !ok {
		fmt.Printf("Projector Info: project '%s' doesn't exists.\n", pName)
		return
	}

	delete(s.Projects, pName)

	if err := storage.Write(s); err != nil {
		fmt.Printf("%s", err.Error())
	}

	fmt.Printf("Projector Info: '%s' project was removed.\n", pName)
}
