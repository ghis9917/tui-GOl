package main

import (
	"fmt"
	"math/rand"

	"github.com/eiannone/keyboard"
)

type Board struct {
	cells        [][]Cell
	height       int
	width        int
	stats        Stats
	started      bool
	selectedCell Coordinate
	cfg          Config
}

type Stats struct {
	generation       int
	population       []float64
	populationChange []float64
}

func NewBoard(cfg Config) (board *Board) {

	b := &Board{
		cells:        make([][]Cell, cfg.height),
		height:       cfg.height,
		width:        cfg.width,
		started:      false,
		selectedCell: Coordinate{i: 0, j: 0},
		cfg:          cfg,
		stats: Stats{
			generation:       0,
			population:       []float64{},
			populationChange: []float64{},
		},
	}
	limit := int(float64(cfg.width*cfg.height) * cfg.fill)

	for i := range b.cells {
		b.cells[i] = make([]Cell, cfg.width)
		for j := range b.cells[i] {
			if rand.Intn(int(1.0/cfg.fill)) == 0 && limit > 0 {
				b.cells[i][j] = Cell{status: ALIVE, selected: Coordinate{i: i, j: j} == b.selectedCell}
				limit -= 1
			} else {
				b.cells[i][j] = Cell{status: DEAD, selected: Coordinate{i: i, j: j} == b.selectedCell}
			}
		}
	}

	b.CollectPopulationStat() // Initialize stats on population before the first step is run

	return b

}

func (b *Board) SwitchCellStatus() {

	b.cells[b.selectedCell.i][b.selectedCell.j].status = !b.cells[b.selectedCell.i][b.selectedCell.j].status

}

func (b *Board) SwitchCellSelection() {

	b.cells[b.selectedCell.i][b.selectedCell.j].selected = !b.cells[b.selectedCell.i][b.selectedCell.j].selected

}

func (b *Board) CleanSelection() {

	for i := range b.cells {
		for j := range b.cells[i] {

			b.cells[i][j].selected = false

		}
	}

}

func (b *Board) Evolve() (end bool) {

	b.PrintStats() // Stats on top of the board
	// TODO: Add a print for the setup of the board letting the user know the parameters set for the current simulation

	b.PolliceVerso()
	b.PrintBoard()

	b.stats.generation += 1

	end = b.CollectPopulationStat()

	return end

}

func (b *Board) PolliceVerso() {

	neighboursMap := make([][]int, b.height)

	for i := range b.cells {
		neighboursMap[i] = make([]int, b.width)
		for j := range b.cells[i] {

			neighbours := CountLiveNeighbours(i, j, b)
			neighboursMap[i][j] = neighbours

		}
	}

	for i := range b.height {
		for j := range b.width {

			b.cells[i][j].Fate(neighboursMap[i][j])

		}
	}

}

func (b *Board) GetCell(i, j int) (status, valid bool) {

	iValid := i >= 0 && i < len(b.cells)
	jValid := j >= 0 && j < len(b.cells[0])

	if iValid && jValid {
		return b.cells[i][j].status, true
	}

	return false, false

}

func CountLiveNeighbours(i, j int, b *Board) (neighboursNumber int) {

	count := 0

	directions := []Coordinate{
		{i: -1, j: -1}, {i: -1, j: +0}, {i: -1, j: +1}, // 1 ROW ABOVE
		{i: +0, j: -1}, {i: +0, j: +1}, // SAME ROW
		{i: +1, j: -1}, {i: +1, j: +0}, {i: +1, j: +1}, // 1 ROW DOWN
	}

	for _, c := range directions {

		if status, valid := b.GetCell(i+c.i, j+c.j); valid {
			if status {
				count += 1
			}
		}

	}

	return count
}

func (b *Board) PrintBoard() {

	for i := range b.cells {
		representation := ""
		for j := range b.cells[i] {
			representation += b.cells[i][j].String()
		}
		b.cfg.Println(representation)
	}

}

func (b *Board) CollectPopulationStat() (end bool) {

	population := 0

	for i := range b.cells {
		for j := range b.cells[i] {
			if b.cells[i][j].status {
				population += 1
			}
		}
	}

	b.stats.population = append(b.stats.population, float64(population))

	if len(b.stats.population) > 1 {
		populationChange := float64(population) - b.stats.population[len(b.stats.population)-2]
		b.stats.populationChange = append(b.stats.populationChange, populationChange)
		if len(b.stats.populationChange) > 3 {
			static := true
			for _, pc := range b.stats.populationChange[len(b.stats.populationChange)-3:] {
				static = pc == 0 && static
			}
			return static
		}
	}

	return population == 0

}

