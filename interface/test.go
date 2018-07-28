package main

import (
	"chess/game"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	walk.Resources.SetRootDirPath("images")

	var widgets []Widget
	// 	forestgreen	#228B22	rgb(34,139,34)
	// OLD BRICK 150, 40, 27
	//primes := []string{"BB.png", "BB.png", "BB.png", "BB.png", "BB.png", "BB.png", "BB.png", "BB.png", "BB.png"}

	humanSide := chessgame.WHITE
	aiSide := chessgame.BLACK

	aiPlayer := NewCommandLineAI(aiSide)
	humanPlayer := CommandLinePlayer{humanSide}

	game := chessgame.NewChessGame(aiPlayer, humanPlayer)
	var row int
	var col int
	row = 0
	col = 0
	for _, sqaure := range game.Board.BoardPieces {
		for _, piece := range sqaure {
			if piece == nil {
				widgets = append(widgets,
					ImageView{
						Background: SolidColorBrush{Color: imageBackgroundColor(row, col)},
						Image:      getSquare(row, col),
						Margin:     10,
						Mode:       ImageViewModeIdeal,
					},
				)
			} else {
				widgets = append(widgets,
					ImageView{
						Background: SolidColorBrush{Color: imageBackgroundColor(row, col)},
						Image:      getImage(piece),
						Margin:     10,
						Mode:       ImageViewModeIdeal,
					},
				)
			}
			if col == 7 {
				col = 0
				row++
			} else {
				col++
			}
		}
	}

	MainWindow{
		Title:    "Walk ImageView Example",
		Size:     Size{600, 600},
		Layout:   Grid{Columns: 8},
		Children: widgets,
	}.Run()
}
