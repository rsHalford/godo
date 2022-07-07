/*
Find command helps you search for todos containing the provided string.

Usage:

	godo find [flags]

Aliases:

	find, fd, f

Flags:

	-c, --case string   choose case sensitivity pattern for search (default "smart")
	-h, --help          help for find
	-T, --tag           show the todos tag
	-t, --title         only show todo titles
*/
package cmd

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"text/tabwriter"
	"unicode"

	"github.com/rsHalford/godo/config"
	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

// findCmd represents the find command.
var findCmd = &cobra.Command{
	Use:     "find",
	Aliases: []string{"fd", "f"},
	Short:   "Search for a given string",
	Long:    `Find command helps you search for todos containing the provided string.`,
	RunE:    findRun,
}

var caseOpt string

func findRun(cmd *cobra.Command, args []string) error {
	var command string = "find"

	todos, err := todo.Todos() // Get todos from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	sort.Sort(todo.Order(todos)) // Sort the todos for terminal printing.

	w := tabwriter.NewWriter(os.Stdout, minwidth, tabwidth, padding, padchar, flags)
	flagSet := cmd.Flags().Lookup("case").Changed // True if --case flag provided.
	caseOpt = searchPattern(flagSet)

	// For each argument, check the title and body of every todo according to case
	// -sensitivity settings. Printing the results depending on the --title flag.
	for _, a := range args {
		switch {
		// Return all case-sensitive matches, according to --title and --tag flags.
		case caseOpt == "sensitive":
			for _, t := range todos {
				printFindMatches(w, t, t.Body, t.Title, a)
			}

		// Change the argument and results to lower-case before matching.
		case caseOpt == "insensitive":
			for _, t := range todos {
				a = strings.ToLower(a)
				body := strings.ToLower(t.Body)
				title := strings.ToLower(t.Title)
				printFindMatches(w, t, body, title, a)
			}

		// Sensitivity is only implemented if the argument contains upper-case.
		// Changing to lower-case if the argument is only lower-case.
		default:
			var hasUpper bool

			// Check each rune of the argument for upper-case characters.
			for _, r := range a {
				if unicode.IsUpper(r) {
					hasUpper = true

					break
				}
			}

			for _, t := range todos {
				body := t.Body
				title := t.Title

				// Perform a case-insensitive search if hasUpper is true.
				if !hasUpper {
					body = strings.ToLower(body)
					title = strings.ToLower(title)
				}

				printFindMatches(w, t, body, title, a)
			}
		}
	}

	// TODO: assign error message to tabwriter.Writer.Flush().
	w.Flush()

	return nil
}

// searchPattern returns the user provided argument for --case/-c, if one is
// set. Otherwise, the config.toml option for caseSensitivity will determine
// the string value for caseOpt. Finally, if the config.toml value is unset,
// default to the "smart" pattern.
func searchPattern(set bool) string {
	switch {
	case set:
		break
	case config.Value("caseSensitivity") != "":
		caseOpt = config.Value("caseSensitivity")
	default:
		caseOpt = "smart"
	}

	return caseOpt
}

// printFindMatches searches the given todos, body and title for a matching
// string against the command's argument, a. And sends all matches to stdout,
// using the Writer to format the the results.
func printFindMatches(w *tabwriter.Writer, t todo.Todo, body, title, a string) {
	if strings.Contains(body, a) || strings.Contains(title, a) {
		switch {
		// Only print the title and tag for the matching todo.
		case tagOpt && titleOpt:
			fmt.Fprintln(w, t.PositionFmt()+t.TagFmt(t.Tag)+t.TitleFmt(t.Title))

		// Only print the title of the matching todo.
		case titleOpt:
			fmt.Fprintln(w, t.PositionFmt()+t.TitleFmt(t.Title))

		// Print the title, tag and body content for the matching todo.
		case tagOpt:
			fmt.Fprintln(w, t.PositionFmt()+t.TagFmt(t.Tag)+t.TitleFmt(t.Title)+
				"\n"+t.Body+"\n")

		// Default to printing just the title and body of the matching todo.
		default:
			fmt.Fprintln(w, t.PositionFmt()+t.TitleFmt(t.Title)+
				"\n"+t.Body+"\n")
		}
	}
}

func init() {
	rootCmd.AddCommand(findCmd)

	// The --title flag argument determines if only the todo title will be printed.
	findCmd.Flags().BoolVarP(&titleOpt, "title", "t", false, "only show todo titles")
	// The --tag flag determines whether the tag for each todo should be shown.
	findCmd.Flags().BoolVarP(&tagOpt, "tag", "T", false, "show the todos tag")
	// The --case flag determines what type of argument case matching occurs.
	findCmd.Flags().StringVarP(&caseOpt, "case", "c", "smart", "choose case sensitivity pattern for search")
}
