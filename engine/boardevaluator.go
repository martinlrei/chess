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
	if isInMiddle(piece) && piece.GetPieceType() == chessgame.KNIGHT || piece.GetPieceType() == chessgame.PAWN || piece.GetPieceType() == chessgame.BISHOP {
		totalPieceScore += 15
	}
	totalPieceScore -= len(chessgame.GetThreateningCoordinates(board, piece.GetCurrentCoordinates(), piece.GetPieceSide())) * pieceValue(piece) * 10
	// get supporting pieces by calling GetThreateningCoordinates for piece's side but opposite coordinates
	oppositePieceSide := chessgame.WHITE
	if piece.GetPieceSide() == chessgame.WHITE {
		oppositePieceSide = chessgame.BLACK
	}
	totalPieceScore += len(chessgame.GetThreateningCoordinates(board, piece.GetCurrentCoordinates(), oppositePieceSide)) * 7
	totalPieceScore += kingInBackRow(piece)
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
		return 4
	}
	if piece.GetPieceType() == chessgame.KNIGHT {
		return 10
	}
	if piece.GetPieceType() == chessgame.BISHOP {
		return 11
	}
	if piece.GetPieceType() == chessgame.ROOK {
		return 15
	}
	if piece.GetPieceType() == chessgame.QUEEN {
		return 25
	}
	return 40
}

func kingInBackRow(piece chessgame.ChessPiece) int {
	if piece.GetPieceType() != chessgame.KING {
		return 0
	}
	coords := piece.GetCurrentCoordinates()
	var baseRow int
	if piece.GetPieceSide() == chessgame.WHITE {
		baseRow = 0
	} else {
		baseRow = 7
	}
	if coords.Row == baseRow {
		return 25
	}
	return -5
}
