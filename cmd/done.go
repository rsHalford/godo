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
	"strconv"

	"github.com/rsHalford/godo/config"
	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"d", "do"},
	Short:   "toggle todo as done",
	Long:    `You can label a todo as done or active with the done command.`,
	Run:     doneRun,
}

func doneRun(cmd *cobra.Command, args []string) {
	items, err := todo.GetTodos()
	if err != nil {
		fmt.Println("No entries found")
		return
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("\"%v\" is not a valid argument\n", args[0])
		return
	}
	if i > 0 && i <= len(items) {
		if items[i-1].Status != true {
			items[i-1].Status = true
			fmt.Printf("%q %v\n", items[i-1].Body, "marked done")
			if config.GetString("api") != "" {
				todo.UpdateRemoteTodo(config.GetString("api"), fmt.Sprint(items[i-1].ID), items[i-1])
				sort.Sort(todo.Order(items))
			} else {
				sort.Sort(todo.Order(items))
				todo.SaveTodos(viper.GetString("datafile"), items)
			}
		} else {
			items[i-1].Status = false
			fmt.Printf("%q %v\n", items[i-1].Body, "marked active")
			if config.GetString("api") != "" {
				todo.UpdateRemoteTodo(config.GetString("api"), fmt.Sprint(items[i-1].ID), items[i-1])
				sort.Sort(todo.Order(items))
			} else {
				sort.Sort(todo.Order(items))
				todo.SaveTodos(viper.GetString("datafile"), items)
			}
		}
	} else {
		fmt.Printf("\"%v\" doesn't match any todos\n", i)
	}
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
