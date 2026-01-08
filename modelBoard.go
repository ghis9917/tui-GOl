package main

import (
	"fmt"
	"math/rand"

	"github.com/eiannone/keyboard"
	"github.com/guptarohit/asciigraph"
)

type Board struct {
	cells        [][]Cell
	height       int
	width        int
	stats        Stats
	started      bool
	selectedCell Coordinate
}

type Stats struct {
	generation       int
	population       []float64
	populationChange []float64
}

func NewBoard(width, height int) *Board {

	b := &Board{
		cells:  make([][]Cell, height),
		height: height,
		width:  width,
		stats: Stats{
			generation:       0,
			population:       []float64{},
			populationChange: []float64{},
		},
		started:      false,
		selectedCell: Coordinate{i: 0, j: 0},
	}
	limit := MAX_INITIAL_POPULATION

	for i := range b.cells {
		b.cells[i] = make([]Cell, width)
		for j := range b.cells[i] {
			if rand.Intn(SPARSITY) == 1 && limit > 0 {
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

	b.PolliceVerso()
	b.PrintBoard()

	b.stats.generation += 1

	end = b.CollectPopulationStat()
	b.PrintStats()

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

func CountLiveNeighbours(i, j int, b *Board) int {

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

	representation := ""

	for i := range b.cells {
		for j := range b.cells[i] {
			representation += b.cells[i][j].String()
		}
		representation += "\n"
	}

	fmt.Println(representation)

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
	}

	return population == 0

}

func (b *Board) PrintStats() {

	fmt.Printf("\t* # Generations: %v\n", b.stats.generation)
	currentPopulation := b.stats.population[len(b.stats.population)-1]
	fmt.Printf("\t* Current Population: %v\n", currentPopulation)
	deltaPopulation := b.stats.populationChange[len(b.stats.populationChange)-1]
	fmt.Printf("\t* Delta Population: %v\n", deltaPopulation)

	graph := asciigraph.Plot(b.stats.population, asciigraph.Height(5), asciigraph.Width(WIDTH))
	fmt.Print(graph)

}

func (b *Board) HandleKeyStrokes(keysEvents <-chan keyboard.KeyEvent) {

	defer keyboard.Close()

	event := <-keysEvents
	if event.Err != nil {
		panic(event.Err)
	}
	fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)

	if !b.started { // Handle selection event only if simulation hasn't started yet
		if event.Rune == 's' {
			if b.selectedCell.i+1 < len(b.cells) {
				b.SwitchCellSelection()
				b.selectedCell.i += 1
				b.SwitchCellSelection()
			}
			return
		}

		if event.Rune == 'w' {
			if b.selectedCell.i-1 >= 0 {
				b.SwitchCellSelection()
				b.selectedCell.i -= 1
				b.SwitchCellSelection()
			}
			return
		}

		if event.Rune == 'd' {
			if b.selectedCell.j+1 < len(b.cells[0]) {
				b.SwitchCellSelection()
				b.selectedCell.j += 1
				b.SwitchCellSelection()
			}
			return
		}

		if event.Rune == 'a' {
			if b.selectedCell.j-1 >= 0 {
				b.SwitchCellSelection()
				b.selectedCell.j -= 1
				b.SwitchCellSelection()
			}
			return
		}
	}

	if event.Key == keyboard.KeyEsc || event.Key == keyboard.KeyCtrlC {
		b.started = true
		b.CleanSelection()
		return
	}

	switch event.Key {
	case keyboard.KeyEnter:
		b.started = true
		b.CleanSelection()
	case keyboard.KeySpace:
		b.SwitchCellStatus()
	default:
		return
	}

}
