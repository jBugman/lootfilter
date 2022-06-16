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
					Height:      CmpLTE("3"),
					Width:       CmpEQ("1"),

					TextColor:   ColorChrome,
					BorderColor: ColorChrome,
				},
			},
			section{
				block: block{
					Visibility:  Show,
					SocketGroup: "RGB",
					Height:      CmpEQ("2"),
					Width:       CmpEQ("2"),

					TextColor:   ColorChrome,
					BorderColor: ColorChrome,
				},
			},
		)
	}

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class: append(PresetGear, PresetShield...),

				Rarity: CmpLT(string(RarityUnique)),

				Sockets: CmpLT("6"),

				DropLevel: CmpLT("58"),
			},

			Hide: HideFully,

			NonInfluenced: true,
		},
	)

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class: PresetGear,

				Rarity: CmpLT(string(RarityRare)),

				LinkedSockets: CmpLT("4"),
				Sockets:       CmpLT("6"),

				ItemLevel: CmpGT("45"),
			},

			Hide: HideFully,

			NonInfluenced: true,
		},
	)

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class: PresetCaster,

				Rarity: CmpLT(string(RarityUnique)),

				Sockets: CmpLT("6"),

				DropLevel: CmpLT("58"),
			},

			Hide: HideFully,

			NonInfluenced: true,
		},
	)

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class: append(PresetBow, append(PresetMelee1H, PresetMelee2H...)...),

				Rarity: CmpLT(string(RarityUnique)),

				Sockets: CmpLT("6"),
			},

			Hide: HideFully,

			NonInfluenced: true,
		},
	)

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				BaseTypes: append(PresetBadBelts, PresetBadJewelry...),

				Rarity: CmpLT(string(RarityUnique)),
			},

			Hide: HideClickable,

			NonInfluenced: true,
		},
	)

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class: Classes{"Shield", "Wand"},

				ItemLevel: CmpLT("82"),

				Rarity: CmpLT(string(RarityUnique)),
			},

			Hide: HideClickable,

			NonInfluenced: true,
		},
	)

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				BaseTypes: BaseTypes{"Diamond Ring", "Amethyst Ring", "Unset Ring", "Moonstone Ring"},

				ItemLevel: CmpLT("82"),

				Rarity: CmpLT(string(RarityUnique)),
			},

			Hide: HideClickable,

			NonInfluenced: true,
		},
	)

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class: PresetCaster,

				ItemLevel: CmpLT("82"),

				Rarity: CmpLT(string(RarityUnique)),
			},

			Hide: HideClickable,

			NonInfluenced: true,
		},
	)

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
			Visibility: Show,
			Sockets:    CmpEQ("6"),

			TextColor:       Color6S,
			BorderColor:     Color6S,
			BackgroundColor: ColorBG,

			FontSize: defaltFontSize + 3,
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
			Visibility: Hide,
			Class:      Classes{"Gem"},

			AlternateQuality: FALSE,

			Quality:  CmpLT("7"),
			GemLevel: CmpLT("18"),

			Continue: true,
		},
		Hide: HideFully,
	})

	ss = append(ss, section{
		block: block{
			Visibility: Show,
			Class:      Classes{"Gem"},
			BaseTypes:  PresetGoodGems,

			FontSize: defaltFontSize + 4,

			TextColor:       ColorGem,
			BorderColor:     ColorGem,
			BackgroundColor: ColorBG,
		},
	})

	ss = append(ss, section{
		block: block{
			Visibility: Show,
			Class:      Classes{"Fragment"},

			FontSize: defaltFontSize + 3,

			TextColor:   ColorFragment,
			BorderColor: ColorFragment,
		},
	})

	ss = append(ss, section{
		block: block{
			Visibility: Show,
			Class:      Classes{"Map"},

			FontSize: defaltFontSize + 1,

			TextColor:       ColorMaps,
			BorderColor:     ColorMaps,
			BackgroundColor: ColorBG,
		},
	})

	ss = append(ss, section{
		block: block{
			Visibility: Show,
			Class:      Classes{"Currency"},

			FontSize: defaltFontSize + 3,

			TextColor:   ColorCurrency,
			BorderColor: ColorCurrency,
		},
	})

	ss = append(ss, section{
		block: block{
			Visibility: Hide,
			BaseTypes:  BaseTypes{"Scroll of Wisdom", "Portal Scroll"},
			AreaLevel:  CmpGT("55"),

			StackSize: CmpLT("5"),
		},
	})

	ss = append(ss, section{
		block: block{
			Visibility: Hide,
			BaseTypes:  BaseTypes{"Engineer's Shard", "Transmutation Shard"},
		},
		Hide: HideFully,
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
