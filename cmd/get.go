/*
Get lets you select a specific todo by passing it's list position as an
argument.

Usage:

	godo get [flags]

Aliases:

	get, g

Flags:

	-b, --body   get only the todo body
	-h, --help   help for get
	-T, --tag    show the todos tag
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

// getCmd represents the get command.
var getCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"g"},
	Short:   "Get a specific todo",
	Long: `Get lets you select a specific todo by passing it's list position as an
argument.`,
	RunE: getRun,
}

func getRun(cmd *cobra.Command, args []string) error {
	var command string = "get"

	todos, err := todo.Todos() // Get todos from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	w := tabwriter.NewWriter(os.Stdout, minwidth, tabwidth, padding, padchar, flags)

	// Convert todo position argument to an integer.
	p, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%v: %q %w", command, args[0], err)
	}

	if p > 0 && p <= len(todos) { // Validate position argument.
		t := todos[p-1]

		// Print only the body of the todo if the --body flag is given.
		// Otherwise print the position and title too. And the todo tag, if the
		// --tag is provided.
		switch {
		case bodyOpt:
			fmt.Fprintln(w, t.Body)
		case tagOpt:
			fmt.Fprintln(w, t.PositionFmt()+t.TagFmt(t.Tag)+t.TitleFmt(t.Title)+
				"\n"+t.Body)
		default:
			fmt.Fprintln(w, t.PositionFmt()+t.TitleFmt(t.Title)+"\n"+t.Body)
		}
	} else {
		return fmt.Errorf("(%v) todo does not exist at position \"%d\"\n", command, p)
	}

	w.Flush()

	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)

	// The --body flag argument determines if only the todo body will be printed.
	getCmd.Flags().BoolVarP(&bodyOpt, "body", "b", false, "get only the todo body")
	// The --tag flag determines whether the tag for the todo should be shown.
	getCmd.Flags().BoolVarP(&tagOpt, "tag", "T", false, "show the todos tag")
}
