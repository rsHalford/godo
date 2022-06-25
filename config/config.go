/*
Package config creates the structure for the configuration of GoDo. This includes
setting the environment variables that the godo command and subcommands use to
operate according to the user's specification. These variables can also be set
using the config.yaml file, that will be created automatically if it does not
already exist.

general:
  # change the file path for saving local notes (defaults to "~/.local/share/godo/godos.json" if unset)
  # $GODO_GENERAL_DATA_FILE
  dataFile: ""

# set preferences for editing notes
edit:
  # default to either editing the note title or body (defaults to "title" if unset)
  # $GODO_EDIT_DEFAULT
  default: "body"
  # determine which editor to make edits in (defaults to the environment's $EDITOR if unset)
  # $GODO_EDIT_EDITOR
  editor: "vim"
  # append an extension to the temporary file's buffer for editing (e.g. "org", "md", "txt")
  # $GODO_EDIT_FILETYPE
  filetype: "md"

find:
  # choose between "smart", "sensitive" or "insensitive" search patterns (defaults to "smart" if unset)
  # "smart" - if the search argument is all lower-case, all results are shown. Only becoming case-sensitive
  # if upper-case characters are provided.
  # $GODO_FIND_CASE_SENSITIVITY
  caseSensitivity: "smart"

# change the colour of the output
theme:
  # use case-insesitive color names or hexadecimals, and prepend with "bg" to change the background instead.
  # $GODO_THEME_PRIMARY
  primary: "bg#00385c"
  # $GODO_THEME_SECONDARY
  secondary: "#00add8"
  # $GODO_THEME_POSITION
  position: "grey"
  # $GODO_THEME_TAG
  tag: "magenta"
  # $GODO_THEME_TITLE
  title: "brightwhite"
  # $GODO_THEME_PRIORITY
  priority: "yellow"
  # $GODO_THEME_STATUS
  status: "white"
*/
package config

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type General struct {
	DataFile string `yaml:"dataFile" env:"DATA_FILE"`
}

type Edit struct {
	Default  string `yaml:"default" env:"DEFAULT"`
	Editor   string `yaml:"editor" env:"EDITOR"`
	Filetype string `yaml:"filetype" env:"FILETYPE"`
}

type Find struct {
	CaseSensitivity string `yaml:"caseSensitivity" env:"CASE_SENSITIVITY"`
}

type Theme struct {
	Primary   string `yaml:"primary" env:"PRIMARY" env-default:"Blue"`
	Secondary string `yaml:"secondary" env:"SECONDARY" env-default:"Yellow"`
	Position  string `yaml:"position" env:"POSITION" env-default:"Grey"`
	Tag       string `yaml:"tag" env:"TAG" env-default:"Magenta"`
	Title     string `yaml:"title" env:"TITLE" env-default:"BrightWhite"`
	Priority  string `yaml:"priority" env:"PRIORITY" env-default:"Yellow"`
	Status    string `yaml:"status" env:"STATUS" env-default:"White"`
}

// Config struct defines the config.yaml and related environment variables.
type Config struct {
	General General `yaml:"general" env-prefix:"GODO_GENERAL_"`
	Edit    Edit    `yaml:"edit" env-prefix:"GODO_EDIT_"`
	Find    Find    `yaml:"find" env-prefix:"GODO_FIND_"`
	Theme   Theme   `yaml:"theme" env-prefix:"GODO_THEME_"`
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
		configBoilerplate := fmt.Sprintf(`general:
  # change the file path for saving local notes (defaults to "%s/godos.json" if unset)
  # $GODO_GENERAL_DATA_FILE
  dataFile: ""

# set preferences for editing notes
edit:
  # default to either editing the note title or body (defaults to "title" if unset)
  # $GODO_EDIT_DEFAULT
  default: "body"
  # determine which editor to make edits in (defaults to the environment's $EDITOR if unset)
  # $GODO_EDIT_EDITOR
  editor: "vim"
  # append an extension to the temporary file's buffer for editing (e.g. "org", "md", "txt")
  # $GODO_EDIT_FILETYPE
  filetype: "md"

find:
  # choose between "smart", "sensitive" or "insensitive" search patterns (defaults to "smart" if unset)
  # "smart" - if the search argument is all lower-case, all results are shown. Only becoming case-sensitive
  # if upper-case characters are provided.
  # $GODO_FIND_CASE_SENSITIVITY
  caseSensitivity: "smart"

# change the colour of the output
theme:
  # use case-insesitive color names or hexadecimals, and prepend with "bg" to change the background instead.
  # $GODO_THEME_PRIMARY
  primary: "bg#00385c"
  # $GODO_THEME_SECONDARY
  secondary: "#00add8"
  # $GODO_THEME_POSITION
  position: "grey"
  # $GODO_THEME_TAG
  tag: "magenta"
  # $GODO_THEME_TITLE
  title: "brightwhite"
  # $GODO_THEME_PRIORITY
  priority: "yellow"
  # $GODO_THEME_STATUS
  status: "white"`, dataDir)

		_, err = f.WriteString(configBoilerplate)
		if err != nil {
			return fmt.Errorf("write to empty file: %w", err)
		}
	}

	return nil
}

// defineConfig assigns the file path for the configuration file, before checking
// the existence of the file and creating it if it doesn't exist.
func defineConfig() (cfgPath string, err error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Printf("determining user configuration location: %v", err)
	}

	// Assign the config.yaml filepath within the default root
	// directory godo/, to use in the user-specific configuration data.
	godoCfgDir := cfgDir + "/godo"
	cfgPath = godoCfgDir + "/config.yaml"

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

// Value takes a key and returns the matching value from the config.yaml.
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

	case "theme_status":
		value := cfg.Theme.Status

		return value

	default:
		fmt.Println("No configuration key provided")
	}

	return ""
}

// InitTheme retrieves all color values set in the config.yaml file.
func InitTheme() *Theme {
	return &Theme{
		Primary:   Value("theme_primary"),
		Secondary: Value("theme_secondary"),
		Position:  Value("theme_position"),
		Tag:       Value("theme_tag"),
		Title:     Value("theme_title"),
		Priority:  Value("theme_priority"),
		Status:    Value("theme_status"),
	}
}
