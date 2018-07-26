package main

import (
	"chess/game"
)

func initialize() {

	// TODO: We should create a GUI that allows the user to choose white or black
	// For now we can default to WHITE for Human Play

	//for feedBack != "white" && feedBack != "black"

	humanSide := chessgame.WHITE
	aiSide := chessgame.BLACK

	aiPlayer := NewCommandLineAI(aiSide)
	humanPlayer := CommandLinePlayer{humanSide}

	game := chessgame.NewChessGame(aiPlayer, humanPlayer)
	outCome := game.PlayGame()

	//TODO: Display the board to the GUI PrintBoard(*game.Board, humanPlayer.GetSide())

	//We should show a message box and start a new game and then refresh the screen when we have an outcome
	if outCome == chessgame.WHITEVICTORY {
		//fmt.Println("Checkmate! White is victorious")
	} else if outCome == chessgame.BLACKVICTORY {
		//fmt.Println("Checkmate! Black is victorious")
	} else {
		//fmt.Println("Stalemate!")
	}

}
