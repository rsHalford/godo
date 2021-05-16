GoDo is a command line todo list application, written in Go.


# Table of Contents

- [About](#about)
- [Getting Started](#getting-started)
  - [Requirements](#requirements)
  - [Install](#install)
- [Database](#database)
  - [Server API](#server-api)
- [Licence](#licence)

# About

GoDo aims to help you get organised with your tasks. Designed to be 
simple and accessible.

Providing the ability to:

- Add - `godo add [todo text]`
- Edit - `godo edit [todo id]`
- Remove - `godo remove [todo id]`
- List - `godo list`
- Search - `godo find [todo text]`
- Prioritise - `godo prioritise [todo id]`
- Mark complete - `godo done [todo id]`

Go get things done and checked off the list.


# Getting Started

## Requirements

- Go

## Install

To install GoDo, all you have to do is run `go get` to install the latest version.

```sh
go get -u github.com/rsHalford/godo
```

## Database

By default, GoDo will create a JSON file to store your todo list.

### Server API

The config.yaml can be edited to set-up a connection to a [RESTful API](https://github.com/rsHalford/goapi)

```
go get -u github.com/rsHalford/goapi
```

Just make sure to edit the configuration file, to provide the username, password and base API URL.

Once a connection is setup, this will now be the default store of data for your todo list.

# Licence

GoDo is released under the GNU General Public License v3.0. See [LICENSE](https://github.com/rsHalford/godo/LICENSE)
