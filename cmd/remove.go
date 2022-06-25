/*
Remove will delete a todo from your list, forever.

Usage:

	godo remove [flags]

Aliases:

	remove, rm, r

Flags:

	-h, --help   help for remove
*/
package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	c "github.com/jwalton/gchalk"
	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command.
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm", "r"},
	Short:   "Remove a todo",
	Long:    `Remove will delete a todo from your list, forever.`,
	RunE:    removeRun,
}

func removeRun(cmd *cobra.Command, args []string) error {
	var command string = "remove"

	items, err := todo.Todos() // Get todo items from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	i, err := strconv.Atoi(args[0]) // Convert todo id argument to an integer.
	if err != nil {
		return fmt.Errorf("%v: %q %w", command, args[0], err)
	}

	if i > 0 && i <= len(items) { // Validate id argument.
		var filename string
		var isConfirmed bool
		var title string = items[i-1].Title

		// Check for confirmation of the todo's removal.
		if isConfirmed, err = confirmRemove(title); isConfirmed {
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			// Remove the todo item from the items slice, by copying the
			// higher-numbered elements down by one.
			items = items[:i-1+copy(items[i-1:], items[i:])]

			sort.Sort(todo.Order(items)) // Sort the items before saving.

			// Pass the filename of the local todo store to the filename variable.
			filename, err = todo.LocalTodos()
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			// Using SaveLocal to add the new todo(s) to the local JSON store.
			err = todo.SaveLocal(filename, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			fmt.Printf("\n%q %v\n", title, "deleted")
		}
	} else {
		return fmt.Errorf("%v: %q %w", command, i, err)
	}

	return nil
}

// confirmRemove prompts the user to confirm that they want to remove the todo.
func confirmRemove(title string) (bool, error) {
	var response string

	// Print which todo is being removed and prompt for confirmation from the user.
	fmt.Printf("%s Removing todo...\n%s %q\n%s Proceed with removal? (y/n): ",
		c.StyleMust(Theme.Primary)("::"),
		c.StyleMust(Theme.Secondary)("-->"),
		title,
		c.StyleMust(Theme.Primary)("::"))

	if _, err := fmt.Scanln(&response); err != nil {
		return false, fmt.Errorf("reading response: %w", err)
	}

	// Convert the user input to lower case and return the appropriate boolean response.
	switch strings.ToLower(response) {
	case "y", "yes":
		return true, nil
	case "n", "no":
		return false, nil
	default:
		fmt.Println("Please type (y)es or (n)o and press enter:")

		// Rerun the confirmRemove function if the user input is invalid.
		return confirmRemove(title)
	}
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
