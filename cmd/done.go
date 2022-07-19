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

		// Set the boolean value of done to the opposite of it's current value.
		// Then update the changes.
		if !todos[p-1].Done {
			todos[p-1].Done = true

			fmt.Printf("%s Marked done...\n%s %q\n",
				c.StyleMust(Theme.Primary)("::"),
				c.StyleMust(Theme.Secondary)("-->"), todos[p-1].Title)

			err = updateTodo(p, command, todos)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		} else {
			todos[p-1].Done = false

			fmt.Printf("%s Marked active...\n%s %q\n",
				c.StyleMust(Theme.Primary)("::"),
				c.StyleMust(Theme.Secondary)("-->"), todos[p-1].Title)

			err = updateTodo(p, command, todos)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		}
	} else {
		return fmt.Errorf("(%v) todo does not exist at position \"%d\"\n", command, p)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
