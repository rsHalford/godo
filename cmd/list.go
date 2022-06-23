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
	"sort"
	"text/tabwriter"

	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

var (
	priorityOpt bool
	doneOpt     bool
	allOpt      bool
	tagOpt      bool
)

const (
	minwidth int  = 3
	tabwidth int  = 0
	padding  int  = 1
	padchar  byte = ' '
	flags    uint = 0
)

// listCmd represents the list command.
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "list your todos",
	Long:    `Listing all of your todos.`,
	RunE:    listRun,
}

func listRun(cmd *cobra.Command, args []string) error {
	var command string = "list"

	items, err := todo.Todos() // Get todo items from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	sort.Sort(todo.Order(items)) // Sort the items for terminal printing.

	// Create a new writer with defined formatting.
	w := tabwriter.NewWriter(os.Stdout, minwidth, tabwidth, padding, padchar, flags)

	// Print as a list each todo item that qualifies via the flag arguments given.
	if priorityOpt {
		for _, i := range items {
			if i.Priority && (allOpt || i.Status == doneOpt) {
				if tagOpt {
					fmt.Fprintln(w, i.Label()+i.TagFmt(i.Tag)+i.TitleFmt(i.Title))
				} else {
					fmt.Fprintln(w, i.Label()+i.TitleFmt(i.Title))
				}
			}
		}
	} else {
		for _, i := range items {
			if allOpt || i.Status == doneOpt {
				if tagOpt {
					fmt.Fprintln(w, i.Label()+i.TagFmt(i.Tag)+i.TitleFmt(i.Title))
				} else {
					fmt.Fprintln(w, i.Label()+i.TitleFmt(i.Title))
				}
			}
		}
	}

	w.Flush()

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)

	// The --priority/--done/--all flag arguments determine which items to list.
	listCmd.Flags().BoolVarP(&priorityOpt, "priority", "p", false, "show prioritised todos")
	listCmd.Flags().BoolVarP(&doneOpt, "done", "d", false, "show completed todos")
	listCmd.Flags().BoolVarP(&allOpt, "all", "a", false, "show all todos")
	// The --tag flag determines whether the tag for each todo should be shown.
	listCmd.Flags().BoolVarP(&tagOpt, "tag", "T", false, "show the todo's tag")
}
