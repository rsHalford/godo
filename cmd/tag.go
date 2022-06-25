/*
Tag is used to add or replace a tag to a todo.

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
	Short:   "Add a tag to your todo",
	Long:    `Tag is used to add or replace a tag for a todo.`,
	RunE:    tagRun,
}

func tagRun(cmd *cobra.Command, args []string) error {
	var command string = "tag"

	items, err := todo.Todos() // Get todo items from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	i, err := strconv.Atoi(args[0]) // Convert todo id argument to an integer.
	if err != nil {
		return fmt.Errorf("%v: %q %w", command, args[0], err)
	}

	if i > 0 && i <= len(items) { // Validate id argument.
		// Add the tag to the todo item Tag field. Then update the changes.
		items[i-1].Tag = args[1] // Assign the tag arguments for the todo.

		fmt.Printf("%s Adding tag...\n%s %q: %s\n",
			c.StyleMust(Theme.Primary)("::"),
			c.StyleMust(Theme.Secondary)("-->"),
			items[i-1].Title,
			c.StyleMust(Theme.Tag)(items[i-1].Tag))

		err = updateTodo(i, command, items)
		if err != nil {
			return fmt.Errorf("%v: %w", command, err)
		}
	} else {
		return fmt.Errorf("%v: %q %w", i, command, err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
