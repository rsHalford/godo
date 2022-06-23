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
	"io/fs"
	"os"

	"github.com/rsHalford/godo/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
// TODO: remove header to an info/man page instead, or just when help is called.
var rootCmd = &cobra.Command{
	Use:   "godo",
	Short: "GoDo is a todo application.",
	Long:  "GoDo: A command line todo list application",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

const perm fs.FileMode = 0o755

var Theme *config.Theme

func init() {
	Theme = config.InitTheme() // Load in colour theme.
	cobra.OnInitialize(initData)
}

// initData checks the configuration file for a location to save todos.
// Creating a godo directory in the user's home directory, if not set.
func initData() {
	// Check config.yaml for user defined local filepath.
	if config.Value("dataFile") == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("user home directory: %v", err)
		}

		// Set default directory for local godos.json.
		godoDirectory := home + "/.local/share/godo"

		// Create godo directory if it doesn't already exist.
		if _, err := os.Stat(godoDirectory); os.IsNotExist(err) {
			err := os.Mkdir(godoDirectory, perm)
			if err != nil {
				fmt.Printf("making new directory: %v", err)
			}
		}
	}
}
