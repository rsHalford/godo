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

	items, err := todo.Todos()
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%v: \"%v\" %w", command, args[0], err)
	}

	if i > 0 && i <= len(items) {
		var filename string
		var isConfirmed bool

		if isConfirmed, err = confirmRemove(items[i-1].Title); isConfirmed {
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			fmt.Printf("%q %v\n", items[i-1].Title, "deleted")

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

				sort.Sort(todo.Order(items))
			} else {
				items = items[:i-1+copy(items[i-1:], items[i:])]

				sort.Sort(todo.Order(items))

				filename, err = todo.LocalTodos()
				if err != nil {
					return fmt.Errorf("%v: %w", command, err)
				}

				err = todo.SaveLocal(filename, items)
				if err != nil {
					return fmt.Errorf("%v: %w", command, err)
				}
			}
		}
	} else {
		return fmt.Errorf("%v: \"%v\" %w", command, i, err)
	}

	return nil
}

func confirmRemove(title string) (bool, error) {
	var response string

	fmt.Printf("\033[34m::\033[0m Removing todo...\n\n\033[33m-->\033[0m %q\n\n\033[32m::\033[0m Proceed with removal? (y/n): ", title)

	if _, err := fmt.Scanln(&response); err != nil {
		return false, fmt.Errorf("reading response: %w", err)
	}

	switch strings.ToLower(response) {
	case "y", "yes":
		return true, nil
	case "n", "no":
		return false, nil
	default:
		fmt.Println("Please type (y)es or (n)o and press enter:")

		return confirmRemove(title)
	}
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
