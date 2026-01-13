package main

type Cell struct {
	status   bool
	selected bool
}

func (c *Cell) String() string {
	representation := ""

	if c.status {
		if c.selected {
			representation = ColoredString(ALIVE_STRING, THEME.Primary, THEME.Selection)
		} else {
			representation = ColoredString(ALIVE_STRING, THEME.Primary, THEME.AliveCell)
		}
	} else {
		if c.selected {
			representation = ColoredString(DEAD_STRING, THEME.Secondary, THEME.Selection)
		} else {
			representation = ColoredString(DEAD_STRING, THEME.Secondary, THEME.DeadCell)
		}
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
