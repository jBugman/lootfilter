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
	Chromatic bool

	MinScrolls I

	Evasion visibility

	Chance BaseTypes
}

func (filter Filter) applyRules() []section {
	var ss []section

	// 6 Sockets
	ss = append(ss, section{
		block: block{
			Visibility: Show,

			Sockets: cmp{EQ, I(6)},

			FontSize:        defaultFontSize + 2,
			TextColor:       Color6S,
			BorderColor:     Color6S,
			BackgroundColor: ColorBG,
		},
	})

	// Veiled Items
	ss = append(ss, section{
		block: block{
			Visibility: Show,

			Rarity:         cmp{EQ, RarityRare},
			Identified:     TRUE,
			HasExplicitMod: `"Veil"`,

			FontSize: defaultFontSize,
		},
	})

	// Chancing bases
	if filter.Chance != nil {
		ss = append(ss,
			section{
				block: block{
					Visibility: Show,

					BaseTypes: filter.Chance,
					Rarity:    cmp{EQ, RarityNormal},

					BorderColor: ColorChance,
				},
			},
		)
	}

	// Chromatic recipe
	if filter.Chromatic {
		ss = append(ss,
			section{
				block: block{
					Visibility:  Show,
					SocketGroup: "RGB",
					Height:      cmp{EQ, I(3)},
					Width:       cmp{EQ, I(1)},

					FontSize:    defaultFontSize,
					TextColor:   ColorChrome,
					BorderColor: ColorChrome,
				},
			},
			section{
				block: block{
					Visibility:  Show,
					SocketGroup: "RGB",
					Height:      cmp{EQ, I(2)},
					Width:       cmp{EQ, I(2)},

					FontSize:    defaultFontSize,
					TextColor:   ColorChrome,
					BorderColor: ColorChrome,
				},
			},
		)
	}

	// Leveling
	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class:         PresetGear,
				Rarity:        cmp{LT, RarityUnique},
				LinkedSockets: cmp{LT, I(4)},
				DropLevel:     cmp{LT, I(40)},
			},
			Hide:          HideClickable,
			NonInfluenced: true,
		},
	)
	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class:         PresetMelee1H.And(PresetMelee2H).And(PresetShield).And(PresetCaster).And(PresetBow),
				Rarity:        cmp{LT, RarityUnique},
				LinkedSockets: cmp{LT, I(3)},
				DropLevel:     cmp{LT, I(40)},
			},
			Hide:          HideClickable,
			NonInfluenced: true,
		},
	)
	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class:         PresetGear,
				Rarity:        cmp{LT, RarityRare},
				LinkedSockets: cmp{LT, I(4)},
				Sockets:       cmp{LT, I(6)},
				ItemLevel:     cmp{LT, I(45)},
			},

			Hide: HideFully,

			NonInfluenced: true,
		},
	)

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class:     PresetGear.And(PresetShield).And(PresetCaster),
				Rarity:    cmp{LT, RarityUnique},
				Sockets:   cmp{LT, I(6)},
				DropLevel: cmp{LT, I(58)},
			},

			Hide: HideFully,

			NonInfluenced: true,
		},
	)

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class:   PresetGear.And(PresetShield).And(PresetCaster),
				Rarity:  cmp{LT, RarityRare},
				Sockets: cmp{LT, I(6)},
			},

			Hide: HideClickable,

			NonInfluenced: true,
		},
	)
	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class:  Classes{"Belt", "Amulet", "Ring"},
				Rarity: cmp{LT, RarityRare},
			},

			Hide: HideClickable,

			NonInfluenced: true,
		},
	)

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class:   PresetBow.And(PresetMelee1H).And(PresetMelee2H),
				Rarity:  cmp{LT, RarityUnique},
				Sockets: cmp{LT, I(6)},
			},

			Hide: HideFully,

			NonInfluenced: true,
		},
	)

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				BaseTypes: PresetBadBelts.And(PresetBadJewelry),
				Rarity:    cmp{LT, RarityUnique},
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
				ItemLevel: cmp{LT, I(82)},
				Rarity:    cmp{LT, RarityUnique},
			},

			Hide: HideClickable,

			NonInfluenced: true,
		},
	)

	ss = append(ss,
		section{
			block: block{
				Visibility: Hide,

				Class:     PresetCaster,
				ItemLevel: cmp{LT, I(82)},
				Rarity:    cmp{LT, RarityUnique},
			},

			Hide: HideClickable,

			NonInfluenced: true,
		},
	)

	// Hide pure evasion bases
	if filter.Evasion == Hide {
		ss = append(ss,
			section{
				block: block{
					Visibility: Hide,

					Class:   PresetGear.And(PresetShield),
					Rarity:  cmp{LT, RarityUnique},
					Sockets: cmp{LT, I(6)},

					BaseEvasion:      cmp{GT, I(0)},
					BaseArmour:       cmp{EQ, I(0)},
					BaseEnergyShield: cmp{EQ, I(0)},
				},

				Hide: HideFully,

				NonInfluenced: true,
			},
		)
	}

	// Flasks
	ss = append(ss, section{
		block: block{
			Visibility: Hide,

			Class:     Classes{"Flask"},
			AreaLevel: cmp{GTE, I(50)},
			ItemLevel: cmp{LTE, I(74)},
			Quality:   cmp{LT, I(7)},
		},
		Hide: HideFully,
	})

	// Gems
	ss = append(ss, section{
		block: block{
			Visibility: Show,

			Class:     Classes{"Gem"},
			BaseTypes: PresetGoodGems,

			FontSize:        defaultFontSize + 4,
			TextColor:       ColorGem,
			BorderColor:     ColorGem,
			BackgroundColor: ColorBG,
		},
	})
	ss = append(ss, section{
		block: block{
			Visibility: Hide,

			Class:            Classes{"Gem"},
			AlternateQuality: FALSE,
			Quality:          cmp{LT, I(7)},
			GemLevel:         cmp{LT, I(18)},
		},
		Hide: HideFully,
	})

	// Fragments
	ss = append(ss, section{
		block: block{
			Visibility: Show,

			Class: Classes{"Fragment"},

			FontSize:    defaultFontSize + 3,
			TextColor:   ColorFragment,
			BorderColor: ColorFragment,
		},
	})

	// Maps
	ss = append(ss, section{
		block: block{
			Visibility: Show,

			Class: Classes{"Map"},

			FontSize:        defaultFontSize + 1,
			TextColor:       ColorMaps,
			BorderColor:     ColorMaps,
			BackgroundColor: ColorBG,
		},
	})

	// Currency
	if filter.MinScrolls > 0 {
		ss = append(ss, section{
			block: block{
				Visibility: Hide,

				BaseTypes: BaseTypes{"Scroll of Wisdom", "Portal Scroll"},
				AreaLevel: cmp{GT, I(55)},
				StackSize: cmp{LT, filter.MinScrolls},
			},
		})
	}
	ss = append(ss, section{
		block: block{
			Visibility: Hide,

			BaseTypes: BaseTypes{"Engineer's Shard", "Transmutation Shard"},
		},
		Hide: HideFully,
	})
	ss = append(ss, section{
		block: block{
			Visibility: Show,

			Class: Classes{"Currency"},

			FontSize:    defaultFontSize + 3,
			TextColor:   ColorCurrency,
			BorderColor: ColorCurrency,
		},
	})

	return ss
}

func Render(f Filter) string {
	var buf strings.Builder
	for _, block := range f.applyRules() {
		_, _ = buf.WriteString(block.String())
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
		sec.Elder = FALSE
		sec.Shaper = FALSE
		sec.Influence = InfluenceNone
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
