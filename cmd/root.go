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
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile, dataFile string
var version string = "v0.0.1"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "godo",
	Short: "GoDo is a todo application.",
	Long: `GoDo aims to help you get organised with your tasks. Designed to be 
simple and accessible. Go get things done and checked off the list.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	//cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initData)

	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile.")
	}

	//rootCmd.PersistentFlags().StringVar(&configFile, "config", home+"/.config/godo/config.yaml", "configuration file")
	rootCmd.PersistentFlags().StringVar(&dataFile, "datafile", home+"/.local/share/godo/godos.json", "data file to store todos")

	viper.BindPFlag("datafile", rootCmd.PersistentFlags().Lookup("datafile"))
}

func initData() {
	home, err := homedir.Dir()
	if err != nil {
		log.Println("Unable to detect home directory. Please set data file using --datafile.")
	}
	os.Mkdir(home+"/.local/share/godo", 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO: Configuration to be added when JSON API support is added
// initConfig reads in config file and ENV variables if set.
//func initConfig() {
//	if configFile != "" {
//		viper.SetConfigFile(configFile)
//	} else {
//		home, err := homedir.Dir()
//		cobra.CheckErr(err)
//		viper.AddConfigPath(home + "/.config/godo/")
//		viper.SetConfigName("config")
//	}
//	viper.SetEnvPrefix("godo")
//	viper.AutomaticEnv() // read in environment variables that match
//	viper.ReadInConfig()
//}
