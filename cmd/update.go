package cmd

import (
	"fmt"
	"sort"
	"time"

	"github.com/rsHalford/godo/todo"
)

// updateTodo will send the updated item and it's properties to save locally.
func updateTodo(i int, command string, items []todo.Todo) error {
	items[i-1].UpdatedAt = time.Now().UTC()

	sort.Sort(todo.Order(items)) // Sort the items before saving.

	// Pass the filename of the local todo store to the filename variable.
	filename, err := todo.LocalTodos()
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	// Using SaveLocal to add the new todo(s) to the local JSON store.
	err = todo.SaveLocal(filename, items)
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	return nil
}
