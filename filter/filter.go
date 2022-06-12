package filter

import (
	"strings"
)

var _t = true
var _f = false

var (
	// TRUE is boolean true.
	TRUE = &_t
	// FALSE is boolean false.
	FALSE = &_f
)

// Filter represents high level filter config.
type Filter struct {
	sections []section

	Chromatic bool

	Evasion visibility

	Chance BaseTypes

	Leveling4L bool
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
					BorderColor: ColorChrome,
				},
			},
			section{
				block: block{
					Visibility:  Show,
					SocketGroup: "RGB",
					Height:      2,
					Width:       2,
					BorderColor: ColorChrome,
				},
			},
		)
	}

	if f.Evasion == Hide {
		ss = append(ss,
			section{
				block: block{
					Visibility: Hide,

					BaseEvasion:      CmpGT("0"),
					BaseArmour:       CmpEQ("0"),
					BaseEnergyShield: CmpEQ("0"),
				},
				Gear:          true,
				NonInfluenced: true,
				Hide:          HideFully,
			},
		)
	}

	if f.Leveling4L {
		ss = append(ss,
			section{
				block: block{
					Visibility:    Hide,
					Rarity:        CmpLT(string(RarityUnique)),
					LinkedSockets: CmpLT("4"),
				},
				Hide:          HideClickable,
				Gear:          true,
				NonInfluenced: true,
			},
		)
	}

	if f.Chance != nil {
		ss = append(ss,
			section{
				block: block{
					Visibility:  Show,
					BaseTypes:   f.Chance,
					Rarity:      CmpEQ(string(RarityNormal)),
					BorderColor: ColorChance,
				},
				NonInfluenced: true,
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

type section struct {
	block

	OneHanded bool
	TwoHanded bool

	Gear   bool
	Shield bool

	NonInfluenced bool

	Hide hideLevel
}

type hideLevel string

const (
	// HideFully hides item, minimizes font and disables other decorations and sounds.
	HideFully hideLevel = "fully"
	// HideClickable hides item, makes font smaller and disables other decorations and sounds.
	HideClickable hideLevel = "clickable"
)

func (sec section) String() string {
	if sec.Gear {
		sec.Class = Classes{"Helmet", "Body armour", "Gloves", "Boots"}
	}

	if sec.NonInfluenced {
		sec.ElderItem = FALSE
		sec.ShaperItem = FALSE
		sec.HasInfluence = InfluenceNone
	}

	switch sec.Hide {
	case HideFully:
		sec.FontSize = minFontSize
		sec.BackgroundColor = ColorZero
		sec.BorderColor = ColorZero
		sec.DisableDropSound = true

	case HideClickable:
		sec.FontSize = clickableFontSize
		sec.BackgroundColor = ColorZero
		sec.BorderColor = ColorZero
		sec.DisableDropSound = true
	}

	return sec.block.String()
}

func init() {
	initBlockTemplate()
}
