package filter

// spellchecker:words Ezomyte

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

	Chance: BaseTypes{"Champion Kite Shield", "Ezomyte Dagger", "Fiend Dagger", "Prophecy Wand"},
}

var Default = defaultFilter
