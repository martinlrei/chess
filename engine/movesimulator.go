// movesimulator
package chessengine

import (
	"chess/game"
)

func selectBestMove(side chessgame.Side, board *chessgame.ChessBoard, validMoves map[chessgame.Coordinate]map[chessgame.Coordinate]bool, depth int) (chessgame.Coordinate, chessgame.Coordinate) {
	currentScore := 0
	first := true
	var bestFromCoord chessgame.Coordinate
	var bestToCoord chessgame.Coordinate
	useMin := (depth % 2) == 1
	for from, allToCoords := range validMoves {
		for to, _ := range allToCoords {
			formerToPiece := board.BoardPieces[to.Row][to.Column]
			board.BoardPieces[to.Row][to.Column] = board.BoardPieces[from.Row][from.Column]
			board.BoardPieces[from.Row][from.Column] = nil
			score := evaluateAllMoves(getOppositeSide(side), board, useMin, depth, 1)
			if board.BoardPieces[to.Row][to.Column] != nil && board.GetPieceSide(to) != side {
				if useMin {
					score -= pieceValue(board.BoardPieces[to.Row][to.Column]) * 10
				} else {
					score += pieceValue(board.BoardPieces[to.Row][to.Column]) * 10
				}
			}
			currentScore = compareScores(useMin, currentScore, score, first)
			if score == currentScore {
				currentScore = score
				bestFromCoord = from
				bestToCoord = to
			}
			first = false
			board.BoardPieces[from.Row][from.Column] = board.BoardPieces[to.Row][to.Column]
			board.BoardPieces[to.Row][to.Column] = formerToPiece
		}
	}
	return bestFromCoord, bestToCoord
}

func evaluateAllMoves(side chessgame.Side, board *chessgame.ChessBoard, useMin bool, maxDepth int, currentDepth int) int {
	if maxDepth == currentDepth {
		return evaluateBoard(board, side)
	}
	score := 0
	first := true
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if board.BoardPieces[row][col] != nil && board.BoardPieces[row][col].GetPieceSide() == side {
				moves := board.BoardPieces[row][col].ValidMoves(board)
				for toCoord, _ := range moves {
					fromCoord := chessgame.Coordinate{row, col}
					formerToPiece := board.BoardPieces[toCoord.Row][toCoord.Column]
					board.BoardPieces[toCoord.Row][toCoord.Column] = board.BoardPieces[fromCoord.Row][fromCoord.Column]
					board.BoardPieces[fromCoord.Row][fromCoord.Column] = nil
					oppositeSide := getOppositeSide(side)
					moveScore := evaluateAllMoves(oppositeSide, board, !useMin, maxDepth, currentDepth+1)
					score = compareScores(useMin, score, moveScore, first)
					first = false
					board.BoardPieces[fromCoord.Row][fromCoord.Column] = board.BoardPieces[toCoord.Row][toCoord.Column]
					board.BoardPieces[toCoord.Row][toCoord.Column] = formerToPiece
				}
			}
		}
	}
	return score
}

func compareScores(useMin bool, currentScore int, challengerScore int, first bool) int {
	if first {
		return challengerScore
	}
	if useMin {
		return getMin(currentScore, challengerScore)
	}
	return getMax(currentScore, challengerScore)
}

func getMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func getOppositeSide(side chessgame.Side) chessgame.Side {
	if side == chessgame.WHITE {
		return chessgame.BLACK
	}
	return chessgame.WHITE
}
