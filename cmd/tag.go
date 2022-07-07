/*
Tag is used to add or replace a tag for a todo.

Usage:

	godo tag [flags]

Aliases:

	tag, t

Flags:

	-h, --help   help for tag
*/
package cmd

import (
	"fmt"
	"strconv"

	c "github.com/jwalton/gchalk"
	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

// tagCmd represents the tag command.
var tagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"t"},
	Short:   "Change a tag for a todo",
	Long:    `Tag is used to add or replace a tag for a todo.`,
	RunE:    tagRun,
}

func tagRun(cmd *cobra.Command, args []string) error {
	var command string = "tag"

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

		// Add the tag to the todo Tag field. Then update the changes.
		t.Tag = args[1] // Assign the tag arguments for the todo.

		fmt.Printf("%s Adding tag...\n%s %q: %s\n",
			c.StyleMust(Theme.Primary)("::"),
			c.StyleMust(Theme.Secondary)("-->"),
			t.Title,
			c.StyleMust(Theme.Tag)(t.Tag))

		err = updateTodo(p, command, todos)
		if err != nil {
			return fmt.Errorf("%v: %w", command, err)
		}
	} else {
		return fmt.Errorf("(%v) todo does not exist at position \"%d\"\n", command, p)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
