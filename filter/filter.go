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

	LevelingHideNon4L bool
	LevelingHideNon3L bool
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

					Class: append(PresetGear, PresetShield...),

					Rarity: CmpLT(string(RarityUnique)),

					Sockets: CmpLT("6"),

					BaseEvasion:      CmpGT("0"),
					BaseArmour:       CmpEQ("0"),
					BaseEnergyShield: CmpEQ("0"),
				},

				Hide: HideFully,

				NonInfluenced: true,
			},
		)
	}

	if f.LevelingHideNon4L {
		ss = append(ss,
			section{
				block: block{
					Visibility: Hide,

					Class: PresetGear,

					Rarity: CmpLT(string(RarityUnique)),

					LinkedSockets: CmpLT("4"),

					DropLevel: CmpLT("40"),
				},
				Hide:          HideClickable,
				NonInfluenced: true,
			},
		)
	}

	if f.LevelingHideNon3L {
		ss = append(ss,
			section{
				block: block{
					Visibility: Hide,

					Class: append(
						PresetMelee1H, append(
							PresetMelee2H, append(
								PresetShield, append(
									PresetCaster,
									PresetBow...)...)...)...),

					Rarity: CmpLT(string(RarityUnique)),

					LinkedSockets: CmpLT("3"),

					DropLevel: CmpLT("40"),
				},
				Hide:          HideClickable,
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

	ss = append(ss, section{
		block: block{
			Visibility:  Show,
			Sockets:     CmpEQ("6"),
			BorderColor: Color6S,
		},
	})

	ss = append(ss, section{
		block: block{
			Visibility: Hide,
			Class:      Classes{"Flask"},

			AreaLevel: CmpGTE("50"),
			ItemLevel: CmpLTE("74"),

			Quality: CmpLT("7"),
		},
		Hide: HideFully,
	})

	ss = append(ss, section{
		block: block{
			Visibility: Show,
			Class:      Classes{"Currency"},

			FontSize:    defaltFontSize + 2,
			TextColor:   ColorCurrency,
			BorderColor: ColorCurrency,
		},
	})

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
