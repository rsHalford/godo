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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "get a sepcific todo",
	Long:    `Get lets you select a specific todo by passing it's ID as an argument.`,
	Run:     getRun,
}

func getRun(cmd *cobra.Command, args []string) {
	items, err := todo.GetTodos()
	if err != nil {
		fmt.Println("No entries found")
		return
	}
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("\"%v\" is not a valid argument\n", args[0])
		return
	}
	if i > 0 && i <= len(items) {
		item := items[i-1]
		fmt.Fprintln(w, "\033[90m"+item.Label()+"\t\t"+"\033[0m"+item.PriorityFlag()+item.StatusFlag()+item.Title+"\033[0m\n"+item.Body)
	} else {
		fmt.Printf("\"%v\" is not a valid argument\n", args[0])
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
