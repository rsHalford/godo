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

	// Create a new writer with defined formatting.
	w := tabwriter.NewWriter(os.Stdout, minwidth, tabwidth, padding, padchar, flags)

	// Assign the bool value if an argument for the --case/-c flag is provided.
	flagSet := cmd.Flags().Lookup("case").Changed

	// If the user provides an argument for --case/-c, that will be used.
	// Otherwise, the config.toml option for caseSensitivity will determine the
	// string value for caseOpt. Finally, if the config.toml value is unset,
	// default to the "smart" pattern.
	switch {
	case flagSet:
		break
	case config.Value("caseSensitivity") != "":
		caseOpt = config.Value("caseSensitivity")
	default:
		caseOpt = "smart"
	}

	// For every argument string, go through every todo and check both the title
	// and body for the string, depending on case-sensitivity settings. Then
	// print the todo title - exclusively if the --title/-t flag is used - and
	// also print the body.
	for _, a := range args {
		switch {
		// For sensitive search results, for each argument return all
		// case-sensitive matches, according to the --title and --tag flag
		// arguments.
		case caseOpt == "sensitive":
			for _, t := range todos {
				printFindMatches(w, t, t.Body, t.Title, a)
			}

		// For insensitive search results, change both argument and results
		// to lower-case. Then for each todo, return all matches in their
		// original format, according to the --title and --tag flag arguments.
		case caseOpt == "insensitive":
			for _, t := range todos {
				a = strings.ToLower(a)
				body := strings.ToLower(t.Body)
				title := strings.ToLower(t.Title)
				printFindMatches(w, t, body, title, a)
			}

		// Implement a smart search, where case sensitivity is only implemented
		// if the command argument contains an upper-case character. And only
		// change results to lower-case if the argument only contains lower-case
		// characters. Then for each todo return all matches in their original
		// format, according to the --title and --tag flag arguments.
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

	// TODO: assign error message to tabwriter.Writer.Flush()
	w.Flush()

	return nil
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
