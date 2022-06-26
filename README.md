<p align="center">
  <img width="80%" src="godo-full.svg" alt="GoDo: A command line todo list application.">
</p>

<p align="center">
  <a href="https://pkg.go.dev/github.com/rsHalford/godo">
    <img src="https://pkg.go.dev/badge/github.com/rsHalford/godo.svg" alt="Go Reference">
  </a>
  <a href="https://github.com/pre-commit/pre-commit">
    <img src="https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white" alt="pre-commit">
  </a>
  <a href="https://goreportcard.com/report/github.com/rsHalford/godo">
    <img src="https://goreportcard.com/badge/github.com/rsHalford/godo" alt="Go Report Card">
  </a>
</p>

---

# About

GoDo aims to help you get organised with your todos. Designed to be simple and accessible.

```
Usage:
  godo [command]

Available Commands:
  add         Add a new todo
  completion  Generate the autocompletion script for the specified shell
  done        Toggle todo as done
  edit        Edit a todo (default: edit title)
  find        Search for a given string
  get         Get a specific todo
  help        Help about any command
  list        List todos
  priority    Label a todo as a priority
  remove      Remove a todo
  tag         Change the tag of a todo

Flags:
  -h, --help      help for godo
  -v, --version   version for godo

Use "godo [command] --help" for more information about a command.
```

Go get things done and checked off the list.

---

# Getting Started

## Requirements

All you need is to have Go [installed](https://go.dev/dl/) to compile GoDo.

## Installation

To install GoDo, all you have to do is run the `go install` command.

```sh
$ go install github.com/rsHalford/godo@latest
```

### Flake

If you're into Nix, GoDo has been setup so that you can just add the repo's URL to your flake.nix inputs. Then the overlay can be called by `environment.systemPackages` or per user with `home.packages`.

```nix
inputs = {
  nixpkgs.url = "nixpkgs/nixos-unstable";

  godo-flake = {
    url = "github:rsHalford/godo";
    inputs.nixpkgs.follows = "nixpkgs";
  };
};
```

Instructions on how to add GoDo as a flake can vary wildly depending on the system configuration. You can view [my dotfiles](https://github.com/rsHalford/dotfiles) to see a working example, and message me or create an issue on either repository if you need additional help getting setup.

## Configuration

You can edit the `config.toml` to set your preferred default settings, helping shorten your most used `godo` arguments. For example you might prefer to use Vim as your favourite terminal editor, with all your notes being done with markdown syntax.

On Linux this file will be read from `${XDG_CONFIG_HOME:-$HOME/.config}/godo/config.toml`.

```toml
[general]
# change the file path for saving todos (defaults to "~/.local/share/godo/godos.json" if unset)
data-file = "" # $GODO_GENERAL_DATA_FILE

[edit]
# default to editing the todo title or body (defaults to "title" if unset)
default = "body" # $GODO_EDIT_DEFAULT
# which editor to make edits in (defaults to the environment's $EDITOR if unset)
editor = "vim" # $GODO_EDIT_EDITOR
# append an extension to the temporary file's buffer for editing (e.g. "org", "md", "txt")
filetype = "md" # $GODO_EDIT_FILETYPE

[find]
# "smart", "sensitive" or "insensitive" search patterns (defaults to "smart" if unset)
# "smart" - if the search argument is all lower-case, all results are shown.
# Only becoming case-sensitive if upper-case characters are provided.
case-sensitivity = "smart" # $GODO_FIND_CASE_SENSITIVITY

[theme]
# use case-insesitive color names or hexadecimals, and prepend with "bg" to change the background instead.
primary = "bg#00385c" # $GODO_THEME_PRIMARY
secondary = "#00add8" # $GODO_THEME_SECONDARY
position = "grey" # $GODO_THEME_POSITION
tag = "magenta" # $GODO_THEME_TAG
title = "brightwhite" # $GODO_THEME_TITLE
priority = "yellow" # $GODO_THEME_PRIORITY
done = "white" # $GODO_THEME_DONE
```

## NO_COLOR Support

GoDo accepts the following methods to change no/color output:

To disable color:
- stdout is not a TTY
- NO_COLOR=true
- TERM=dumb
- FORCE_COLOR=0

To enable:
- FORCE_COLOR=1, 2, or 3 (16 color, 256 color and 16.7m color respectively)
- FORCE_COLOR=true or no value

---

# Contributing

To help contribute to GoDo, you could either send in your feature requests as an issue or take it upon yourself to send in a pull request after following the [CONTRIBUTING](https://github.com/rsHalford/godo/blob/main/CONTRIBUTING.md) guide.

My long-term aims are actually to re-implement the more "advanced/unneeded/complex" features like the GUI, API access and mobile way to interact with GoDo. Potentially making them available as separate plugins that can be added with the users discretion, rather than ship GoDo with everything already included.

Thanks in advance for taking an interest!

## Development Environment

GoDo development is setup to utilise the following tools;

- [git-chglog](https://github.com/git-chglog/git-chglog)
  - For updating the CHANGELOG.md
- [golangci-lint](https://golangci-lint.run/)
  - Go linters aggregator
- [gopls](https://github.com/golang/tools/blob/master/gopls/README.md)
  - Go language server
- [pre-commit](https://pre-commit.com/)
  - Setup and run pre-commit hooks

### [direnv](https://direnv.net/)

Combined with nix flakes and [nix-direnv](https://github.com/nix-community/nix-direnv), it is possible to create a `.envrc` file to make a portable isolated environment for development.

Meaning the above tools can be installed and setup just for this project by just using the below file contents.

```sh
use flake
layout go
```

---

# Licence

GoDo is released under the GNU General Public License v3.0.

See [LICENSE](https://github.com/rsHalford/godo/blob/main/LICENSE).
