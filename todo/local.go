package todo

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"

	"github.com/rsHalford/godo/config"
)

const perm fs.FileMode = 0o600

// LocalTodos will assign dataFile a user defined filename for the local godos.json.
// Creating one within the godo data directory if not set if config.yaml.
func LocalTodos() (filename string, err error) {
	// If dataFile configuration key is set, return that value.
	if config.Value("dataFile") != "" {
		dataFile := config.Value("dataFile")

		return dataFile, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("user home directory: %w", err)
	}

	// Use default value for dataFile if configuration key is not set.
	dataFile := home + "/.local/share/godo/godos.json"

	return dataFile, nil
}

// ReadLocal reads the contents of the file provided. Parsing the items
// from the JSON, and assigning a position value to each item.
func ReadLocal(filename string) (items []Todo, err error) {
	// Open the godos.json file with the given path for decoding.
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("reading %v: %w", filename, err)
	}

	defer f.Close()

	// The JSON-encoded data of the file is then parsed and stored
	// within the matching values of items' Todo struct.
	err = json.NewDecoder(f).Decode(&items)
	if err != nil {
		return nil, fmt.Errorf("parsing JSON: %w", err)
	}

	// Each item, i is then assigned a position value
	// that is greater than the previous by 1.
	for i := range items {
		items[i].position = i + 1
	}

	return
}

// SaveLocal encodes the todo items as JSON objects,
// before writing the changes to the provided file.
func SaveLocal(filename string, items []Todo) error {
	// The todo items are encoded as JSON objects, relating to the Todo struct.
	data, err := json.Marshal(items)
	if err != nil {
		return fmt.Errorf("encoding JSON: %w", err)
	}

	// The data is then written to the provided file.
	// Creating one with the provided permissions, if it doesn't exist.
	err = os.WriteFile(filename, data, perm)
	if err != nil {
		return fmt.Errorf("writing to %v: %w", filename, err)
	}

	return nil
}
