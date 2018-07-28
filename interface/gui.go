package main

import (
	"chess/game"

	"github.com/lxn/walk"
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

func getImage(piecenumber chessgame.ChessPiece) string {

	var result string
	if piecenumber.GetPieceSide() == chessgame.WHITE {
		result = "W"
	} else {
		result = "B"
	}

	switch piecenumber.GetPieceType() {
	case chessgame.ROOK:
		result += "R"
	case chessgame.PAWN:
		result += "P"
	case chessgame.KNIGHT:
		result += "N"
	case chessgame.KING:
		result += "K"
	case chessgame.QUEEN:
		result += "Q"
	case chessgame.BISHOP:
		result += "B"
	default:
		result = "BR"
	}

	return result + ".png"
}

func getSquare(row int, col int) string {

	var result string
	// 00 white 01 red   02 white 03 red 04 white 05 red 06 white 07 red
	// 10 red   11 white 12 red   13 white 14 red 15 white 16 red 17 white
	// 20 white 21 red   22 white 23 red 24 white 25 red 26 white 27 red
	// 30 red   31 white 32 red   33 white 34 red 35 white 36 red 37 white

	if (row+col)%2 == 0 {
		result = "WhiteSquare.png"
	} else {
		result = "RedSquare.png"
	}

	return result
}

func imageBackgroundColor(row int, col int) walk.Color {
	// 	forestgreen	#228B22	rgb(34,139,34)
	// OLD BRICK 150, 40, 27

	var result walk.Color
	if (row+col)%2 == 0 {
		result = walk.RGB(255, 255, 255)
	} else {
		result = walk.RGB(150, 40, 27)
	}

	return result
}
