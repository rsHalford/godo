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
package config

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type General struct {
	DataFile string `yaml:"dataFile" env:"DATA_FILE"`
}

type Edit struct {
	Default  string `yaml:"default" env:"DEFAULT"`
	Editor   string `yaml:"editor" env:"EDITOR"`
	Filetype string `yaml:"filetype" env:"FILETYPE"`
}

type Find struct {
	CaseSensitivity string `yaml:"caseSensitivity" env:"CASE_SENSITIVITY"`
}

// Config struct defines the config.yaml and related environment variables.
type Config struct {
	General General `yaml:"general" env-prefix:"GODO_GENERAL_"`
	Edit    Edit    `yaml:"edit" env-prefix:"GODO_EDIT_"`
	Find    Find    `yaml:"find" env-prefix:"GODO_FIND_"`
}

var cfg Config

// createCfgFile takes the configuration file path to determine if it exists,
// creating the file if missing.
func createCfgFile(cfgFile string) error {
	var f *os.File

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("user home directory: %v", err)
	}

	// Set default directory for local godos.json.
	dataDir := home + "/.local/share/godo"

	// If cfgFile represents a filepath that does not exist on the system,
	// one will be created.
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		f, err = os.Create(cfgFile)
		if err != nil {
			return fmt.Errorf("creating %v: %w", cfgFile, err)
		}

		defer f.Close()

		// The file is created with boilerplate for configuration options.
		configBoilerplate := fmt.Sprintf(`general:
  # change the file path for saving local notes (defaults to "%s" if unset)
  # $GODO_GENERAL_DATA_FILE
  dataFile: ""

# set preferences for editing notes
edit:
  # default to either editing the note title or body (defaults to "title" if unset)
  # $GODO_EDIT_DEFAULT
  default: "body"
  # determine which editor to make edits in (defaults to the environment's $EDITOR if unset)
  # $GODO_EDIT_EDITOR
  editor: "vim"
  # append an extension to the temporary file's buffer for editing (e.g. "org", "md", "txt")
  # $GODO_EDIT_FILETYPE
  filetype: "md"

find:
  # choose between "smart", "sensitive" or "insensitive" search patterns (defaults to "smart" if unset)
  # "smart" - if the search argument is all lower-case, all results are shown. Only becoming case-sensitive
  # if upper-case characters are provided.
  # $GODO_FIND_CASE_SENSITIVITY
  caseSensitivity: "smart"`, dataDir)

		_, err = f.WriteString(configBoilerplate)
		if err != nil {
			return fmt.Errorf("write to empty file: %w", err)
		}
	}

	return nil
}

// defineConfig assigns the file path for the configuration file, before checking
// the existence of the file and creating it if it doesn't exist.
func defineConfig() (cfgPath string, err error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("determining user configuration location: %v", err)
	}

	// Assign the config.yaml filepath within the default root
	// directory godo/, to use in the user-specific configuration data.
	godoCfgDir := cfgDir + "/godo"
	cfgPath = godoCfgDir + "/config.yaml"

	// Create the configuration directory if it doesn't already exist,
	// including a configuration file with acceptable options defined.
	if _, err = os.Stat(godoCfgDir); os.IsNotExist(err) {
		var perm fs.FileMode = 0o755

		err = os.Mkdir(godoCfgDir, perm)
		if err != nil {
			return "", fmt.Errorf("making new directory: %w", err)
		}

		err = createCfgFile(cfgPath)
		if err != nil {
			return "", fmt.Errorf("making new file: %w", err)
		}
	}

	// Create a configuration file if the directory is empty.
	err = createCfgFile(cfgPath)
	if err != nil {
		return "", fmt.Errorf("making new file: %w", err)
	}

	return cfgPath, nil
}

// Value takes a key and returns the matching value from the config.yaml.
func Value(key string) string {
	cfgPath, err := defineConfig()
	if err != nil {
		fmt.Printf("finding configuration file: %v", err)
	}

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		fmt.Printf("parsing configuration: %v", err)
	}

	switch key {
	case "dataFile":
		value := cfg.General.DataFile

		return value

	case "editing_default":
		value := cfg.Edit.Default

		return value

	case "editing_editor":
		value := cfg.Edit.Editor

		return value

	case "editing_filetype":
		value := cfg.Edit.Filetype

		return value

	case "caseSensitivity":
		value := cfg.Find.CaseSensitivity

		return value

	default:
		fmt.Println("No configuration key provided")
	}

	return ""
}
