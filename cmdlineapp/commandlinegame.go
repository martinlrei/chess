// main
package main

import (
	"chess/engine"
	"chess/game"
	"fmt"
)

func main() {
	fmt.Println("Pick a side. Enter \"white\" or \"black\"")
	var feedBack string
	fmt.Scanln(&feedBack)
	for feedBack != "white" && feedBack != "black" {
		fmt.Scanln(&feedBack)
	}
	fmt.Println("Pick a difficulty. Enter \"basic\" \"easy\" \"medium\" \"hard\" or \"expert\"")
	var difficultyString string
	fmt.Scanln(&difficultyString)
	for difficultyString != "basic" && difficultyString != "easy" && difficultyString != "medium" && difficultyString != "hard" && difficultyString != "expert" {
		fmt.Scanln(&difficultyString)
	}
	difficulty := difficultyFromString(difficultyString)
	humanSide := chessgame.WHITE
	aiSide := chessgame.BLACK
	if feedBack == "black" {
		humanSide = chessgame.BLACK
		aiSide = chessgame.WHITE
	}
	aiPlayer := NewCommandLineAI(aiSide, difficulty)
	humanPlayer := CommandLinePlayer{humanSide}
	game := chessgame.NewChessGame(aiPlayer, humanPlayer)
	if aiPlayer.GetSide() == chessgame.BLACK {
		game = chessgame.NewChessGame(humanPlayer, aiPlayer)
	}
	outCome := game.PlayGame()
	PrintBoard(*game.Board, humanPlayer.GetSide())
	if outCome == chessgame.WHITEVICTORY {
		fmt.Println("Checkmate! White is victorious")
	} else if outCome == chessgame.BLACKVICTORY {
		fmt.Println("Checkmate! Black is victorious")
	} else {
		fmt.Println("Stalemate!")
	}
}

func PrintBoard(board chessgame.ChessBoard, side chessgame.Side) {
	colString := "\t    h   g   f   e   d   c   b   a  "
	if side == chessgame.WHITE {
		colString = "\t    a   b   c   d   e   f   g   h  "
	}
	initial := 0
	boundary := 8
	if side == chessgame.WHITE {
		initial = 7
		boundary = -1
	}
	fmt.Println(colString)
	fmt.Println("\t  ---------------------------------")
	row := initial
	for row != boundary {
		fmt.Printf("\t%v | ", row+1)
		col := initial
		for col != boundary {
			PrintPiece(board.BoardPieces[row][col])
			fmt.Print("| ")
			if side == chessgame.WHITE {
				col--
			} else {
				col++
			}
		}
		fmt.Print("\n")
		fmt.Println("\t  ---------------------------------")
		if side == chessgame.WHITE {
			row--
		} else {
			row++
		}
	}
}

func PrintPiece(piece chessgame.ChessPiece) {
	if piece == nil {
		fmt.Printf("  ")
	} else if piece.GetPieceSide() == chessgame.BLACK && piece.GetPieceType() == chessgame.PAWN {
		fmt.Print("bP")
	} else if piece.GetPieceSide() == chessgame.WHITE && piece.GetPieceType() == chessgame.PAWN {
		fmt.Printf("wP")
	} else if piece.GetPieceSide() == chessgame.BLACK && piece.GetPieceType() == chessgame.ROOK {
		fmt.Printf("bR")
	} else if piece.GetPieceSide() == chessgame.WHITE && piece.GetPieceType() == chessgame.ROOK {
		fmt.Printf("wR")
	} else if piece.GetPieceSide() == chessgame.BLACK && piece.GetPieceType() == chessgame.KNIGHT {
		fmt.Printf("bN")
	} else if piece.GetPieceSide() == chessgame.WHITE && piece.GetPieceType() == chessgame.KNIGHT {
		fmt.Printf("wN")
	} else if piece.GetPieceSide() == chessgame.BLACK && piece.GetPieceType() == chessgame.BISHOP {
		fmt.Printf("bB")
	} else if piece.GetPieceSide() == chessgame.WHITE && piece.GetPieceType() == chessgame.BISHOP {
		fmt.Printf("wB")
	} else if piece.GetPieceSide() == chessgame.BLACK && piece.GetPieceType() == chessgame.QUEEN {
		fmt.Printf("bQ")
	} else if piece.GetPieceSide() == chessgame.WHITE && piece.GetPieceType() == chessgame.QUEEN {
		fmt.Printf("wQ")
	} else if piece.GetPieceSide() == chessgame.BLACK && piece.GetPieceType() == chessgame.KING {
		fmt.Printf("bK")
	} else if piece.GetPieceSide() == chessgame.WHITE && piece.GetPieceType() == chessgame.KING {
		fmt.Printf("wK")
	}
}

func difficultyFromString(difficultyStr string) chessengine.EngineDifficulty {
	if difficultyStr == "basic" {
		return chessengine.BASIC
	}
	if difficultyStr == "easy" {
		return chessengine.EASY
	}
	if difficultyStr == "medium" {
		return chessengine.MEDIUM
	}
	if difficultyStr == "hard" {
		return chessengine.HARD
	}
	return chessengine.EXPERT
}
