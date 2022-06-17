package filter

// spellchecker:words Ezomyte

var defaultFilter = Filter{
	Chromatic: true,

	Evasion: Hide,

	MinScrolls: 5,

	Chance: BaseTypes{"Champion Kite Shield", "Ezomyte Dagger", "Fiend Dagger", "Prophecy Wand"},
}

// Default is default SSF filter config.
var Default = defaultFilter
