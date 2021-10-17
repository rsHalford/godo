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
	"strings"
	"text/tabwriter"

	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:     "find",
	Aliases: []string{"fd", "f"},
	Short:   "search for a given string",
	Long: `The find command helps you search for todos containing the
provided string`,
	Run: findRun,
}

func findRun(cmd *cobra.Command, args []string) {
	items, err := todo.GetTodos()
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	sort.Sort(todo.Order(items))
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, a := range args {
		for _, i := range items {
			if strings.Contains(i.Body, a) || strings.Contains(i.Title, a) {
				if titleOpt == true {
					fmt.Fprintln(w, "\033[90m"+i.Label()+"\t\t"+"\033[0m"+i.PriorityFlag()+i.StatusFlag()+i.Title+"\033[0m")
				} else {
					fmt.Fprintln(w, "\033[90m"+i.Label()+"\t\t"+"\033[0m"+i.PriorityFlag()+i.StatusFlag()+i.Title+"\033[0m\n"+i.Body+"\n")
				}
			}
		}
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	findCmd.Flags().BoolVarP(&titleOpt, "title", "t", false, "only show item titles")
}
