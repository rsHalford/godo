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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateRemoteTodo(url, username, password string, item Todo) error {
	data, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("encoding JSON: %w", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("POST requesting data: %w", err)
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("data response: %w", err)
	}

	resp.Body.Close()

	return nil
}

func GetRemoteTodos(url, username, password string) ([]Todo, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("GET requesting data: %w", err)
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("data response: %w", err)
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response data: %w", err)
	}

	var items []Todo

	if err := json.Unmarshal(bodyBytes, &items); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func UpdateRemoteTodo(url, username, password, id string, todo Todo) error {
	data, err := json.Marshal(todo)
	if err != nil {
		return fmt.Errorf("encoding JSON: %w", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, url+"/"+id, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("PUT requesting data: %w", err)
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("data response: %w", err)
	}

	resp.Body.Close()

	return nil
}

func DeleteRemoteTodo(url, username, password, id string) error {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, url+"/"+id, nil)
	if err != nil {
		return fmt.Errorf("DELETE requesting data: %w", err)
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("data response: %w", err)
	}

	resp.Body.Close()

	return nil
}
