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
	doneOpt bool
	allOpt  bool
	// individualOpt int
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "l"},
	Short:   "list your todos",
	Long:    `Listing all of your todos.`,
	Run:     listRun,
}

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.GetTodos()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	sort.Sort(todo.Order(items))
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		if allOpt || i.Status == doneOpt {
			fmt.Fprintln(w, "\033[90m"+i.Label()+"\t\t"+"\033[0m"+i.PriorityFlag()+i.StatusFlag()+i.Body+"\033[0m")
		}
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().BoolVarP(&doneOpt, "done", "d", false, "show completed todos")
	listCmd.Flags().BoolVarP(&allOpt, "all", "a", false, "show all todos")
	// listCmd.Flags().IntVar(&individualOpt, )
}
