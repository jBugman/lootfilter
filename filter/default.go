package filter

// spellchecker:words Ezomyte

var defaultFilter = Filter{
	Chromatic: true,

	Evasion: Hide,

	MinScrolls: 5,

	FlasksQual: 10,
	FlasksLvl:  82,

	MinGearDropLvl: 70, // show all atlas bases

	GoodBases: BaseTypes{
		"Pig-Faced Bascinet",
		"Full Dragonscale",
		"Dragonscale Gauntlets",
		"Cardinal Round Shield",
		"Hubris Circlet",
	},

	Chance: BaseTypes{
		"Champion Kite Shield",
		"Ezomyte Dagger",
		"Fiend Dagger",
		"Prophecy Wand",
	},

	BadCards: BaseTypes{
		"The Lover",
		"The Carrion Crow",
	},
	GoodCards: BaseTypes{
		"The Apothecary",
	},
}

// Default is default SSF filter config.
var Default = defaultFilter
