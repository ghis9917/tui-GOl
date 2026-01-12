package main

type Symbol struct {
	Filled string
	Empty  string
}

type Symbols struct {
	CIRCLE           Symbol
	CIRCLE_EMOJI     Symbol
	SQUARE           Symbol
	SQUARE_EMOJI     Symbol
	SQUARE_BIG_EMOJI Symbol
	MOON             Symbol
}

var AvailableSymbols = Symbols{
	CIRCLE: Symbol{
		Filled: "●",
		Empty:  "○",
	},
	CIRCLE_EMOJI: Symbol{
		Filled: "⚪️",
		Empty:  "⚫️",
	},
	SQUARE: Symbol{
		Filled: "■",
		Empty:  "□",
	},
	SQUARE_EMOJI: Symbol{
		Filled: "◻️",
		Empty:  "◼️",
	},
	SQUARE_BIG_EMOJI: Symbol{
		Filled: "⬜️",
		Empty:  "⬛️",
	},
	MOON: Symbol{
		Filled: "🌕",
		Empty:  "🌑",
	},
}
