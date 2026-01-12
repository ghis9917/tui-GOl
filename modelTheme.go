package main

type AvailableTheme struct {
	LIGHT Theme
	DARK  Theme
}

type Theme struct {
	Primary    Color
	Secondary  Color
	Selection  Color
	Background Color
}

var Themes = AvailableTheme{
	LIGHT: Theme{
		Primary:    Colors.FG_RED,
		Secondary:  Colors.FG_GREY,
		Selection:  Colors.BG_GREY,
		Background: Colors.BG_WHITE,
	},
	DARK: Theme{
		Primary:    Colors.FG_MAGENTA,
		Secondary:  Colors.FG_CYAN,
		Selection:  Colors.BG_RED,
		Background: Colors.BG_GREY,
	},
}
