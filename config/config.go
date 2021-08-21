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
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	GOAPI struct {
		API      string `yaml:"api" env:"GODO_GOAPI_API"`
		Password string `yaml:"password" env:"GODO_GOAPI_PASSWORD"`
		Username string `yaml:"username" env:"GODO_GOAPI_USERNAME"`
	} `yaml:"goapi" env:"GODO_GOAPI"`
	Editing struct {
		Default  string `yaml:"default" env:"GODO_EDITING_DEFAULT"`
		Editor   string `yaml:"editor" env:"GODO_EDITING_EDITOR"`
	} `yaml:"editing" env:"GODO_EDITING"`
	GUI struct {
		Port	string `yaml:"port" env:"GODO_GUI_PORT"`
	} `yaml:"gui" env:"GODO_GUI"`
}

var cfg Config

func GetString(key string) string {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println(err.Error())
	}
	cfgPath := cfgDir + "/godo/config.yaml"
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		return ""
	}

	switch key {
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
	case "gui_port":
		value := cfg.GUI.Port
		return value
	default:
		fmt.Println("No configuration key provided")
	}

	return ""
}
