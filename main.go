package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"lootfilter/filter"
)

func main() {
	var outName string
	var outPath string

	flag.StringVar(&outName, "name", "generated", "Filter file name")
	flag.StringVar(&outPath, "path", defaultPath(), "Filter location")
	flag.Parse()

	filterFile := filepath.Join(outPath, outName+".filter")

	custom := filter.Default

	if err := os.WriteFile(filterFile, []byte(filter.Render(custom)), 0o644); err != nil {
		fatal(err)
	}
}

func defaultPath() string {
	osUser, err := user.Current()
	if err != nil {
		fatal(err)
	}

	return filepath.Join(osUser.HomeDir, "Documents", "My Games", "Path of Exile")
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
