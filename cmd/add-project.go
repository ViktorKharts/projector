package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
	"github.com/viktorkharts/projector/models"
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
	var fd models.FileData

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

	storage := os.Getenv("HOME") + "/projector-storage.json"
	f, err := os.ReadFile(storage)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			_, err := os.Create(storage)
			if err != nil {
				fmt.Printf("Projector Error: failed to create a storage file.\n%s\n", err.Error())
				os.Exit(1)
			}
			list(cmd, args)
		} else {
			fmt.Printf("Projector Error: failed to read storage file.\n%s\n", err.Error())
			os.Exit(1)
		}
	}

	if len(f) == 0 {
		fd.SelectedProject = p.Name
		fd.Projects = append(fd.Projects, p)

		bd, err := json.Marshal(fd)
		if err != nil {
			fmt.Printf("Projector Error: failed to Marshal Project data before save.\n%s\n", err.Error())
			os.Exit(1)
		}

		if err := os.WriteFile(storage, bd, 0666); err != nil {
			fmt.Printf("Projector Error: failed to write Project data into storage file.\n%s\n", err.Error())
			os.Exit(1)
		}
		return
	}

	if err = json.Unmarshal(f, &fd); err != nil {
		fmt.Printf("Projector Error: failed to parse file byte data into json.\n%s\n", err.Error())
		os.Exit(1)
	}

	fd.SelectedProject = p.Name
	fd.Projects = append(fd.Projects, p)

	bd, err := json.Marshal(fd)
	if err != nil {
		fmt.Printf("Projector Error: failed to Marshal Project data before save.\n%s\n", err.Error())
		os.Exit(1)
	}

	if err := os.WriteFile(storage, bd, 0666); err != nil {
		fmt.Printf("Projector Error: failed to write Project data into storage file.\n%s\n", err.Error())
		os.Exit(1)
	}
}
