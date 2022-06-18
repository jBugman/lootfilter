package filter

// spellchecker:words stretchr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClassesString(t *testing.T) {
	require.Equal(t,
		"",
		(Classes(nil)).String(),
	)

	require.Equal(t,
		`"Helmets" "Amulet"`,
		(Classes{"Helmets", "Amulet"}).String(),
	)
}

func TestColorConversion(t *testing.T) {
	require.EqualValues(t, hex("fff"), "255 255 255 255")
	require.EqualValues(t, hex("#1F54BF"), "31 84 191 255")
}
