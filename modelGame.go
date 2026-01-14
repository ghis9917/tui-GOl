package main

import (
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
)

type GameOfLife struct {
	started bool
	board   *Board
	stats   Stats
	cfg     Config
}

func NewGame(cfg Config) (newGame *GameOfLife) {

	g := &GameOfLife{
		started: false,
		board:   NewBoard(cfg),
		cfg:     cfg,
		stats: Stats{
			generation:       -1,
			population:       []float64{},
			populationChange: []float64{},
		},
	}

	return g

}

func (g *GameOfLife) SetUp() {

	for !g.started {

		g.PrintScreen(BANNER, g.WriteSetUpInstructions)

		keysEvents, err := keyboard.GetKeys(10)
		if err != nil {
			log.Fatal(err)
		}
		g.HandleKeyStrokes(keysEvents)

		time.Sleep(UI_REFRESH_RATE)
	}

	g.UpdateStats(g.board.GetPopulation()) // Initialize stats on population after the setup has been completed

}

func (g *GameOfLife) Run() {

	for g.started {

		g.board.PolliceVerso()
		newPopulation := g.board.GetPopulation()
		end := g.UpdateStats(newPopulation)

		g.PrintScreen(BANNER, g.WriteStats)

		if end {
			break
		}

		time.Sleep(time.Second / time.Duration(g.cfg.fps))
	}

	g.PrintScreen(END_BANNER, g.WriteSummary)

}

func (g *GameOfLife) UpdateStats(newPopulation float64) (end bool) {

	g.stats.generation += 1

	g.stats.population = append(g.stats.population, newPopulation)

	if len(g.stats.population) > 1 {
		populationChange := newPopulation - g.stats.population[len(g.stats.population)-2]
		g.stats.populationChange = append(g.stats.populationChange, populationChange)
		if len(g.stats.populationChange) > 3 {
			static := true
			for _, pc := range g.stats.populationChange[len(g.stats.populationChange)-3:] {
				static = pc == 0 && static
			}
			return static
		}
	}

	return newPopulation == 0

}

func (g *GameOfLife) WriteBanner(banner []string) {
	g.cfg.Println()

	for _, line := range banner {
		g.cfg.Println(line)
	}

	g.cfg.Println()
	g.cfg.PrintConfig()
}

func (g *GameOfLife) WriteSetUpInstructions() {

	g.cfg.Println(ColoredString(ApplyStyle("Simulation Set Up", TextStyle.BOLD, TextStyle.REVERSE), THEME.Secondary, THEME.Background))
	g.cfg.Println("Commands:")

	g.cfg.Println("    - ", ColoredString(ApplyStyle("'a'", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": move to the cell to the left;")
	g.cfg.Println("    - ", ColoredString(ApplyStyle("'w'", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": move to the cell above;")
	g.cfg.Println("    - ", ColoredString(ApplyStyle("'s'", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": move to the cell below;")
	g.cfg.Println("    - ", ColoredString(ApplyStyle("'d'", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": move to the cell to the right;")
	g.cfg.Println("    - ", ColoredString(ApplyStyle("space bar", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": toggle cell status (alive/dead);")
	g.cfg.Println("    - ", ColoredString(ApplyStyle("any other key", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": start simulation;")
	g.cfg.Println()
}

func (g *GameOfLife) WriteStats() {

	g.cfg.Println(
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

	g.cfg.Println(
		ApplyStyle("* # Generations: ", TextStyle.BOLD),
		ColoredString(
			ApplyStyle(
				fmt.Sprintf("%v", g.stats.generation),
				TextStyle.BOLD,
				TextStyle.UNDERLINE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	g.cfg.Println(
		ApplyStyle("* Current Population: ", TextStyle.BOLD),
		ColoredString(
			ApplyStyle(
				fmt.Sprintf("%v", g.stats.population[len(g.stats.population)-1]),
				TextStyle.BOLD,
				TextStyle.UNDERLINE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	if len(g.stats.populationChange) > 0 {
		g.cfg.Println(
			ApplyStyle("* Delta Population: ", TextStyle.BOLD),
			ColoredString(
				ApplyStyle(
					fmt.Sprintf("%v", g.stats.populationChange[len(g.stats.populationChange)-1]),
					TextStyle.BOLD,
					TextStyle.UNDERLINE,
				),
				THEME.Secondary,
				THEME.Background,
			),
		)
	}

	g.cfg.Println()

}

func (g *GameOfLife) WriteSummary() {

	g.cfg.Println(
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

	g.cfg.Println(
		ApplyStyle("* # Generations: ", TextStyle.BOLD),
		ColoredString(
			ApplyStyle(
				fmt.Sprintf("%v", g.stats.generation),
				TextStyle.BOLD,
				TextStyle.UNDERLINE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	g.cfg.Println(
		ApplyStyle("* Initial Population -> Final Population: ", TextStyle.BOLD),
		ColoredString(
			ApplyStyle(
				fmt.Sprintf("%v -> %v", g.stats.population[0], g.stats.population[len(g.stats.population)-1]),
				TextStyle.BOLD,
				TextStyle.UNDERLINE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	if len(g.stats.populationChange) > 0 {
		g.cfg.Println(
			ApplyStyle("* Final Delta Population: ", TextStyle.BOLD),
			ColoredString(
				ApplyStyle(
					fmt.Sprintf("%v", g.stats.population[len(g.stats.population)-1]-g.stats.population[0]),
					TextStyle.BOLD,
					TextStyle.UNDERLINE,
				),
				THEME.Secondary,
				THEME.Background,
			),
		)
	}

	g.cfg.Println()
}

func (g *GameOfLife) WriteBoard(board []string) {
	for _, row := range board {
		g.cfg.Println(row)
	}
	g.cfg.FillVerticalSpace()
}

func (g *GameOfLife) HandleKeyStrokes(keysEvents <-chan keyboard.KeyEvent) {

	defer keyboard.Close()

	event := <-keysEvents
	if event.Err != nil {
		panic(event.Err)
	}
	// fmt.Printf("You pressed: rune %q, key %X\r\n", event.Rune, event.Key)

	if !g.started { // Handle selection event only if simulation hasn't started yet
		if event.Rune == 's' {
			g.board.Move(&g.board.selectedCell.i, 1, g.cfg.height)
			return
		}

		if event.Rune == 'w' {
			g.board.Move(&g.board.selectedCell.i, -1, g.cfg.height)
			return
		}

		if event.Rune == 'd' {
			g.board.Move(&g.board.selectedCell.j, 1, g.cfg.width)
			return
		}

		if event.Rune == 'a' {
			g.board.Move(&g.board.selectedCell.j, -1, g.cfg.width)
			return
		}

		if event.Rune != 'a' && event.Rune != 'w' && event.Rune != 's' && event.Rune != 'd' && event.Rune != '\x00' { // All other letters
			g.started = true
			g.board.CleanSelection()
			return
		}

		switch event.Key {
		case keyboard.KeySpace:
			g.board.SwitchCellStatus()
		case keyboard.KeyEsc:
			fallthrough
		case keyboard.KeyCtrlC:
			fallthrough
		case keyboard.KeyEnter:
			g.started = true
			g.board.CleanSelection()
			fallthrough
		default:
			return
		}

	}

}

func (g *GameOfLife) PrintScreen(banner []string, info func()) {
	g.cfg.ResetOutput()
	g.WriteBanner(banner)
	info()
	g.WriteBoard(g.board.Strings())
	g.cfg.Draw()
}
