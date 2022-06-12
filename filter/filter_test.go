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
