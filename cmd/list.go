/*
List todos.

Usage:

	godo list [flags]

Aliases:

	list, ls, l

Flags:

	-a, --all        show all todos
	-d, --done       show completed todos
	-h, --help       help for list
	-p, --priority   show proritised todos
	-T, --tag        show the todos tag
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
	Short:   "List todos",
	Long:    `List todos.`,
	RunE:    listRun,
}

func listRun(cmd *cobra.Command, args []string) error {
	var command string = "list"

	todos, err := todo.Todos() // Get todos from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	sort.Sort(todo.Order(todos)) // Sort the todos for terminal printing.

	// Create a new writer with defined formatting.
	w := tabwriter.NewWriter(os.Stdout, minwidth, tabwidth, padding, padchar, flags)

	// Print as a list each todo that qualifies via the flag arguments given.
	if priorityOpt {
		for _, t := range todos {
			if t.Priority && (allOpt || t.Done == doneOpt) {
				if tagOpt {
					fmt.Fprintln(w, t.PositionFmt()+t.TagFmt(t.Tag)+
						t.TitleFmt(t.Title))
				} else {
					fmt.Fprintln(w, t.PositionFmt()+t.TitleFmt(t.Title))
				}
			}
		}
	} else {
		for _, t := range todos {
			if allOpt || t.Done == doneOpt {
				if tagOpt {
					fmt.Fprintln(w, t.PositionFmt()+t.TagFmt(t.Tag)+
						t.TitleFmt(t.Title))
				} else {
					fmt.Fprintln(w, t.PositionFmt()+t.TitleFmt(t.Title))
				}
			}
		}
	}

	w.Flush()

	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)

	// The --priority/--done/--all flag arguments determine which todos to list.
	listCmd.Flags().BoolVarP(&priorityOpt, "priority", "p", false, "show prioritised todos")
	listCmd.Flags().BoolVarP(&doneOpt, "done", "d", false, "show completed todos")
	listCmd.Flags().BoolVarP(&allOpt, "all", "a", false, "show all todos")
	// The --tag flag determines whether the tag for each todo should be shown.
	listCmd.Flags().BoolVarP(&tagOpt, "tag", "T", false, "show the todos tag")
}
