package cmd

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
	"github.com/viktorkharts/projector/models"
	"github.com/viktorkharts/projector/storage"
)

func init() {
	rootCmd.AddCommand(listProjectsCmd)
}

var listProjectsCmd = &cobra.Command{
	Use:     "list-p",
	Short:   "List available projects.",
	Long:    "List available projects.",
	Aliases: []string{"lp", "lip", "lsp", "lsip-p", "lits-p", "list-p"},
	Run:     listProjects,
}

func listProjects(cmd *cobra.Command, args []string) {
	s, err := storage.Read()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if reflect.DeepEqual(s, models.Storage{}) {
		fmt.Printf("Projector Info: you have no projects created.\n")
		return
	}

	if len(s.Projects) == 0 {
		fmt.Printf("Projector Info: you have no projects created.\n")
		return
	}

	fmt.Printf("Projects:\n")

	c := 1
	for _, p := range s.Projects {
		fmt.Printf("%d. %s\n", c, p.Name)
		c += 1
	}
}
