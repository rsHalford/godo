/*
Package config creates the structure for the configuration of GoDo. This includes
setting the environment variables that the godo command and subcommands use to
operate according to the user's specification. These variables can also be set
using the config.toml file, that will be created automatically if it does not
already exist.

[general]
# change the file path for saving todos (defaults to "$HOME/.local/share/godo/godos.json" if unset)
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
*/
package config

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type General struct {
	DataFile string `toml:"data-file" env:"DATA_FILE"`
}

type Edit struct {
	Default  string `toml:"default" env:"DEFAULT"`
	Editor   string `toml:"editor" env:"EDITOR"`
	Filetype string `toml:"filetype" env:"FILETYPE"`
}

type Find struct {
	CaseSensitivity string `toml:"case-sensitivity" env:"CASE_SENSITIVITY"`
}

type Theme struct {
	Primary   string `toml:"primary" env:"PRIMARY" env-default:"Blue"`
	Secondary string `toml:"secondary" env:"SECONDARY" env-default:"Yellow"`
	Position  string `toml:"position" env:"POSITION" env-default:"Grey"`
	Tag       string `toml:"tag" env:"TAG" env-default:"Magenta"`
	Title     string `toml:"title" env:"TITLE" env-default:"BrightWhite"`
	Priority  string `toml:"priority" env:"PRIORITY" env-default:"Yellow"`
	Done      string `toml:"done" env:"DONE" env-default:"White"`
}

// Config struct defines the config.toml and related environment variables.
type Config struct {
	General General `toml:"general" env-prefix:"GODO_GENERAL_"`
	Edit    Edit    `toml:"edit" env-prefix:"GODO_EDIT_"`
	Find    Find    `toml:"find" env-prefix:"GODO_FIND_"`
	Theme   Theme   `toml:"theme" env-prefix:"GODO_THEME_"`
}

var cfg Config

// createCfgFile takes the configuration file path to determine if it exists,
// creating the file if missing.
func createCfgFile(cfgFile string) error {
	var f *os.File

	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("user home directory: %v", err)
	}

	// Set default directory for local godos.json.
	dataDir := home + "/.local/share/godo"

	// If cfgFile represents a filepath that does not exist on the system,
	// one will be created.
	if _, err := os.Stat(cfgFile); os.IsNotExist(err) {
		f, err = os.Create(cfgFile)
		if err != nil {
			return fmt.Errorf("creating %v: %w", cfgFile, err)
		}

		defer f.Close()

		// The file is created with boilerplate for configuration options.
		configBoilerplate := fmt.Sprintf(`[general]
# change the file path for saving todos (defaults to "%s/godos.json" if unset)
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
done = "white" # $GODO_THEME_DONE`, dataDir)

		_, err = f.WriteString(configBoilerplate)
		if err != nil {
			return fmt.Errorf("write to empty file: %w", err)
		}
	}

	return nil
}

// defineConfig assigns the file path for the configuration file, before
// checking the existence of the file and creating it if it doesn't exist.
func defineConfig() (cfgPath string, err error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("determining user configuration location: %v", err)
	}

	// Assign the config.toml filepath within the default root directory godo/,
	// to use in the user-specific configuration data.
	godoCfgDir := cfgDir + "/godo"
	cfgPath = godoCfgDir + "/config.toml"

	// Create the configuration directory if it doesn't already exist,
	// including a configuration file with acceptable options defined.
	if _, err = os.Stat(godoCfgDir); os.IsNotExist(err) {
		var perm fs.FileMode = 0o755

		err = os.Mkdir(godoCfgDir, perm)
		if err != nil {
			return "", fmt.Errorf("making new directory: %w", err)
		}

		err = createCfgFile(cfgPath)
		if err != nil {
			return "", fmt.Errorf("making new file: %w", err)
		}
	}

	// Create a configuration file if the directory is empty.
	err = createCfgFile(cfgPath)
	if err != nil {
		return "", fmt.Errorf("making new file: %w", err)
	}

	return cfgPath, nil
}

// Value takes a key and returns the matching value from the config.toml.
func Value(key string) string {
	cfgPath, err := defineConfig()
	if err != nil {
		fmt.Printf("finding configuration file: %v", err)
	}

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		fmt.Printf("parsing configuration: %v", err)
	}

	switch key {
	case "dataFile":
		value := cfg.General.DataFile

		return value

	case "editing_default":
		value := cfg.Edit.Default

		return value

	case "editing_editor":
		value := cfg.Edit.Editor

		return value

	case "editing_filetype":
		value := cfg.Edit.Filetype

		return value

	case "caseSensitivity":
		value := cfg.Find.CaseSensitivity

		return value

	default:
		fmt.Println("No configuration key provided")
	}

	return ""
}

// Theme takes a key and returns the matching value from the config.toml.
func ThemeValue(key string) string {
	cfgPath, err := defineConfig()
	if err != nil {
		fmt.Printf("finding configuration file: %v", err)
	}

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		fmt.Printf("parsing configuration: %v", err)
	}

	switch key {
	case "theme_primary":
		value := cfg.Theme.Primary

		return value

	case "theme_secondary":
		value := cfg.Theme.Secondary

		return value

	case "theme_position":
		value := cfg.Theme.Position

		return value

	case "theme_tag":
		value := cfg.Theme.Tag

		return value

	case "theme_title":
		value := cfg.Theme.Title

		return value

	case "theme_priority":
		value := cfg.Theme.Priority

		return value

	case "theme_done":
		value := cfg.Theme.Done

		return value

	default:
		fmt.Println("No configuration key provided")
	}

	return ""
}

// InitTheme retrieves all color values set in the config.toml file.
func InitTheme() *Theme {
	return &Theme{
		Primary:   ThemeValue("theme_primary"),
		Secondary: ThemeValue("theme_secondary"),
		Position:  ThemeValue("theme_position"),
		Tag:       ThemeValue("theme_tag"),
		Title:     ThemeValue("theme_title"),
		Priority:  ThemeValue("theme_priority"),
		Done:      ThemeValue("theme_done"),
	}
}
