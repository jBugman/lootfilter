package filter

import (
	"strings"
)

// Filter represents high level filter config.
type Filter struct {
	Chromatic bool

	MinScrolls I

	Evasion visibility

	Chance BaseTypes

	FlasksQual I
	FlasksLvl  I
}

func (filter Filter) applyRules() []block {
	var res []block

	// 6 Sockets
	res = append(res, filter.Show(block{
		Sockets: cmp{EQ, I(6)},

		FontSize:        fontSizeDefault + 2,
		TextColor:       Color6S,
		BorderColor:     Color6S,
		BackgroundColor: ColorBG,
	}))

	// Veiled Items
	res = append(res, filter.Show(block{
		Rarity:         cmp{EQ, RarityRare},
		Identified:     TRUE,
		HasExplicitMod: `"Veil"`,
	}))

	// Chancing bases
	if filter.Chance != nil {
		res = append(res, filter.Show(block{
			BaseTypes: filter.Chance,
			Rarity:    cmp{EQ, RarityNormal},

			BorderColor: ColorChance,
		}))
	}

	// Chromatic recipe
	if filter.Chromatic {
		res = append(res, filter.Show(block{
			SocketGroup: "RGB",
			Height:      cmp{EQ, I(3)},
			Width:       cmp{EQ, I(1)},

			TextColor:   ColorChrome,
			BorderColor: ColorChrome,
		}))
		res = append(res, filter.Show(block{
			SocketGroup: "RGB",
			Height:      cmp{EQ, I(2)},
			Width:       cmp{EQ, I(2)},

			TextColor:   ColorChrome,
			BorderColor: ColorChrome,
		}))
	}

	// Leveling
	res = append(res, filter.Hide(block{
		Class:         PresetGear,
		LinkedSockets: cmp{LT, I(4)},
		DropLevel:     cmp{GTE, I(30)},
	}, cfg{
		presetGear: true,
	}))
	res = append(res, filter.Hide(block{
		Class:         PresetMelee1H.And(PresetMelee2H).And(PresetShield).And(PresetCaster).And(PresetBow),
		LinkedSockets: cmp{LT, I(3)},
		DropLevel:     cmp{LT, I(40)},
	}, cfg{
		presetGear: true,
	}))
	res = append(res, filter.Hide(block{
		Class:         PresetGear,
		Rarity:        cmp{LT, RarityRare},
		LinkedSockets: cmp{LT, I(4)},
		ItemLevel:     cmp{LT, I(45)},
	}, cfg{
		presetGear: true,
		minimize:   true,
	}))

	res = append(res, filter.Hide(block{
		Class:     PresetGear.And(PresetShield).And(PresetCaster),
		DropLevel: cmp{LT, I(58)},
	}, cfg{
		presetGear: true,
		minimize:   true,
	}))

	res = append(res, filter.Hide(block{
		Class:   PresetGear.And(PresetShield).And(PresetCaster),
		Rarity:  cmp{LT, RarityRare},
		Sockets: cmp{LT, I(6)},
	}, cfg{
		presetGear: true,
	}))
	res = append(res, filter.Hide(block{
		Class:  Classes{"Belt", "Amulet", "Ring"},
		Rarity: cmp{LT, RarityRare},
	}, cfg{
		nonInfluenced: true,
	}))

	res = append(res, filter.Hide(block{
		Class: PresetBow.And(PresetMelee1H).And(PresetMelee2H),
	}, cfg{
		presetGear: true,
		minimize:   true,
	}))

	res = append(res, filter.Hide(block{
		BaseTypes: PresetBadBelts.And(PresetBadJewelry),
	}, cfg{
		presetJewelry: true,
	}))

	res = append(res, filter.Hide(block{
		BaseTypes: BaseTypes{"Diamond Ring", "Amethyst Ring", "Unset Ring", "Moonstone Ring"},
		ItemLevel: cmp{LT, I(82)},
	}, cfg{
		presetJewelry: true,
	}))

	res = append(res, filter.Hide(block{
		Class:     PresetCaster,
		ItemLevel: cmp{LT, I(82)},
		Rarity:    cmp{LT, RarityUnique},
	}, cfg{
		presetGear: true,
	}))

	// Hide pure evasion bases
	if filter.Evasion == Hide {
		res = append(res, filter.Hide(block{
			Class:            PresetGear.And(PresetShield),
			BaseEvasion:      cmp{GT, I(0)},
			BaseArmour:       cmp{EQ, I(0)},
			BaseEnergyShield: cmp{EQ, I(0)},
		}, cfg{
			presetGear: true,
			minimize:   true,
		}))
	}

	// Abyss Jewels
	colorAbyss := hex("#36733A")
	res = append(res, filter.Show(block{
		Class: Classes{"Abyss Jewels"},

		TextColor:       colorAbyss,
		BorderColor:     colorAbyss,
		BackgroundColor: ColorBG,
	}))

	// Cluster Jewels
	colorCluster := hex("#8C8C8C")
	res = append(res, filter.Show(block{
		BaseTypes: BaseTypes{"Cluster Jewel"},

		TextColor:       colorCluster,
		BorderColor:     colorCluster,
		BackgroundColor: ColorBG,
	}))

	// Flasks
	if filter.FlasksQual > 0 {
		res = append(res, filter.Show(block{
			Class:   Classes{"Flask"},
			Quality: cmp{GTE, filter.FlasksQual},
		}))
	}
	res = append(res, filter.Hide(block{
		Class:     Classes{"Flask"},
		BaseTypes: PresetBadFlasks,
	}, cfg{
		minimize: true,
	}))
	res = append(res, filter.Show(block{ // Leveling
		Class:     Classes{"Flask"},
		AreaLevel: cmp{LT, I(50)},
	}))
	if filter.FlasksLvl > 0 {
		res = append(res, filter.Show(block{
			Class:     Classes{"Flask"},
			ItemLevel: cmp{GTE, filter.FlasksLvl},
		}))
	}
	res = append(res, filter.Hide(block{
		Class: Classes{"Flask"},
	}, cfg{
		minimize: true,
	}))

	// Gems
	res = append(res, filter.Show(block{
		Class:     Classes{"Gem"},
		BaseTypes: PresetGoodGems,

		FontSize:        fontSizeDefault + 4,
		TextColor:       ColorGem,
		BorderColor:     ColorGem,
		BackgroundColor: ColorBG,
	}))
	res = append(res, filter.Hide(block{
		Class:            Classes{"Gem"},
		AlternateQuality: FALSE,
		Quality:          cmp{LT, I(7)},
		GemLevel:         cmp{LT, I(18)},
	}, cfg{
		minimize: true,
	}))

	// Fragments
	res = append(res, filter.Show(block{
		Class: Classes{"Fragment"},

		FontSize:    fontSizeDefault + 2,
		TextColor:   ColorFragment,
		BorderColor: ColorFragment,
	}))

	// Maps
	res = append(res, filter.Show(block{
		Class: Classes{"Map"},

		FontSize:        fontSizeDefault + 2,
		TextColor:       ColorMaps,
		BorderColor:     ColorMaps,
		BackgroundColor: ColorBG,
	}))

	// Currency
	if filter.MinScrolls > 0 {
		res = append(res, filter.Hide(block{
			BaseTypes: BaseTypes{"Scroll of Wisdom", "Portal Scroll"},
			AreaLevel: cmp{GT, I(55)},
			StackSize: cmp{LT, filter.MinScrolls},
		}, nil))
	}
	res = append(res, filter.Hide(block{
		BaseTypes: BaseTypes{"Engineer's Shard", "Transmutation Shard"},
	}, cfg{
		minimize: true,
	}))
	res = append(res, filter.Show(block{
		Class: Classes{"Currency"},

		FontSize:    fontSizeDefault + 3,
		TextColor:   ColorCurrency,
		BorderColor: ColorCurrency,
	}))

	// Sentinels
	colorSentinel := hex("#1F54BF")
	res = append(res, filter.Hide(block{
		Class:     Classes{"Sentinel"},
		BaseTypes: BaseTypes{"Emberstone", "Brimstone"},
	}, cfg{
		minimize: true,
	}))
	res = append(res, filter.Show(block{
		Class: Classes{"Sentinel"},

		TextColor:       colorSentinel,
		BorderColor:     colorSentinel,
		BackgroundColor: ColorBG,
	}))

	return res
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

func (Filter) Show(b block) block {
	b.Visibility = Show
	if b.FontSize == 0 {
		b.FontSize = fontSizeDefault
	}
	return b
}

func (Filter) Hide(b block, cfg hideConfig) block {
	b.Visibility = Hide
	b.FontSize = fontSizeHidden
	b.DisableDropSound = true

	if cfg == nil {
		return b
	}

	if cfg.PresetGear() {
		b.Sockets = cmp{LT, I(6)}
	}
	presetItem := cfg.PresetGear() || cfg.PresetJewelry()
	if presetItem {
		if b.Rarity == nil {
			b.Rarity = cmp{LT, RarityUnique}
		}
	}

	if cfg.NonInfluenced() || presetItem {
		b.Elder = FALSE
		b.Shaper = FALSE
		b.Influence = InfluenceNone
	}
	if cfg.Minimize() {
		b.FontSize = fontSizeMin
		b.BackgroundColor = ColorZero
		b.BorderColor = ColorZero
		b.DisableDropSound = true
	}

	return b
}

type hideConfig interface {
	NonInfluenced() bool

	Minimize() bool

	PresetGear() bool
	PresetJewelry() bool
}

type cfg struct {
	nonInfluenced bool
	minimize      bool

	presetGear    bool
	presetJewelry bool
}

func (c cfg) NonInfluenced() bool {
	return c.nonInfluenced
}

func (c cfg) Minimize() bool {
	return c.minimize
}

func (c cfg) PresetGear() bool {
	return c.presetGear
}

func (c cfg) PresetJewelry() bool {
	return c.presetJewelry
}
