package main

type AvailableTheme struct {
	LIGHT               Theme
	LIGHT_HIGH_CONTRAST Theme
	DARK                Theme
	DARK_HIGH_CONTRAST  Theme
}

type Theme struct {
	Primary    Color
	Secondary  Color
	AliveCell  Color
	DeadCell   Color
	Selection  Color
	Background Color
}

var Themes = AvailableTheme{
	LIGHT: Theme{
		Primary:    Colors.FG_RED,
		Secondary:  Colors.FG_BLUE,
		AliveCell:  Colors.BG_WHITE,
		DeadCell:   Colors.BG_WHITE,
		Selection:  Colors.BG_DARK_GREY,
		Background: Colors.BG_WHITE,
	},
	LIGHT_HIGH_CONTRAST: Theme{
		Primary:    Colors.FG_RED,
		Secondary:  Colors.FG_BLUE,
		AliveCell:  Colors.BG_RED,
		DeadCell:   Colors.BG_BLUE,
		Selection:  Colors.BG_DARK_GREY,
		Background: Colors.BG_WHITE,
	},
	DARK: Theme{
		Primary:    Colors.FG_WHITE,
		Secondary:  Colors.FG_BRIGHT_CYAN,
		AliveCell:  Colors.BG_GREY,
		DeadCell:   Colors.BG_DARK_GREY,
		Selection:  Colors.BG_MAGENTA,
		Background: Colors.BG_BLACK,
	},
	DARK_HIGH_CONTRAST: Theme{
		Primary:    Colors.FG_WHITE,
		Secondary:  Colors.FG_MAGENTA,
		AliveCell:  Colors.BG_WHITE,
		DeadCell:   Colors.BG_BLACK,
		Selection:  Colors.BG_MAGENTA,
		Background: Colors.BG_BLACK,
	},
}
