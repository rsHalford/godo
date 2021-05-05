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
	"os"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:     "find",
	Aliases: []string{"fd", "f"},
	Short:   "search for a given string",
	Long: `The find command helps you search for todos containing a certain 
provide string`,
	Run: findRun,
}

func findRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}
	sort.Sort(todo.Order(items))
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	fmt.Fprintln(w, "#\tD\tP\tT\t")
	for _, a := range args {
		for _, i := range items {
			if strings.Contains(i.Text, a) {
				fmt.Fprintln(w, i.Label()+"\t"+i.PretDone()+"\t"+i.PretPriority()+"\t"+i.Text+"\t")
			}
		}
	}
	w.Flush()
}

//func iterItems([]string)

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
