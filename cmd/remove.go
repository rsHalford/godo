/*
Copyright © 2021 Richard Halford <richard@xhalford.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/rsHalford/godo/config"
	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

// removeCmd represents the remove command.
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm", "r"},
	Short:   "remove a todo",
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

		// Check for confirmation of the todo's removal.
		if isConfirmed, err = confirmRemove(items[i-1].Title); isConfirmed {
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			// Delete from an API if one is set in the configuration file.
			if config.Value("goapi_api") != "" {
				err = todo.DeleteRemote(
					config.Value("goapi_api"),
					config.Value("goapi_username"),
					config.Value("goapi_password"),
					fmt.Sprint(items[i-1].ID),
				)
				if err != nil {
					return fmt.Errorf("%v: %w", command, err)
				}
			} else {
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
			}

			fmt.Printf("\n%q %v\n", items[i-1].Title, "deleted")
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
	fmt.Printf("\033[34m::\033[0m Removing todo...\n\n\033[33m-->\033[0m %q\n\n\033[32m::\033[0m Proceed with removal? (y/n): ", title)

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
