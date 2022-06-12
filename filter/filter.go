package filter

import (
	"strings"
)

var _t = true
var _f = false
var TRUE = &_t
var FALSE = &_f

const defaltFontSize = 34

type Filter struct {
	sections []section

	Chromatic bool

	Evasion visibility
}

func (f Filter) virtualSections() []section {
	ss := f.sections

	if f.Chromatic {
		ss = append(ss,
			section{
				block: block{
					Visibility:  Show,
					SocketGroup: "RGB",
					Height:      3,
					Width:       1,
					BorderColor: "235 235 235 255",
				},
			},
			section{
				block: block{
					Visibility:  Show,
					SocketGroup: "RGB",
					Height:      2,
					Width:       2,
					BorderColor: "235 235 235 255",
				},
			},
		)
	}

	if f.Evasion == Hide {
		ss = append(ss,
			section{
				block: block{
					Visibility: Hide,

					BaseEvasion:      cmpGT("0"),
					BaseArmour:       cmpEQ("0"),
					BaseEnergyShield: cmpEQ("0"),
				},
				Gear:          true,
				NonInfluenced: true,
				Hide:          HideFully,
			},
		)
	}

	return ss
}

func (f Filter) String() string {
	var buf strings.Builder
	for _, sec := range f.virtualSections() {
		_, _ = buf.WriteString(sec.String())
		_, _ = buf.WriteString("\n")
	}
	res := strings.TrimSpace(buf.String())
	res += "\n"
	return res
}

var defaultFilter = Filter{
	sections: []section{
		{
			block: block{
				Visibility: Show,
				Continue:   true,
				FontSize:   34,
			},
		},
	},

	Chromatic: true,

	Evasion: Hide,
}

var Default = defaultFilter

type section struct {
	block

	Gear bool

	NonInfluenced bool

	Hide hideLevel
}

type hideLevel string

const (
	HideFully hideLevel = "fully"
)

func (sec section) String() string {
	if sec.Gear {
		sec.Class = classes{"Helmet", "Body armour", "Gloves", "Boots"}
	}

	if sec.NonInfluenced {
		sec.ElderItem = FALSE
		sec.ShaperItem = FALSE
		sec.HasInfluence = InfluenceNone
	}

	if sec.Hide == "fully" {
		sec.FontSize = 18
		sec.BackgroundColor = ColorZero
		sec.BorderColor = ColorZero
		sec.DisableDropSound = true
	}

	return sec.block.String()
}

func init() {
	initBlockTemplate()
}
