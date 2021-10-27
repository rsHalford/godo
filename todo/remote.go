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

// CreateRemote will encode the todo item into JSON.
// Then make a POST request to the API, using basic authentication.
func CreateRemote(url, username, password string, item Todo) error {
	// The todo item is encoded as JSON objects, relating to the Todo struct.
	data, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("encoding JSON: %w", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("POST requesting data: %w", err)
	}

	// Assign the basic authentication credentials to the request.
	req.SetBasicAuth(username, password)

	// Send the HTTP request, and assign the response to resp.
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("data response: %w", err)
	}

	resp.Body.Close()

	return nil
}

// RemoteTodos sends a request to the API url with basic authentication.
// And makes a GET request, which returns a response body. That is read by
// parsing the items from the JSON, and assigning a position value to each item.
func RemoteTodos(url, username, password string) ([]Todo, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("GET requesting data: %w", err)
	}

	// Assign the basic authentication credentials to the request.
	req.SetBasicAuth(username, password)

	// Send the HTTP request, and assign the response to resp.
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("data response: %w", err)
	}

	defer resp.Body.Close()

	// bodyBytes holds the contents of the file within a []byte.
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response data: %w", err)
	}

	var items []Todo

	// The JSON-encoded data of bodyBytes is then parsed and stored
	// within the matching values from the Todo struct.
	if err := json.Unmarshal(bodyBytes, &items); err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	// Each item, i is then assigned a position value
	// that is greater than the previous by 1.
	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}

// UpdateRemote will take in the todo and encode it into JSON. Before
// sending a PUT request to the API's url, using the todo id,
// to send the updated todo data.
func UpdateRemote(url, username, password, id string, todo Todo) error {
	// The todo item are encoded as JSON objects, relating to the Todo struct.
	data, err := json.Marshal(todo)
	if err != nil {
		return fmt.Errorf("encoding JSON: %w", err)
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodPut, url+"/"+id, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("PUT requesting data: %w", err)
	}

	// Assign the basic authentication credentials to the request.
	req.SetBasicAuth(username, password)

	// Send the HTTP request, and assign the response to resp.
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("data response: %w", err)
	}

	resp.Body.Close()

	return nil
}

// DeleteRemote uses the todo id to send a DELETE request
// to the specific API url, using basic authentication.
func DeleteRemote(url, username, password, id string) error {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodDelete, url+"/"+id, nil)
	if err != nil {
		return fmt.Errorf("DELETE requesting data: %w", err)
	}

	// Assign the basic authentication credentials to the request.
	req.SetBasicAuth(username, password)

	// Send the HTTP request, and assign the response to resp.
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("data response: %w", err)
	}

	resp.Body.Close()

	return nil
}
