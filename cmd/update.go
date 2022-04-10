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

	"github.com/rsHalford/godo/todo"
)

// updateTodo will send the updated item and it's properties to save locally.
func updateTodo(command string, items []todo.Todo) error {
	sort.Sort(todo.Order(items)) // Sort the items before saving.

	// Pass the filename of the local todo store to the filename variable.
	filename, err := todo.LocalTodos()
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
