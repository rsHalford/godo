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
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/rsHalford/godo/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:     "edit",
	Aliases: []string{"e"},
	Short:   "edit a todo",
	Long:    `Edit a todo by passing the list number of the todo.`,
	Run:     editRun,
}

func editRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	i, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatalln(args[0], "is not a valid label\n", err)
	}
	if i > 0 && i <= len(items) {
		items[i-1].Text = createTemp([]byte(items[i-1].Text))
		sort.Sort(todo.Order(items))
		todo.SaveItems(viper.GetString("datafile"), items)
	} else {
		log.Println(i, "doesn't match any items")
	}
}

func createTemp(todoText []byte) string {
	tmpFile, err := ioutil.TempFile(os.TempDir(), "godo-")
	if err != nil {
		log.Fatal("Unable to create temporary file", err)
	}
	defer os.Remove(tmpFile.Name())
	log.Println("Created temporary file: " + tmpFile.Name())
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

//const defaultEditor = "nvim"

func editTemp(filename string) error {
	editor := os.Getenv("EDITOR")
	//if editor == "" {
	//	editor = defaultEditor
	//}
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

//func editorInput() ([]byte, error) {
//	file, err := ioutil.TempFile(os.TempDir(), "*")
//	if err != nil {
//		return []byte{}, err
//	}
//	filename := file.Name()
//	defer os.Remove(filename)
//	if err := file.Close(); err != nil {
//		return []byte{}, err
//	}
//	if err = openEditor(filename); err != nil {
//		return []byte{}, err
//	}
//	bytes, err := ioutil.ReadFile(filename)
//	if err = openEditor(filename); err != nil {
//		return []byte{}, err
//	}
//	return bytes, nil
//}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
