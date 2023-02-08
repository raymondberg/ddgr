package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/raymondberg/ddgr/pkg"
)

func Run() {
	const CONFIG_FILENAME string = ".dogrc"

	currentUser, _ := user.Current()
	config, err := pkg.ConfigFromFile(filepath.Join(currentUser.HomeDir, CONFIG_FILENAME))

	if err != nil {
		os.Exit(1)
	}

	var idSearchString string
	flag.StringVar(&idSearchString, "s", "", "search for a specific id")
	flag.Parse()

	if idSearchString == "" {
		query := strings.Join(os.Args[1:], " ")
		fmt.Printf("Searching for '%s'\n", query)

		result := pkg.Search(query, config)
		pkg.DumpTable(result)
	} else {
		fmt.Print("Searching by id not yet supported")
	}
}
