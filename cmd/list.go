/*
List your todos.

Usage:

	godo list [flags]

Aliases:

	list, ls, l

Flags:

	-a, --all        show all todos
	-d, --done       show completed todos
	-h, --help       help for list
	-p, --priority   show proritised todos
	-T, --tag        show the todo's tag
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
	Short:   "List your todos",
	Long:    `List your todos.`,
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
