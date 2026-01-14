package main

type Theme struct {
	Name       string
	Primary    Color
	Secondary  Color
	AliveCell  Color
	DeadCell   Color
	Selection  Color
	Background Color
}

type AvailableTheme struct {
	LIGHT          Theme
	DARK           Theme
	LIGHT_RED_BLUE Theme
	DARK_RED_BLUE  Theme
	DARK_MAGENTA   Theme
}

var Themes = AvailableTheme{
	LIGHT: Theme{
		Name:       "LIGHT",
		Primary:    Colors.FG_DARK_GREY,
		Secondary:  Colors.FG_BLACK,
		AliveCell:  Colors.BG_DARK_GREY,
		DeadCell:   Colors.BG_GREY,
		Selection:  Colors.BG_RED,
		Background: Colors.BG_WHITE,
	},
	DARK: Theme{
		Name:       "DARK",
		Primary:    Colors.FG_WHITE,
		Secondary:  Colors.FG_BLUE,
		AliveCell:  Colors.BG_WHITE,
		DeadCell:   Colors.BG_DARK_GREY,
		Selection:  Colors.BG_BLUE,
		Background: Colors.BG_BLACK,
	},
	LIGHT_RED_BLUE: Theme{
		Name:       "Light Red & Blue",
		Primary:    Colors.FG_RED,
		Secondary:  Colors.FG_BLUE,
		AliveCell:  Colors.BG_RED,
		DeadCell:   Colors.BG_BLUE,
		Selection:  Colors.BG_BLACK,
		Background: Colors.BG_WHITE,
	},
	DARK_RED_BLUE: Theme{
		Name:       "Dark Red & Blue",
		Primary:    Colors.FG_BLUE,
		Secondary:  Colors.FG_RED,
		AliveCell:  Colors.BG_BLUE,
		DeadCell:   Colors.BG_RED,
		Selection:  Colors.BG_WHITE,
		Background: Colors.BG_BLACK,
	},
	DARK_MAGENTA: Theme{
		Name:       "Dark Magenta",
		Primary:    Colors.FG_BRIGHT_MAGENTA,
		Secondary:  Colors.FG_WHITE,
		AliveCell:  Colors.BG_BRIGHT_MAGENTA,
		DeadCell:   Colors.BG_WHITE,
		Selection:  Colors.BG_DARK_GREY,
		Background: Colors.BG_BLACK,
	},
}
