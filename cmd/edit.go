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
	"io/ioutil"
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
)

func editRun(cmd *cobra.Command, args []string) error {
	var command string = "edit"

	items, err := todo.Todos()
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	i, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%v: \"%v\" %w", command, args[0], err)
	}

	if i > 0 && i <= len(items) {
		switch {
		case titleOpt:
			items[i-1].Title, err = createTemp([]byte(items[i-1].Title))
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			err = Update(i, command, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		case config.Value("editing_default") == "body" || bodyOpt:
			items[i-1].Body, err = createTemp([]byte(items[i-1].Body))
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			err = Update(i, command, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		default:
			items[i-1].Title, err = createTemp([]byte(items[i-1].Title))
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}

			err = Update(i, command, items)
			if err != nil {
				return fmt.Errorf("%v: %w", command, err)
			}
		}
	} else {
		return fmt.Errorf("%v: \"%v\" %w", command, i, err)
	}

	return nil
}

func createTemp(text []byte) (file string, err error) {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "godo-")
	if err != nil {
		return "", fmt.Errorf("create temporary file: %w", err)
	}
	defer os.Remove(tmpFile.Name())
	if _, err = tmpFile.Write(text); err != nil {
		return "", fmt.Errorf("write temporary file: %w", err)
	}
	if err = editTemp(tmpFile.Name()); err != nil {
		return "", fmt.Errorf("edit temporary file: %w", err)
	}
	data, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		return "", fmt.Errorf("read temporary file: %w", err)
	}
	todo := string(data)
	file = strings.TrimSuffix(todo, "\n")
	if err := tmpFile.Close(); err != nil {
		return "", fmt.Errorf("close temporary file: %w", err)
	}

	return
}

func editTemp(filename string) error {
	editor := defaultEditor()
	executable, err := exec.LookPath(editor)
	if err != nil {
		return fmt.Errorf("execute default editor: %w", err)
	}
	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func defaultEditor() (editor string) {
	if config.Value("editing_editor") == "" {
		editor = os.Getenv("EDITOR")

		return
	}
	editor = config.Value("editing_editor")

	return
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().BoolVarP(&titleOpt, "title", "t", false, "edit item title")
	editCmd.Flags().BoolVarP(&bodyOpt, "body", "b", false, "edit item body")
}
