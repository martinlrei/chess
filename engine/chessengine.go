// chessengineplayer
package chessengine

import (
	"chess/game"
	"math/rand"
)

type EngineDifficulty int

const (
	BASIC EngineDifficulty = iota + 1
	EASY
	MEDIUM
	HARD
	EXPERT
)

func MakeEvaluatedMove(side chessgame.Side, board *chessgame.ChessBoard, validMoves map[chessgame.Coordinate]map[chessgame.Coordinate]bool) (chessgame.Coordinate, chessgame.Coordinate) {
	return selectBestMove(side, board, validMoves, 3)
}

func MakeRandomMove(validMoves map[chessgame.Coordinate]map[chessgame.Coordinate]bool) (chessgame.Coordinate, chessgame.Coordinate) {
	randNum := rand.Intn(len(validMoves))
	counter := 0
	var fromCoord chessgame.Coordinate
	var toCoord chessgame.Coordinate
	for k, _ := range validMoves {
		if counter == randNum {
			fromCoord = k
			break
		}
		counter++
	}
	counter = 0
	randNum = rand.Intn(len(validMoves[fromCoord]))
	for k, _ := range validMoves[fromCoord] {
		if counter == randNum {
			toCoord = k
			break
		}
		counter++
	}
	return fromCoord, toCoord
}
