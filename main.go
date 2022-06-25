/*
Godo is a command line todo list application.

Usage:

	godo [command]

Available Commands:

	add         Add new todo
	completion  Generate the autocompletion script for the specified shell
	done        Toggle todo as done
	edit        Edit a todo (default: edit title)
	find        Search for a given string
	get         Get a specific todo
	help        Help about any command
	list        List your todos
	priority    Label a todo as a priority
	remove      Remove a todo
	tag         Add a tag to your todo

Flags:

	-h, --help      help for godo
	-v, --version   version for godo

Use "godo [command] --help" for more information about a command.
*/
package main

import "github.com/rsHalford/godo/cmd"

func main() {
	cmd.Execute()
}
