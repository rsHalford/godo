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

	c "github.com/jwalton/gchalk"
	"github.com/rsHalford/godo/config"
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

var Theme *config.Theme

func init() {
	Theme = config.InitTheme() // Load in colour theme.
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

// Prioritise will set the priority value of the Todo item as true.
func (i *Todo) Prioritise(priority bool) {
	if priority {
		i.Priority = true
	}
}

// TitleFmt will return back the given title string with the appropriate color
// and styling, according to it's priority and status.
func (i *Todo) TitleFmt(s string) string {
	switch {
	case i.Priority && i.Status:
		s = c.WithStyleMust(Theme.Priority).Strikethrough(s) + "\t"
		return s

	case i.Priority && !i.Status:
		s = c.StyleMust(Theme.Priority)(s) + "\t"
		return s

	case i.Status && !i.Priority:
		s = c.WithStyleMust(Theme.Title).Strikethrough(s) + "\t"
		return s

	default:
		s = s + "\t"
		return s
	}
}

// TagFmt first determines whether the tag string is empty, and adds a space
// to fix formatting alignments. Before returning it with color and style
// formatting.
func (i *Todo) TagFmt(s string) string {
	if i.Tag == "" {
		s = " "
	}
	s = c.WithItalic().StyleMust(Theme.Tag)(s) + "\t"
	return s
}

// Label will convert the position integer value of the item to a string. Then
// color and style, before returning.
func (i *Todo) Label() (s string) {
	s = strconv.Itoa(i.position)
	s = c.StyleMust(Theme.Position)(s) + "\t"
	return s
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
