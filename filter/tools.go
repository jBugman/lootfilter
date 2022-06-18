package filter

import (
	"fmt"
	"strconv"
	"strings"
)

func hex(hex string) color {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 3 && len(hex) != 6 {
		panic("not a hex color")
	}
	bs := []rune(hex)
	if len(bs) == 3 {
		bs = []rune{bs[0], bs[0], bs[1], bs[1], bs[2], bs[2]}
	}
	r, err := strconv.ParseUint(string(bs[0:0+2]), 16, 8)
	if err != nil {
		panic(err)
	}
	g, err := strconv.ParseUint(string(bs[2:2+2]), 16, 8)
	if err != nil {
		panic(err)
	}
	b, err := strconv.ParseUint(string(bs[4:4+2]), 16, 8)
	if err != nil {
		panic(err)
	}
	return color(fmt.Sprintf("%d %d %d 255", r, g, b))
}
