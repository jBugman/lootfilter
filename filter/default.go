package filter

// spellchecker:words Ezomyte

const defaltFontSize = 34

var defaultFilter = Filter{
	sections: []section{
		{
			block: block{
				Visibility: Show,
				Continue:   true,
				FontSize:   defaltFontSize,
			},
		},
	},

	Chromatic: true,

	Evasion: Hide,

	Chance: BaseTypes{"Champion Kite Shield", "Ezomyte Dagger", "Fiend Dagger", "Prophecy Wand"},

	LevelingHideNon4L: true,
	LevelingHideNon3L: true,
}

// Default is default SSF filter config.
var Default = defaultFilter
