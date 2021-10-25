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
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

// getCmd represents the get command.
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a specific todo",
	Long:  `Get lets you select a specific todo by passing it's ID as an argument.`,
	RunE:  getRun,
}

func getRun(cmd *cobra.Command, args []string) error {
	var command string = "get"

	items, err := todo.Todos()
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	w := tabwriter.NewWriter(os.Stdout, minwidth, tabwidth, padding, padchar, flags)

	i, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%v: \"%v\" %w", command, args[0], err)
	}

	if i > 0 && i <= len(items) {
		item := items[i-1]
		if bodyOpt {
			fmt.Fprintln(w, item.Body)
		} else {
			fmt.Fprintln(w, "\033[90m"+item.Label()+"\t\t"+"\033[0m"+item.PriorityFlag()+item.StatusFlag()+item.Title+"\033[0m\n"+item.Body)
		}
	} else {
		return fmt.Errorf("%v: \"%v\" %w", command, i, err)
	}

	w.Flush()

	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&bodyOpt, "body", "b", false, "get only item body")
}
