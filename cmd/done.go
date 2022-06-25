/*
Done will label a todo as done or active.

Usage:

	godo done [flags]

Aliases:

	done, do, d

Flags:

	-h, --help   help for done
*/
package cmd

import (
	"fmt"
	"strconv"

	c "github.com/jwalton/gchalk"
	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command.
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"do", "d"},
	Short:   "Toggle todo as done",
	Long:    `Done will label a todo as done or active.`,
	RunE:    doneRun,
}

func doneRun(cmd *cobra.Command, args []string) error {
	var command string = "done"

	items, err := todo.Todos() // Get todo items from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	i, err := strconv.Atoi(args[0]) // Convert todo id argument to an integer.
	if err != nil {
		return fmt.Errorf("%v: %q %w", command, args[0], err)
	}

	if i > 0 && i <= len(items) { // Validate id argument.
		// Set the status of the todo item to the opposite of it's current
		// boolean value. Then update the changes.
		if !items[i-1].Status {
			items[i-1].Status = true

			fmt.Printf("%s Marked done...\n%s %q\n",
				c.StyleMust(Theme.Primary)("::"),
				c.StyleMust(Theme.Secondary)("-->"),
				items[i-1].Title)

			err = updateTodo(i, command, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		} else {
			items[i-1].Status = false

			fmt.Printf("%s Marked active...\n%s %q\n",
				c.StyleMust(Theme.Primary)("::"),
				c.StyleMust(Theme.Secondary)("-->"),
				items[i-1].Title)

			err = updateTodo(i, command, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		}
	} else {
		return fmt.Errorf("%v: %q %w", i, command, err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
