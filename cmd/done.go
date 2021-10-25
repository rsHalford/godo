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

// doneCmd represents the done command.
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"d", "do"},
	Short:   "toggle todo as done",
	Long:    `You can label a todo as done or active with the done command.`,
	RunE:    doneRun,
}

func doneRun(cmd *cobra.Command, args []string) error {
	var command string = "done"

	items, err := todo.Todos()
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%v: \"%v\" %w", command, args[0], err)
	}

	if i > 0 && i <= len(items) {
		if !items[i-1].Status {
			items[i-1].Status = true

			fmt.Printf("\033[34m::\033[0m Marked done...\n\n\033[33m-->\033[0m %q\n", items[i-1].Title)

			err = Update(i, command, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		} else {
			items[i-1].Status = false

			fmt.Printf("\033[34m::\033[0m Marked active...\n\n\033[33m-->\033[0m %q\n", items[i-1].Title)

			err = Update(i, command, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		}
	} else {
		return fmt.Errorf("%v: \"%v\" %w", i, command, err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
