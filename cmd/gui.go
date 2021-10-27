/*
Copyright © 2021 Richard Halford <richard@xhalford.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"net/http"

	"github.com/rsHalford/godo/config"
	"github.com/rsHalford/godo/gui"
	"github.com/spf13/cobra"
)

// guiCmd represents the gui command.
var guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "view GoDo using your browser as a gui",
	Long:  `Gui lets you run GoDo as a GUI application in your browser, locally.`,
	Run:   guiRun,
}

// guiHandler gets the GUI filesystem's public directory to be served.
func guiHandler() http.Handler {
	fsys := fs.FS(gui.Gui)
	contentStatic, _ := fs.Sub(fsys, "public")

	return http.FileServer(http.FS(contentStatic))
}

// handleRequests serves the GUI handler on the given port.
func handleRequests(port string) {
	mux := http.NewServeMux()
	mux.Handle("/", guiHandler())
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func guiRun(cmd *cobra.Command, args []string) {
	// Port used is set to 5000, unless assigned in the configuration file.
	port := "5000"
	if config.Value("gui_port") != "" {
		port = config.Value("gui_port")
	}

	fmt.Printf("\033[34m::\033[0m GoDo GUI is currently running on \033[33m\033[4mhttp://localhost:%s\033[0m\n", port)

	handleRequests(port)
}

func init() {
	rootCmd.AddCommand(guiCmd)
}
