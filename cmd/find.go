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
	Short:   "search for a given string",
	Long: `The find command helps you search for todos containing the
provided string`,
	RunE: findRun,
}

var caseOpt string

func findRun(cmd *cobra.Command, args []string) error {
	var command string = "find"

	items, err := todo.Todos() // Get todo items from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	sort.Sort(todo.Order(items)) // Sort the items for terminal printing.

	// Create a new writer with defined formatting.
	w := tabwriter.NewWriter(os.Stdout, minwidth, tabwidth, padding, padchar, flags)

	// Assign the bool value if an argument for the --case/-c flag is provided.
	flagSet := cmd.Flags().Lookup("case").Changed

	// If the user provides an argument for --case/-c, that will be used. Otherwise,
	// the config.yaml option for caseSensitivity will determine the string value
	// for caseOpt. Finally, if the config.yaml value is unset, default to the
	// "smart" pattern.
	switch {
	case flagSet:
		break
	case config.Value("caseSensitivity") != "":
		caseOpt = config.Value("caseSensitivity")
	default:
		caseOpt = "smart"
	}

	// For every argument string, go through every todo item and check both
	// the title and body for the string, depending on case-sensitivity settings.
	// Then print the todo item title - exclusively if the --title/-t flag is used
	// - and also print the body.
	for _, a := range args {
		switch {
		// For sensitive search results, for each item return all case-sensitive
		// matches, according to the --title and --tag flag arguments.
		case caseOpt == "sensitive":
			for _, i := range items {
				printFindMatches(w, i, i.Body, i.Title, a)
			}

		// For insensitive search results, change both argument and results
		// to lower-case. Then for each item return all matches in their original
		// format, according to the --title and --tag flag arguments.
		case caseOpt == "insensitive":
			for _, i := range items {
				a = strings.ToLower(a)
				body := strings.ToLower(i.Body)
				title := strings.ToLower(i.Title)
				printFindMatches(w, i, body, title, a)
			}

		// Implement a smart search, where case sensitivity is only implemented
		// if the command argument contains an upper-case character. And only
		// change results to lower-case if the argument only contains lower-case
		// characters. Then for each item return all matches in their original
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

			for _, i := range items {
				body := i.Body
				title := i.Title

				// Perform a case-insensitive search if hasUpper is true.
				if !hasUpper {
					body = strings.ToLower(body)
					title = strings.ToLower(title)
				}

				printFindMatches(w, i, body, title, a)
			}
		}
	}

	// TODO: assign error message to tabwriter.Writer.Flush()
	w.Flush()

	return nil
}

// printFindMatches searches the given todo item's, body and title for a matching string
// against the command's argument, a. And sends all matches to stdout, using the Writer
// to format the the results.
func printFindMatches(w *tabwriter.Writer, i todo.Todo, body, title, a string) {
	if strings.Contains(body, a) || strings.Contains(title, a) {
		switch {
		// Only print the title and tag for the matching item.
		case tagOpt && titleOpt:
			fmt.Fprintln(w, i.Label()+i.TagFmt(i.Tag)+i.TitleFmt(i.Title))

		// Only print the title of the matching item.
		case titleOpt:
			fmt.Fprintln(w, i.Label()+i.TitleFmt(i.Title))

		// Print the title, tag and body content for the matching item.
		case tagOpt:
			fmt.Fprintln(w, i.Label()+i.TagFmt(i.Tag)+i.TitleFmt(i.Title)+"\n"+i.Body+"\n")

		// Default to printing just the title and body of the mathcing todo item.
		default:
			fmt.Fprintln(w, i.Label()+i.TitleFmt(i.Title)+"\n"+i.Body+"\n")
		}
	}
}

func init() {
	rootCmd.AddCommand(findCmd)

	// The --title flag argument determines if only the item title will be printed.
	findCmd.Flags().BoolVarP(&titleOpt, "title", "t", false, "only show item titles")
	// The --tag flag determines whether the tag for each todo should be shown.
	findCmd.Flags().BoolVarP(&tagOpt, "tag", "T", false, "show the todo's tag")
	// The --case flage determines what type of argument case matching occurs.
	findCmd.Flags().StringVarP(&caseOpt, "case", "c", "smart", "choose case sensitivity pattern for search")
}
