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
package todo

import (
	"encoding/json"
	"fmt"
	"github.com/rsHalford/godo/config"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Priority bool   `json:"priority"`
	position int
	Status   bool `json:"status"`
}

func SaveTodos(filename string, items []Todo) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetTodos() ([]Todo, error) {
	if config.GetString("goapi_api") != "" {
		items, err := GetRemoteTodos(config.GetString("goapi_api"), config.GetString("goapi_username"), config.GetString("goapi_password"))
		if err != nil {
			fmt.Print(err.Error())
		}
		return items, nil
	} else {
		dataFile := LocalTodos()
		items, err := ReadTodos(dataFile)
		if err != nil {
			fmt.Print(err.Error())
		}
		return items, nil
	}
}

func LocalTodos() (filename string) {
	if config.GetString("data_file") != "" {
		dataFile := config.GetString("data_file")
		return dataFile
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Println("Unable to detect home directory.")
		}
		dataFile := home + "/.local/share/godo/godos.json"
		return dataFile
	}
}

func ReadTodos(filename string) ([]Todo, error) {
	bodyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var items []Todo

	if err := json.Unmarshal(bodyBytes, &items); err != nil {
		log.Fatal(err)
	}

	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func (i *Todo) Prioritise(pri bool) {
	if pri {
		i.Priority = true
	}
}

func (i *Todo) PriorityFlag() (color string) {
	if i.Priority {
		return "\033[33m"
	}
	return
}

func (i *Todo) StatusFlag() (strike string) {
	if i.Status {
		return "\033[9m"
	}
	return
}

func (i *Todo) Label() string {
	return strconv.Itoa(i.position)
}

type Order []Todo

func (s Order) Len() int {
	return len(s)
}
func (s Order) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Order) Less(i, j int) bool {
	if s[i].Status == s[j].Status {
		if s[i].Priority == s[j].Priority {
			return s[i].position < s[j].position
		}
		return s[i].Priority && !s[j].Priority
	}
	return !s[i].Status
}
