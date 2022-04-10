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
	"fmt"
	"os"
	"strconv"
	"time"
)

// Todo struct defines the key:value pair types and JSON layout.
type Todo struct {
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Tag       string    `json:"tag"`
	Priority  bool      `json:"priority"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	position  int
}

// Todos determines where to retrieve todo data from locally by using the user
// defined configuration file.
func Todos() ([]Todo, error) {
	// The dataFile variable is assigned using the LocalTodos() function.
	dataFile, err := LocalTodos()
	if err != nil {
		return nil, fmt.Errorf("local filepath: %w", err)
	}

	// If dataFile represents a filepath that does not exist on the system,
	// one will be created.
	if _, err = os.Stat(dataFile); os.IsNotExist(err) {
		var f *os.File

		f, err = os.Create(dataFile)
		if err != nil {
			return nil, fmt.Errorf("creating %v: %w", dataFile, err)
		}

		defer f.Close()

		// The file is created with empty square-brackets.
		// To be read successfully as an empty JSON file.
		_, err = f.Write([]byte("[]"))
		if err != nil {
			return nil, fmt.Errorf("write to empty file: %w", err)
		}
	}

	// The contents of the local file will have it's items parsed,
	// and a position value for ordering.
	items, err := ReadLocal(dataFile)
	if err != nil {
		return nil, fmt.Errorf("reading from %v: %w", dataFile, err)
	}

	return items, nil
}

// Tag will add the tag string for the Todo item.
func (i *Todo) Tagging(tag string) {
	if tag != "" {
		i.Tag = tag
	}
}

// TagFlag will return a magenta foreground and italic ANSI escape codes,
// if an item has a tag.
func (i *Todo) TagFlag() (color string) {
	return "\033[3m\033[35m"
}

// Prioritise will set the priority value of the Todo item as true.
func (i *Todo) Prioritise(priority bool) {
	if priority {
		i.Priority = true
	}
}

// PriorityFlag will return a yellow foreground ANSI escape code,
// if an item set as a priority.
func (i *Todo) PriorityFlag() (color string) {
	if i.Priority {
		return "\033[33m"
	}

	return
}

// StatusFlag will return a strike-through ANSI escape code,
// if an item set as a done.
func (i *Todo) StatusFlag() (strike string) {
	if i.Status {
		return "\033[9m"
	}

	return
}

// Label will convert the position integer value of the item to a string.
// To allow the command-line interface to print it's value.
func (i *Todo) Label() (position string) {
	return strconv.Itoa(i.position)
}

// Order helps sort to organise the todo items for printing.
// Items are separated by their status, then priority,
// and then finally in ascending position order.
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
