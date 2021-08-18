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

// priorityCmd represents the priority command
var priorityCmd = &cobra.Command{
	Use:     "priority",
	Aliases: []string{"pri", "p"},
	Short:   "label a todo as a priority",
	Long:    `You can toggle a todo as being labelled a priority with the priority command.`,
	Run:     priorityRun,
}

func priorityRun(cmd *cobra.Command, args []string) {
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
		if items[i-1].Priority != true {
			items[i-1].Priority = true
			fmt.Printf("\033[34m::\033[0m Setting priority...\n\n\033[33m-->\033[0m %q\n", items[i-1].Title)
			if config.GetString("api") != "" {
				todo.UpdateRemoteTodo(config.GetString("api"), fmt.Sprint(items[i-1].ID), items[i-1])
				sort.Sort(todo.Order(items))
			} else {
				sort.Sort(todo.Order(items))
				todo.SaveTodos(viper.GetString("datafile"), items)
			}
		} else {
			items[i-1].Priority = false
			fmt.Printf("\033[34m::\033[0m Removing priority...\n\n\033[33m-->\033[0m %q\n", items[i-1].Title)
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
	rootCmd.AddCommand(priorityCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// priorityCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// priorityCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
