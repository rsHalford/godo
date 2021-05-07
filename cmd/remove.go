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
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm", "r"},
	Short:   "remove a todo",
	Long:    `Remove will delete a todo from your list, forever.`,
	Run:     removeRun,
}

func removeRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadTodos(viper.GetString("datafile"))
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
		if isConfirmed := confirmRemove(); isConfirmed {
			fmt.Printf("%q %v\n", items[i-1].Text, "deleted")
			items = items[:i-1+copy(items[i-1:], items[i:])]
			sort.Sort(todo.Order(items))
			todo.SaveTodos(viper.GetString("datafile"), items)
		}
	} else {
		fmt.Printf("\"%v\" doesn't match any todos\n", i)
	}
}

func confirmRemove() bool {
	var response string
	fmt.Printf("Confirm deletion? (y/n): ")
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		fmt.Println("Please type (y)es or (n)o and press enter:")
		return confirmRemove()
	}
}

func init() {
	rootCmd.AddCommand(removeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// remove all todos marked as done
}
