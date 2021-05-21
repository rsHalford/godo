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
)

type Config struct {
	Username string `yaml:"username" env:"GODO_USERNAME"`
	Password string `yaml:"password" env:"GODO_PASSWORD"`
	API      string `yaml:"api" env:"GODO_API"`
	Editor   string `yaml:"editor" env:"GODO_EDITOR"`
}

var (
	cfg Config
	//cfgPath = "${XDG_CONFIG_HOME:-$HOME/.config}/godo/config.yaml"
	cfgPath = "./config.yaml"
)

func GetString(key string) string {
	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		fmt.Println(err)
	}

	switch key {
	case "api":
		value := cfg.API
		return value
	case "username":
		value := cfg.Username
		return value
	case "password":
		value := cfg.Password
		return value
	case "editor":
		value := cfg.Editor
		return value
	default:
		fmt.Println("No configuration key provided")
	}

	return ""
}
