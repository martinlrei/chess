package main

import (
	"chess/engine"
	"chess/game"
)

type CommandLineAI struct {
	side chessgame.Side
}

func (ai CommandLineAI) MakeMove(boardClone chessgame.ChessBoard, validMoves map[chessgame.Coordinate]map[chessgame.Coordinate]bool) (chessgame.Coordinate, chessgame.Coordinate) {
	/*side := chessgame.WHITE

	if ai.GetSide() == chessgame.WHITE {
		side = chessgame.BLACK
	}*/
	// PrintBoard(boardClone, side)

	return chessengine.MakeRandomMove(validMoves)
}

func (ai CommandLineAI) PromotePawn() chessgame.PieceType {
	return chessgame.QUEEN
}

func (ai CommandLineAI) GetSide() chessgame.Side {
	return ai.side
}

func NewCommandLineAI(side chessgame.Side) CommandLineAI {
	return CommandLineAI{side}
}
