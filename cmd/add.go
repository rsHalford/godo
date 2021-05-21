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

	"github.com/rsHalford/godo/config"
	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority bool

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "add a new todo",
	Long:    `Add will create a new todo to the list.`,
	Run:     addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := todo.GetTodos()
	if err != nil {
		fmt.Println("Setting up database...")
	}
	for _, x := range args {
		item := todo.Todo{Body: x}
		item.Prioritise(priority)
		if config.GetString("api") != "" {
			todo.CreateRemoteTodo(config.GetString("api"), item)
		}
		items = append(items, item)
	}
	if config.GetString("api") == "" {
		todo.SaveTodos(viper.GetString("datafile"), items)
	}
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().BoolVarP(&priority, "priority", "p", false, "assign priority to your todo")
}
