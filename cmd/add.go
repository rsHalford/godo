/*
Add will create a new todo to the list.

Usage:

	godo add [flags]

Aliases:

	add, a

Flags:

	-h, --help         help for add
	-p, --priority     assign priority to the todo
	-T, --tag string   add a tag to the todo
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command.
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add a new todo",
	Long:    `Add will create a new todo to the list.`,
	RunE:    addRun,
}

var (
	priority bool
	tagStr   string
)

func addRun(cmd *cobra.Command, args []string) error {
	var command string = "add"

	todos, err := todo.Todos() // Get todos from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	// Each argument given to the add command will be assigned to an individual todo.
	for _, x := range args {
		var t time.Time = time.Now().UTC()

		todo := todo.Todo{Title: x, CreatedAt: t, UpdatedAt: t}

		// Mark the todo priority as true, if the --priority flag is provided.
		// Currently this flag affects all new todos being declared.
		todo.Prioritise(priority)

		// Add the given tag to the todo, if the --tag flag is provided.
		todo.Tagging(tagStr)

		todos = append(todos, todo) // Add the new complete todo to todos.
	}

	// Save the new todo(s) locally.
	var filename string

	// Pass the filename of the local todo store to the filename variable.
	filename, err = todo.LocalTodos()
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	// Using SaveLocal to add the new todo(s) to the local JSON store.
	err = todo.SaveLocal(filename, todos)
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)

	// The --priority boolean flag is used to set the priority object of the new todo(s) as true.
	addCmd.Flags().BoolVarP(&priority, "priority", "p", false, "assign priority to the todo")
	// The --tag string flag is used to add a tag to the new todo(s).
	addCmd.Flags().StringVarP(&tagStr, "tag", "T", "", "add a tag to the todo")
}
