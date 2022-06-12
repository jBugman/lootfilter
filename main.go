package main

import (
	"fmt"
	"os"
	"os/user"
	"path"

	"lootfilter/filter"
)

func main() {
	osUser, err := user.Current()
	if err != nil {
		fatal(err)
	}

	home := osUser.HomeDir
	poeFolder := path.Join(home, "Documents", "My Games", "Path of Exile")

	filterFile := path.Join(poeFolder, "generated.filter")

	custom := filter.Default

	if err := os.WriteFile(filterFile, []byte(custom.String()), 0o644); err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
