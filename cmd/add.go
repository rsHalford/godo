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
	"time"

	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

// addCmd represents the add command.
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "add a new todo",
	Long:    `Add will create a new todo to the list.`,
	RunE:    addRun,
}

var (
	priority bool
	tagStr   string
)

func addRun(cmd *cobra.Command, args []string) error {
	var command string = "add"

	items, err := todo.Todos() // Get todo items from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	// Each argument given to the add command will be assigned to an individual item.
	for _, x := range args {
		var t time.Time = time.Now().UTC()

		item := todo.Todo{Title: x, CreatedAt: t, UpdatedAt: t}

		// Mark the item's priority as true, if the --priority flag is provided.
		// Currently this flag affects all new todo items being declared.
		item.Prioritise(priority)

		// Add the given tag to the item, if the --tag flag is provided.
		item.Tagging(tagStr)

		items = append(items, item) // Add the new complete item to items.
	}

	// Save the new todo(s) locally.
	var filename string

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

	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)

	// The --priority boolean flag is used to set the priority object of the new todo(s) as true.
	addCmd.Flags().BoolVarP(&priority, "priority", "p", false, "assign priority to your todo")
	// The --tag string flag is used to add a tag to the new todo(s).
	addCmd.Flags().StringVarP(&tagStr, "tag", "T", "", "add tag to new todo")
}
