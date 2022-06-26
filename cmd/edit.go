/*
Edit a todo by passing the list position number of the todo. Defaults to
editing the todo title, if not set in config.toml.

Usage:

	godo edit [flags]

Aliases:

	edit, ed, e

Flags:

	-b, --body               edit todo body
	-e, --extension string   set filetype extension for editing
	-h, --help               help for edit
	-t, --title              edit todo title
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
	Aliases: []string{"ed", "e"},
	Short:   "Edit a todo (default: edit title)",
	Long: `Edit a todo by passing the list position number of the todo. Defaults to
editing the todo title, if not set in config.toml.`,
	RunE: editRun,
}

var (
	titleOpt bool
	bodyOpt  bool
	extOpt   string
)

func editRun(cmd *cobra.Command, args []string) error {
	var command string = "edit"

	todos, err := todo.Todos() // Get todos from the configured source.
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	// Convert todo position argument to an integer.
	p, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%v: %q %w", command, args[0], err)
	}

	if p > 0 && p <= len(todos) { // Validate position argument.
		// Perform the edit on either the title or body. Depending on which
		// arguments or configuration settings have been set.
		switch {
		case titleOpt:
			todos[p-1].Title, err = createTemp([]byte(todos[p-1].Title))
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			err = updateTodo(p, command, todos)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		case config.Value("editing_default") == "body" || bodyOpt:
			todos[p-1].Body, err = createTemp([]byte(todos[p-1].Body))
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			err = updateTodo(p, command, todos)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		default:
			todos[p-1].Title, err = createTemp([]byte(todos[p-1].Title))
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			err = updateTodo(p, command, todos)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		}
	} else {
		return fmt.Errorf("%v: %q %w", command, p, err)
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

// defaultEditor determines which editor has been chosen within the
// configuration file. If not specified, the system $EDITOR environment
// variable will be used.
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

// tempFiletype will check if there is a string variable flag set for the
// tempFiletype before checking the config.toml. Returning a formatted
// extension string. If neither is set an empty string is returned.
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
	editCmd.Flags().BoolVarP(&titleOpt, "title", "t", false, "edit todo title")
	editCmd.Flags().BoolVarP(&bodyOpt, "body", "b", false, "edit todo body")
	editCmd.Flags().StringVarP(&extOpt, "extension", "e", "", "set filetype extension for editing")
}
