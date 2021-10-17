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
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/rsHalford/godo/config"
	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	titleOpt bool
	bodyOpt  bool
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Short:   "edit a todo (default: edit title)",
	Long:    `Edit a todo by passing the list number of the todo. Defaults to editing the todo title, if not set in godo.yaml`,
	Run:     editRun,
}

func editRun(cmd *cobra.Command, args []string) {
	items, err := todo.GetTodos()
	if err != nil {
		fmt.Println("No entries found")
		return
	}
	i, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("\"%v\" is not a valid argument\n", args[0])
		return
	}
	if i > 0 && i <= len(items) {
		switch {
		case titleOpt:
			items[i-1].Title = createTemp([]byte(items[i-1].Title))
			if config.GetString("goapi_api") != "" {
				todo.UpdateRemoteTodo(config.GetString("goapi_api"), config.GetString("goapi_username"), config.GetString("goapi_password"), fmt.Sprint(items[i-1].ID), items[i-1])
				sort.Sort(todo.Order(items))
			} else {
				sort.Sort(todo.Order(items))
				todo.SaveTodos(viper.GetString("datafile"), items)
			}
		case config.GetString("editing_default") == "body" || bodyOpt:
			items[i-1].Body = createTemp([]byte(items[i-1].Body))
			if config.GetString("goapi_api") != "" {
				todo.UpdateRemoteTodo(config.GetString("goapi_api"), config.GetString("goapi_username"), config.GetString("goapi_password"), fmt.Sprint(items[i-1].ID), items[i-1])
				sort.Sort(todo.Order(items))
			} else {
				sort.Sort(todo.Order(items))
				todo.SaveTodos(viper.GetString("datafile"), items)
			}
		default:
			items[i-1].Title = createTemp([]byte(items[i-1].Title))
			if config.GetString("goapi_api") != "" {
				todo.UpdateRemoteTodo(config.GetString("goapi_api"), config.GetString("goapi_username"), config.GetString("goapi_password"), fmt.Sprint(items[i-1].ID), items[i-1])
				sort.Sort(todo.Order(items))
			} else {
				sort.Sort(todo.Order(items))
				todo.SaveTodos(viper.GetString("datafile"), items)
			}
		}
	} else {
		fmt.Printf("\"%v\" doesn't match any todos\n", i)
	}
}

func createTemp(todoText []byte) string {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "godo-")
	if err != nil {
		log.Fatal("Unable to create temporary file", err)
	}
	defer os.Remove(tmpFile.Name())
	if _, err := tmpFile.Write(todoText); err != nil {
		log.Fatal("Failed to write initial text to temporary file", err)
	}
	editTemp(tmpFile.Name())
	data, err := ioutil.ReadFile(tmpFile.Name())
	todo := string(data)
	todo = strings.TrimSuffix(todo, "\n")
	if err != nil {
		log.Println("File reading error", err)
	}
	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}
	return todo
}

func editTemp(filename string) error {
	editor := defaultEditor()
	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}
	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func defaultEditor() string {
	if config.GetString("editing_editor") == "" {
		editor := os.Getenv("EDITOR")
		return editor
	}
	editor := config.GetString("editing_editor")
	return editor
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	editCmd.Flags().BoolVarP(&titleOpt, "title", "t", false, "edit item title")
	editCmd.Flags().BoolVarP(&bodyOpt, "body", "b", false, "edit item body")
}
