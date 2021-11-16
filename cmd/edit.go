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
	"os/exec"
	"strconv"
	"strings"

	"github.com/rsHalford/godo/config"
	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command.
var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Short:   "edit a todo (default: edit title)",
	Long:    `Edit a todo by passing the list number of the todo. Defaults to editing the todo title, if not set in godo.yaml`,
	RunE:    editRun,
}

var (
	titleOpt bool
	bodyOpt  bool
	extOpt   string
)

func editRun(cmd *cobra.Command, args []string) error {
	var command string = "edit"

	items, err := todo.Todos() // Get todo items from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	i, err := strconv.Atoi(args[0]) // Convert todo id argument to an integer.
	if err != nil {
		return fmt.Errorf("%v: \"%v\" %w", command, args[0], err)
	}

	if i > 0 && i <= len(items) { // Validate id argument.
		// Perform the edit on either the title or body. Depending on
		// which arguments or configuration settings have been set.
		switch {
		case titleOpt:
			items[i-1].Title, err = createTemp([]byte(items[i-1].Title))
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			err = updateTodo(i, command, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		case config.Value("editing_default") == "body" || bodyOpt:
			items[i-1].Body, err = createTemp([]byte(items[i-1].Body))
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			err = updateTodo(i, command, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		default:
			items[i-1].Title, err = createTemp([]byte(items[i-1].Title))
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			err = updateTodo(i, command, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		}
	} else {
		return fmt.Errorf("%v: \"%v\" %w", command, i, err)
	}

	return nil
}

// createTemp creates a temporary file within the system's temporary directory.
// Writes the current title or body to the file for editing. Then reads the
// file, converting the edited data to be returned, before deleting it.
func createTemp(text []byte) (string, error) {
	ext := tempFiletype() // Check and return filetype extension if set.

	f, err := os.CreateTemp(os.TempDir(), "godo-"+ext)
	if err != nil {
		return "", fmt.Errorf("create temporary file: %w", err)
	}

	defer f.Close()           // Close the file at the end of the function.
	defer os.Remove(f.Name()) // Remove the file at the end of the function.

	_, err = f.Write(text) // Add the current title/body to the temporary file.
	if err != nil {
		return "", fmt.Errorf("write temporary file: %w", err)
	}

	err = editTemp(f.Name()) // Open the file with an editor.
	if err != nil {
		return "", fmt.Errorf("edit temporary file: %w", err)
	}

	data, err := os.ReadFile(f.Name()) // Read the edited file contents to data.
	if err != nil {
		return "", fmt.Errorf("read temporary file: %w", err)
	}

	todo := string(data) // Convert data []byte to string.

	file := strings.TrimSuffix(todo, "\n") // Remove any end of file whitespace.

	return file, nil
}

// editTemp opens a temporary file with an editor.
func editTemp(filename string) error {
	editor, err := defaultEditor() // Identify an executable editor's path.
	if err != nil {
		return fmt.Errorf("editTemp: %w", err)
	}

	executable, err := exec.LookPath(editor)
	if err != nil {
		return fmt.Errorf("find editor executable path: %w", err)
	}

	// Run the identified editor's executable, with the provided file.
	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// defaultEditor determines which editor has been chosen within the configuration file.
// If not specified, the system $EDITOR environment variable will be used.
func defaultEditor() (string, error) {
	if config.Value("editing_editor") == "" {
		editor := os.Getenv("EDITOR")

		// If the system environment variable is also not set, inform the user.
		if editor == "" {
			return "", fmt.Errorf("environment variable $EDITOR not set")
		}

		return editor, nil
	}
	editor := config.Value("editing_editor")

	return editor, nil
}

// tempFiletype will check if there is a string variable flag set for the tempFiletype
// before checking the config.yaml. Returning a formatted extension string. If neither
// is set an empty string is returned.
func tempFiletype() (extension string) {
	switch {
	case extOpt != "":
		extension = "*." + extOpt

		return

	case config.Value("editing_filetype") != "":
		extension = "*." + config.Value("editing_filetype")

		return

	default:
		return ""
	}
}

func init() {
	rootCmd.AddCommand(editCmd)

	// The --title/--body flag arguments determine which part of the todo to edit.
	editCmd.Flags().BoolVarP(&titleOpt, "title", "t", false, "edit item title")
	editCmd.Flags().BoolVarP(&bodyOpt, "body", "b", false, "edit item body")
	editCmd.Flags().StringVarP(&extOpt, "extension", "e", "", "set temporary filetype extension")
}
