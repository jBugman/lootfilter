//go:build tools
// +build tools

package main

// spellchecker:words mgechev, stretchr, honnef

import (
	_ "github.com/mgechev/revive"
	_ "github.com/stretchr/testify/require"
	_ "honnef.co/go/tools/cmd/staticcheck"
)
