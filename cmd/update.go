package cmd

import (
	"fmt"
	"sort"
	"time"

	"github.com/rsHalford/godo/todo"
)

// updateTodo will send the updated todo and it's properties to save locally.
func updateTodo(position int, command string, todos []todo.Todo) error {
	todos[position-1].UpdatedAt = time.Now().UTC()

	sort.Sort(todo.Order(todos)) // Sort the todos before saving.

	// Pass the filename of the local todo store to the filename variable.
	filename, err := todo.LocalTodos()
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	// Using SaveLocal to add the new todo(s) to the local JSON store.
	err = todo.SaveLocal(filename, todos)
	if err != nil {
		return fmt.Errorf("%v: %w", command, err)
	}

	return nil
}
