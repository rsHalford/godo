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

	"github.com/rsHalford/godo/config"
	"github.com/rsHalford/godo/todo"
)

func Update(i int, command string, items []todo.Todo) error {
	if config.Value("goapi_api") != "" {
		err := todo.UpdateRemote(
			config.Value("goapi_api"),
			config.Value("goapi_username"),
			config.Value("goapi_password"),
			fmt.Sprint(items[i-1].ID), items[i-1],
		)
		if err != nil {
			return fmt.Errorf("%v: %w", command, err)
		}

		sort.Sort(todo.Order(items))
	} else {
		sort.Sort(todo.Order(items))

		filename, err := todo.LocalTodos()
		if err != nil {
			return fmt.Errorf("%v: %w", command, err)
		}

		err = todo.SaveLocal(filename, items)
		if err != nil {
			return fmt.Errorf("%v: %w", command, err)
		}
	}

	return nil
}
