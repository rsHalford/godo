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
	"github.com/spf13/viper"
)

var (
	defaults = map[string]interface{}{
		"username": "",
		"password": "",
		"api":      "",
		"editor":   "",
	}
	configName     = "config"
	configType     = "yaml"
	configPath     = "."
	configPathFull = configPath + "/" + configName + "." + configType
)

type Config struct {
	Username string
	Password string
	Api      string
	Editor   string
}

func InitConfig() {
	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.SafeWriteConfigAs(configPathFull)
	viper.ReadInConfig()
}
