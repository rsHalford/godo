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

	items, err := todo.Todos() // Get todo items from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	i, err := strconv.Atoi(args[0]) // Convert todo id argument to an integer.
	if err != nil {
		return fmt.Errorf("%v: %q %w", command, args[0], err)
	}

	if i > 0 && i <= len(items) { // Validate id argument.
		// Set the priority of the todo item to the opposite of it's current
		// boolean value. Then update the changes.
		if !items[i-1].Priority {
			items[i-1].Priority = true

			fmt.Printf("%s Setting priority...\n%s %q\n",
				c.StyleMust(Theme.Primary)("::"),
				c.StyleMust(Theme.Secondary)("-->"),
				items[i-1].Title)

			err = updateTodo(i, command, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		} else {
			items[i-1].Priority = false

			fmt.Printf("%s Removing priority...\n%s %q\n",
				c.StyleMust(Theme.Primary)("::"),
				c.StyleMust(Theme.Secondary)("-->"),
				items[i-1].Title)

			err = updateTodo(i, command, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		}
	} else {
		return fmt.Errorf("%v: %q %w", command, i, err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(priorityCmd)
}
