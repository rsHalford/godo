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
	"io/ioutil"
	"strconv"
)

type Item struct {
	Text     string
	Priority bool
	position int
	Done     bool
}

func SaveItems(filename string, items []Item) error {
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

func ReadItems(filename string) ([]Item, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	for i := range items {
		items[i].position = i + 1
	}
	return items, nil
}

func (i *Item) Prioritise(pri bool) {
	if pri {
		i.Priority = true
	}
}

func (i *Item) PretPriority() string {
	if i.Priority {
		return "*"
	}
	return " "
}

func (i *Item) PretDone() string {
	if i.Done {
		return "X"
	}
	return ""
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position)
}

type Order []Item

func (s Order) Len() int {
	return len(s)
}
func (s Order) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s Order) Less(i, j int) bool {
	if s[i].Done == s[j].Done {
		if s[i].Priority == s[j].Priority {
			return s[i].position < s[j].position
		}
		return s[i].Priority && !s[j].Priority
	}
	return !s[i].Done
}
