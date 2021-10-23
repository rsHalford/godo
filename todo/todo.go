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
	"io/fs"
	"os"
	"strconv"

	"github.com/rsHalford/godo/config"
)

type Todo struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Priority bool   `json:"priority"`
	position int
	Status   bool `json:"status"`
}

const perm fs.FileMode = 0o600

func SaveTodos(filename string, items []Todo) error {
	data, err := json.Marshal(items)
	if err != nil {
		return fmt.Errorf("encoding JSON: %w", err)
	}

	err = os.WriteFile(filename, data, perm)
	if err != nil {
		return fmt.Errorf("writing to %v: %w", filename, err)
	}

	return nil
}

func GetTodos() ([]Todo, error) {
	if config.GetString("goapi_api") != "" {
		items, err := GetRemoteTodos(config.GetString("goapi_api"), config.GetString("goapi_username"), config.GetString("goapi_password"))
		if err != nil {
			return nil, fmt.Errorf("fetching remote todos: %w", err)
		}

		return items, nil
	}

	dataFile, err := LocalTodos()
	if err != nil {
		return nil, fmt.Errorf("local filepath: %w", err)
	}

	if _, err = os.Stat(dataFile); os.IsNotExist(err) {
		var f *os.File

		f, err = os.Create(dataFile)
		if err != nil {
			return nil, fmt.Errorf("creating %v: %w", dataFile, err)
		}

		defer f.Close()

		_, err = f.Write([]byte("[]"))
		if err != nil {
			return nil, fmt.Errorf("write to empty file: %w", err)
		}
	}

	items, err := ReadTodos(dataFile)
	if err != nil {
		return nil, fmt.Errorf("reading from %v: %w", dataFile, err)
	}

	return items, nil
}

func LocalTodos() (filename string, err error) {
	if config.GetString("dataFile") != "" {
		dataFile := config.GetString("dataFile")

		return dataFile, nil
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("user home directory: %w", err)
		}

		dataFile := home + "/.local/share/godo/godos.json"

		return dataFile, nil
	}
}

func ReadTodos(filename string) (items []Todo, err error) {
	bodyBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading %v: %w", filename, err)
	}

	err = json.Unmarshal(bodyBytes, &items)
	if err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	for i := range items {
		items[i].position = i + 1
	}

	return
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

func (i *Todo) Label() (position string) {
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
