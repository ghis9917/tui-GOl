package main

type Cell struct {
	status   bool
	selected bool
}

func (c *Cell) String() string {
	representation := ""

	var bg Color
	if c.selected {
		bg = THEME.Selection //Change to keep contrast with emoji/unicode color
	} else {
		bg = THEME.Background
	}

	if c.status {
		representation = ColoredString(ALIVE_STRING, THEME.Primary, bg) //Applied only to unicode symbols
	} else {
		representation = ColoredString(DEAD_STRING, THEME.Secondary, bg) //Applied only to unicode symbols
	}

	return representation

}

func (c *Cell) Fate(neighbours int) {
	if c.status == ALIVE {

		if neighbours < 2 || neighbours >= 4 {
			c.status = DEAD
		}

	} else {

		if neighbours == 3 {
			c.status = ALIVE
		}

	}
}
