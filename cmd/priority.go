/*
Priority will toggle a todo as being labelled a priority or not.

Usage:

	godo priority [flags]

Aliases:

	priority, pri, p

Flags:

	-h, --help   help for priority
*/
package cmd

import (
	"fmt"
	"strconv"

	c "github.com/jwalton/gchalk"
	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

// priorityCmd represents the priority command.
var priorityCmd = &cobra.Command{
	Use:     "priority",
	Aliases: []string{"pri", "p"},
	Short:   "Label a todo as a priority",
	Long:    `Priority will toggle a todo as being labelled a priority or not.`,
	RunE:    priorityRun,
}

func priorityRun(cmd *cobra.Command, args []string) error {
	var command string = "priority"

	todos, err := todo.Todos() // Get todos from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	// Convert todo position argument to an integer.
	p, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%v: %q %w", command, args[0], err)
	}

	if p > 0 && p <= len(todos) { // Validate position argument.
		t := todos[p-1]

		// Set the priority of the todo to the opposite of it's current boolean
		// value. Then update the changes.
		if !t.Priority {
			t.Priority = true

			fmt.Printf("%s Setting priority...\n%s %q\n",
				c.StyleMust(Theme.Primary)("::"),
				c.StyleMust(Theme.Secondary)("-->"), t.Title)

			err = updateTodo(p, command, todos)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		} else {
			t.Priority = false

			fmt.Printf("%s Removing priority...\n%s %q\n",
				c.StyleMust(Theme.Primary)("::"),
				c.StyleMust(Theme.Secondary)("-->"), t.Title)

			err = updateTodo(p, command, todos)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		}
	} else {
		return fmt.Errorf("%v: %q %w", command, p, err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(priorityCmd)
}
