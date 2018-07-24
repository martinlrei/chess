// movesimulator
package chessengine

import (
	"chess/game"
)

func selectBestMove(side chessgame.Side, board *chessgame.ChessBoard, validMoves map[chessgame.Coordinate]map[chessgame.Coordinate]bool) (chessgame.Coordinate, chessgame.Coordinate) {
	maxScore := 0
	first := true
	var bestFromCoord chessgame.Coordinate
	var bestToCoord chessgame.Coordinate
	for from, allToCoords := range validMoves {
		for to, _ := range allToCoords {
			formerToPiece := board.BoardPieces[to.Row][to.Column]
			board.BoardPieces[to.Row][to.Column] = board.BoardPieces[from.Row][from.Column]
			board.BoardPieces[from.Row][from.Column] = nil
			score := 0
			if board.BoardPieces[to.Row][to.Column] != nil && board.GetPieceSide(to) != side {
				score += pieceValue(board.BoardPieces[to.Row][to.Column]) * 20
			}
			score += evaluateBoard(board, side)
			if score >= maxScore || first {
				maxScore = score
				bestFromCoord = from
				bestToCoord = to
				first = false
			}
			board.BoardPieces[from.Row][from.Column] = board.BoardPieces[to.Row][to.Column]
			board.BoardPieces[to.Row][to.Column] = formerToPiece
		}
	}
	return bestFromCoord, bestToCoord
}
