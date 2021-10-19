# GoDo

## A command line todo list application

[![Go Reference](https://pkg.go.dev/badge/github.com/rsHalford/godo.svg)](https://pkg.go.dev/github.com/rsHalford/godo)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)

---

# Table of Contents

- [About](#about)
- [Getting Started](#getting-started)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [Configuration](#configuration)
- [Licence](#licence)

---

# About

GoDo aims to help you get organised with your tasks. Designed to be simple and accessible.

```sh
Usage:
  godo [command]

Available Commands:
  add         add a new todo
  done        toggle todo as done
  edit        edit a todo
  find        search for a given string
  help        help about any command
  list        list your todos
  priority    label a todo as a priority
  remove      remove a todo
  version     print godo's version

Flags:
      --datafile string   data file to store todos
  -h, --help              help for godo

Use "godo [command] --help" for more information about a command.
```

### Go get things done and checked off the list.

---

# Getting Started

## Requirements

The best way to ensure GoDo will work on your machine, is to compile it yourself.

- Go (to compile applications)

To do this, all you need is to have Go - [the programming language](https://golang.org/doc/install) - installed on your computer.

## Installation

To install GoDo, all you have to do is run the `go get` command.

```sh
$ go get -u github.com/rsHalford/godo
```

## Configuration

The `config.yaml` can be edited to set-up a connection to [GoAPI](https://github.com/rsHalford/goapi). As well as, select your preferred text editor.

> This file will read from `${XDG_CONFIG_HOME:-$HOME/.config}/godo/config.yaml`.

```yaml
username: "admin"
password: "secret"
api: "https://example.com/api/v1/todo"
editor: "vim"
```

If you prefer to use a local JSON file to store your todo list, leave the api address blank - `""`.

By default, GoDo will use whatever text editor you have set as your system's default - `$EDITOR`.

---

# Licence

GoDo is released under the GNU General Public License v3.0.

👉 See [LICENSE](https://github.com/rsHalford/godo/blob/main/LICENSE).
