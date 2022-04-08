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
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config struct defines the config.yaml and related environment variables.
type Config struct {
	GENERAL struct {
		DATA_FILE string `yaml:"dataFile" env:"GODO_GENERAL_DATA_FILE"`
	} `yaml:"general"`
	GOAPI struct {
		API      string `yaml:"api" env:"GODO_GOAPI_API"`
		Password string `yaml:"password" env:"GODO_GOAPI_PASSWORD"`
		Username string `yaml:"username" env:"GODO_GOAPI_USERNAME"`
	} `yaml:"goapi"`
	Editing struct {
		Default  string `yaml:"default" env:"GODO_EDITING_DEFAULT"`
		Editor   string `yaml:"editor" env:"GODO_EDITING_EDITOR"`
		Filetype string `yaml:"filetype" env:"GODO_EDITING_FILETYPE"`
	} `yaml:"editing"`
	GUI struct {
		Port string `yaml:"port" env:"GODO_GUI_PORT"`
	} `yaml:"gui"`
}

var cfg Config

// Value takes a key and returns the matching value from the config.yaml.
func Value(key string) string {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("determining user configuration location: %v", err)
	}

	// Assign the config.yaml filepath within the default root
	// directory, to use for user-specific configuration data.
	cfgPath := cfgDir + "/godo/config.yaml"

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		fmt.Printf("parsing configuration: %v", err)
	}

	switch key {
	case "dataFile":
		value := cfg.GENERAL.DATA_FILE

		return value

	case "goapi_api":
		value := cfg.GOAPI.API

		return value

	case "goapi_password":
		value := cfg.GOAPI.Password

		return value

	case "goapi_username":
		value := cfg.GOAPI.Username

		return value

	case "editing_default":
		value := cfg.Editing.Default

		return value

	case "editing_editor":
		value := cfg.Editing.Editor

		return value

	case "editing_filetype":
		value := cfg.Editing.Filetype

		return value

	case "gui_port":
		value := cfg.GUI.Port

		return value

	default:
		fmt.Println("No configuration key provided")
	}

	return ""
}
