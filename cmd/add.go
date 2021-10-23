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

	"github.com/rsHalford/godo/config"
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

var priority bool

func addRun(cmd *cobra.Command, args []string) error {
	var command string = "add"

	items, err := todo.GetTodos()
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	for _, x := range args {
		item := todo.Todo{Title: x}

		item.Prioritise(priority)

		if config.GetString("goapi_api") != "" {
			err = todo.CreateRemoteTodo(config.GetString("goapi_api"), config.GetString("goapi_username"), config.GetString("goapi_password"), item)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		}

		items = append(items, item)
	}

	if config.GetString("goapi_api") == "" {
		var filename string

		filename, err = todo.LocalTodos()
		if err != nil {
			return fmt.Errorf("%v: %w", command, err)
		}

		err := todo.SaveTodos(filename, items)
		if err != nil {
			return fmt.Errorf("%v: %w", command, err)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVarP(&priority, "priority", "p", false, "assign priority to your todo")
}
