# GoDo

## A command line todo list application

[![Go Reference](https://pkg.go.dev/badge/github.com/rsHalford/godo.svg)](https://pkg.go.dev/github.com/rsHalford/godo)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)

---

# About

GoDo aims to help you get organised with your tasks. Designed to be simple and accessible.

```
Usage:
  godo [command]

Available Commands:
  add         add a new todo
  completion  generate the autocompletion script for the specified shell
  done        toggle todo as done
  edit        edit a todo (default: edit title)
  find        search for a given string
  get         get a specific todo
  help        Help about any command
  list        list your todos
  priority    label a todo as a priority
  remove      remove a todo
  tag         add a tag to your todo
  version     print godo's version

Flags:
  -h, --help   help for godo

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

You can edit the `config.yaml` to set your preferred default settings, helping shorten your most used `godo` arguments. For example you might prefer to use Vim as your favourite terminal editor, with all your notes being done with markdown syntax.

On Linux this file will be read from `${XDG_CONFIG_HOME:-$HOME/.config}/godo/config.yaml`.

```yaml
general:
  # change the file path for saving local notes (defaults to "~/.local/share/godo/godos.json" if unset)
  dataFile: "~/.local/share/godo/godos.json"

# set preferences for editing notes
edit:
  # default to either editing the note title or body (defaults to "title" if unset)
  default: "body"
  # determine which editor to make edits in (defaults to the environment's $EDITOR if unset)
  editor: "vim"
  # append an extension to the temporary file's buffer for editing (e.g. "org", "md", "txt")
  filetype: "md"

find:
  # choose between "smart", "sensitive" or "insensitive" search patterns (defaults to "smart" if unset)
  # "smart" - if the search argument is all lower-case, all results are shown. Only becoming case-sensitive
  # if upper-case characters are provided.
  caseSensitivity: "smart"
```

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
