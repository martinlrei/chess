package main

import (
	"chess/game"
	"fmt"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	walk.Resources.SetRootDirPath("images")

	var widgets []Widget

	// primes := []string{"BB.png", "BB.png", "BB.png", "BB.png", "BB.png", "BB.png", "BB.png", "BB.png", "BB.png"}

	humanSide := chessgame.WHITE
	aiSide := chessgame.BLACK

	aiPlayer := NewCommandLineAI(aiSide)
	humanPlayer := CommandLinePlayer{humanSide}

	game := chessgame.NewChessGame(aiPlayer, humanPlayer)

	for _, sqaure := range game.Board.BoardPieces {
		for _, piece := range sqaure {
			widgets = append(widgets,
				ImageView{
					Background: SolidColorBrush{Color: walk.RGB(255, 191, 0)},
					Image:      "BB.png",
					Mode:       ImageViewModeIdeal,
				},
			)

			fmt.Println(piece.GetPieceType())
		}
	}

	MainWindow{
		Title:    "Walk ImageView Example",
		Size:     Size{600, 600},
		Layout:   Grid{Columns: 8},
		Children: widgets,
	}.Run()
}
