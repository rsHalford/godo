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
	"strconv"

	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

// tagCmd represents the tag command.
var tagCmd = &cobra.Command{
	Use:     "tag",
	Aliases: []string{"t"},
	Short:   "add a tag to your todo",
	Long:    `You can add a tag to a todo, using the tag command.`,
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

		fmt.Printf("\033[34m::\033[0m Adding tag...\n\033[33m-->\033[0m %q: \033[3m\033[35m%v\033[0m\n", items[i-1].Title, items[i-1].Tag)

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
