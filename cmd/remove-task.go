package cmd

import (
	"fmt"
	"slices"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/viktorkharts/projector/storage"
)

func init() {
	rootCmd.AddCommand(addRemoveTaskCmd)
}

var addRemoveTaskCmd = &cobra.Command{
	Use:     "remove-t",
	Short:   "Removes a task from a project.",
	Long:    "Removes a task from a project it is defined in.",
	Aliases: []string{"ret", "remove-t", "rm-t"},
	Run:     removeTask,
}

func removeTask(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Printf("Projector Info: please provide a task number you want to remove.\n")
		return
	}

	if len(args[0]) > 1 {
		fmt.Printf("Projector Error: failed to convert '%s' string input to integer.\n", args[0])
		return
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Projector Error: failed to convert '%s' string input to integer.\n", args[0])
		return
	}

	s, err := storage.Read()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	p := s.Projects[s.SelectedProject]
	t := p.Tasks

	if len(t) == 0 {
		fmt.Printf("Projector Info: no tasks in the current project.\n")
		return
	}

	t = slices.Delete(t, i-1, i)
	p.OverWriteTasks(t)
	s.Projects[s.SelectedProject] = p

	if err := storage.Write(s); err != nil {
		fmt.Printf("%s", err.Error())
	}

	fmt.Printf("Projector Info: task number '%d' was removed.\n", i)
}
