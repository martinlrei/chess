// boardevaluator
package chessengine

import (
	"chess/game"
)

func evaluateBoard(board *chessgame.ChessBoard, side chessgame.Side) int {
	score := 0
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if board.BoardPieces[row][col] != nil && board.BoardPieces[row][col].GetPieceSide() == side {
				score += evaluatePiece(board, board.BoardPieces[row][col])
			}
		}
	}
	return score
}

func evaluatePiece(board *chessgame.ChessBoard, piece chessgame.ChessPiece) int {
	totalPieceScore := 0
	validMoves := piece.ValidMoves(board)
	totalPieceScore += len(validMoves)
	for move := range validMoves {
		if board.BoardPieces[move.Row][move.Column] != nil && board.GetPieceSide(move) != piece.GetPieceSide() {
			totalPieceScore += pieceValue(board.BoardPieces[move.Row][move.Column])
		}
	}
	if isInMiddle(piece) && piece.GetPieceType() == chessgame.KNIGHT || piece.GetPieceType() == chessgame.PAWN {
		totalPieceScore += 15
	}
	totalPieceScore -= len(chessgame.GetThreateningCoordinates(board, piece.GetCurrentCoordinates(), piece.GetPieceSide())) * pieceValue(piece)
	// get supporting pieces by calling GetThreateningCoordinates for piece's side but opposite coordinates
	oppositePieceSide := chessgame.WHITE
	if piece.GetPieceSide() == chessgame.WHITE {
		oppositePieceSide = chessgame.BLACK
	}
	totalPieceScore += len(chessgame.GetThreateningCoordinates(board, piece.GetCurrentCoordinates(), oppositePieceSide))
	return totalPieceScore
}

func isInMiddle(piece chessgame.ChessPiece) bool {
	coord := piece.GetCurrentCoordinates()
	if coord.Row <= 5 && coord.Row >= 3 && coord.Column <= 5 && coord.Column >= 3 {
		return true
	}
	return false
}

func pieceValue(piece chessgame.ChessPiece) int {
	if piece.GetPieceType() == chessgame.PAWN {
		return 1
	}
	if piece.GetPieceType() == chessgame.KNIGHT {
		return 3
	}
	if piece.GetPieceType() == chessgame.BISHOP {
		return 4
	}
	if piece.GetPieceType() == chessgame.ROOK {
		return 5
	}
	if piece.GetPieceType() == chessgame.QUEEN {
		return 7
	}
	return 8
}