func (b *Board) PrintStats() {

	b.cfg.Println(
		ColoredString(
			ApplyStyle(
				"Stats",
				TextStyle.BOLD,
				TextStyle.REVERSE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	b.cfg.Println(
		ApplyStyle("* # Generations: ", TextStyle.BOLD),
		ColoredString(
			ApplyStyle(
				fmt.Sprintf("%v", b.stats.generation),
				TextStyle.BOLD,
				TextStyle.UNDERLINE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	b.cfg.Println(
		ApplyStyle("* Current Population: ", TextStyle.BOLD),
		ColoredString(
			ApplyStyle(
				fmt.Sprintf("%v", b.stats.population[len(b.stats.population)-1]),
				TextStyle.BOLD,
				TextStyle.UNDERLINE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	if len(b.stats.populationChange) > 0 {
		b.cfg.Println(
			ApplyStyle("* Delta Population: ", TextStyle.BOLD),
			ColoredString(
				ApplyStyle(
					fmt.Sprintf("%v", b.stats.populationChange[len(b.stats.populationChange)-1]),
					TextStyle.BOLD,
					TextStyle.UNDERLINE,
				),
				THEME.Secondary,
				THEME.Background,
			),
		)
	}

	b.cfg.Println()

}

func (b *Board) HandleKeyStrokes(keysEvents <-chan keyboard.KeyEvent) {

	defer keyboard.Close()

	event := <-keysEvents
	if event.Err != nil {
		panic(event.Err)
	}
	// fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)

	if !b.started { // Handle selection event only if simulation hasn't started yet
		if event.Rune == 's' {
			b.Move(&b.selectedCell.i, 1, b.height)
			return
		}

		if event.Rune == 'w' {
			b.Move(&b.selectedCell.i, -1, b.height)
			return
		}

		if event.Rune == 'd' {
			b.Move(&b.selectedCell.j, 1, b.width)
			return
		}

		if event.Rune == 'a' {
			b.Move(&b.selectedCell.j, -1, b.width)
			return
		}

		if event.Rune != 'a' && event.Rune != 'w' && event.Rune != 's' && event.Rune != 'd' && event.Rune != '\x00' { // All other letters
			b.started = true
			b.CleanSelection()
			return
		}

		switch event.Key {
		case keyboard.KeySpace:
			b.SwitchCellStatus()
		case keyboard.KeyEsc:
			fallthrough
		case keyboard.KeyCtrlC:
			fallthrough
		case keyboard.KeyEnter:
			b.started = true
			b.CleanSelection()
			fallthrough
		default:
			return
		}

	}

}

func (b *Board) Move(axys *int, direction, upperLimit int) {

	if *axys+direction >= 0 && *axys+direction < upperLimit {
		b.SwitchCellSelection()
		*axys += direction
		b.SwitchCellSelection()
	}

}

func (b *Board) PrintSummary() {

	b.cfg.Println(
		ColoredString(
			ApplyStyle(
				"Stats",
				TextStyle.BOLD,
				TextStyle.REVERSE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	b.cfg.Println(
		ApplyStyle("* # Generations: ", TextStyle.BOLD),
		ColoredString(
			ApplyStyle(
				fmt.Sprintf("%v", b.stats.generation),
				TextStyle.BOLD,
				TextStyle.UNDERLINE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	b.cfg.Println(
		ApplyStyle("* Initial Population -> Final Population: ", TextStyle.BOLD),
		ColoredString(
			ApplyStyle(
				fmt.Sprintf("%v -> %v", b.stats.population[0], b.stats.population[len(b.stats.population)-1]),
				TextStyle.BOLD,
				TextStyle.UNDERLINE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	if len(b.stats.populationChange) > 0 {
		b.cfg.Println(
			ApplyStyle("* Final Delta Population: ", TextStyle.BOLD),
			ColoredString(
				ApplyStyle(
					fmt.Sprintf("%v", b.stats.population[len(b.stats.population)-1]-b.stats.population[0]),
					TextStyle.BOLD,
					TextStyle.UNDERLINE,
				),
				THEME.Secondary,
				THEME.Background,
			),
		)
	}

	b.cfg.Println()

	b.PrintBoard()
}

func (b *Board) PrintSetUpInstructions() {
	b.cfg.Println(ColoredString(ApplyStyle("Simulation Set Up", TextStyle.BOLD, TextStyle.REVERSE), THEME.Secondary, THEME.Background))
	b.cfg.Println("Commands:")

	b.cfg.Println("    - ", ColoredString(ApplyStyle("'a'", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": move to the cell to the left;")
	b.cfg.Println("    - ", ColoredString(ApplyStyle("'w'", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": move to the cell above;")
	b.cfg.Println("    - ", ColoredString(ApplyStyle("'s'", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": move to the cell below;")
	b.cfg.Println("    - ", ColoredString(ApplyStyle("'d'", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": move to the cell to the right;")
	b.cfg.Println("    - ", ColoredString(ApplyStyle("space bar", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": toggle cell status (alive/dead);")
	b.cfg.Println("    - ", ColoredString(ApplyStyle("any other key", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": start simulation;")
	b.cfg.Println()
}

func (b *Board) PrintBanner(banner []string) {
	for _, line := range banner {
		b.cfg.Println(line)
	}
}
