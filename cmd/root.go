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
var rootCmd = &cobra.Command{
	Use:   "godo",
	Short: "GoDo is a todo application.",
	Long: `
GoDo

A command line todo list application

GoDo aims to help you get organised with your tasks.
Designed to be simple and accessible.

Go get things done and checked off the list.

========================================================================`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var dataFile = config.GetString("dataFile")

const perm fs.FileMode = 0o755

func init() {
	cobra.OnInitialize(initData)
}

func initData() {
	if config.GetString("goapi_api") == "" && dataFile == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("user home directory: %v", err)
		}

		godoDirectory := home + "/.local/share/godo"

		if _, err := os.Stat(godoDirectory); os.IsNotExist(err) {
			err := os.Mkdir(godoDirectory, perm)
			if err != nil {
				fmt.Printf("making new directory: %v", err)
			}
		}
	}
}
