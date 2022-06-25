/*
Get lets you select a specific todo by passing it's ID as an argument.

Usage:

	godo get [flags]

Aliases:

	get, g

Flags:

	-b, --body   get only the item body
	-h, --help   help for get
	-T, --tag    show the todo's tag
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
	Long:    `Get lets you select a specific todo by passing it's ID as an argument.`,
	RunE:    getRun,
}

func getRun(cmd *cobra.Command, args []string) error {
	var command string = "get"

	items, err := todo.Todos() // Get todo items from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	// Create a new writer with defined formatting.
	w := tabwriter.NewWriter(os.Stdout, minwidth, tabwidth, padding, padchar, flags)

	i, err := strconv.Atoi(args[0]) // Convert todo id argument to an integer.
	if err != nil {
		return fmt.Errorf("%v: %q %w", command, args[0], err)
	}

	if i > 0 && i <= len(items) { // Validate id argument.
		item := items[i-1]

		// Print only the body of the item if the --body flag is given.
		// Otherwise print the position id and title too. And the todo tag,
		// if the --tag is provided.
		switch {
		case bodyOpt:
			fmt.Fprintln(w, item.Body)
		case tagOpt:
			fmt.Fprintln(w, item.Label()+item.TagFmt(item.Tag)+
				item.TitleFmt(item.Title)+"\n"+item.Body)
		default:
			fmt.Fprintln(w, item.Label()+item.TitleFmt(item.Title)+"\n"+item.Body)
		}
	} else {
		return fmt.Errorf("%v: %q %w", command, i, err)
	}

	w.Flush()

	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)

	// The --body flag argument determines if only the item body will be printed.
	getCmd.Flags().BoolVarP(&bodyOpt, "body", "b", false, "get only the item body")
	// The --tag flag determines whether the tag for the todo should be shown.
	getCmd.Flags().BoolVarP(&tagOpt, "tag", "T", false, "show the todo's tag")
}
