package filter

const minFontSize = 18
const clickableFontSize = 26
const defaultFontSize = 33

type visibility string

const (
	// Hide means block is about hiding loot.
	Hide visibility = "Hide"
	// Show means block is about displaying loot.
	Show visibility = "Show"
)

type influence string

const (
	// InfluenceNone is for no-influence items.
	InfluenceNone influence = "None"
)

type color string

const (
	// ColorZero is for invisibility.
	ColorZero color = "0 0 0 0"

	// ColorCurrency is PoE native color for currency.
	ColorCurrency color = "170 158 130 255"
	// ColorGem is PoE native color for currency.
	ColorGem color = "27 162 155 255"

	// ColorFragment is color for fragments.
	ColorFragment color = "140 0 0 255"

	// ColorMaps is nice color for maps inspired by NeverSink.
	ColorMaps color = "230 204 128 255"

	// ColorChrome is NeverSink's default color for the chromatic recipe.
	ColorChrome color = "235 235 235 255"
	// ColorChance is NeverSink's default color for chancing bases.
	ColorChance color = "33 103 33 255"
	// Color6S is  NeverSink's default color for 6S bases.
	Color6S color = "200 200 200 255"

	// ColorBG is default background color.
	ColorBG color = "51 51 51 255" // #333
)

type rarity string

const (
	// RarityNormal is for white items.
	RarityNormal rarity = "Normal"
	// RarityMagic is for magic items.
	RarityMagic rarity = "Magic"
	// RarityRare is for rare items.
	RarityRare rarity = "Rare"
	// RarityUnique is for unique items.
	RarityUnique rarity = "Unique"
)
