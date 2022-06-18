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
	var configPath string

	flag.StringVar(&outName, "name", "generated", "Filter file name")
	flag.StringVar(&outPath, "path", defaultPath(), "Filter location")
	flag.StringVar(&configPath, "config", "default.config.toml", "Config location")
	flag.Parse()

	configFile, err := os.Open(configPath)
	if err != nil {
		fatal(err)
	}

	config, err := filter.Load(configFile)
	if err != nil {
		fatal(err)
	}

	code := filter.Render(config)

	filterFile := filepath.Join(outPath, outName+".filter")
	if err := os.WriteFile(filterFile, []byte(code), 0o644); err != nil {
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
