// commandlineplayer
package main

import (
	"chess/game"
	"strconv"
	"strings"
)

type CommandLinePlayer struct {
	side chessgame.Side
}

func (player CommandLinePlayer) MakeMove(boardClone chessgame.ChessBoard, validMoves map[chessgame.Coordinate]map[chessgame.Coordinate]bool) (chessgame.Coordinate, chessgame.Coordinate) {
	// PrintBoard(boardClone, player.GetSide())

	return parseMove("a7->a5")
}

func (player CommandLinePlayer) PromotePawn() chessgame.PieceType {
	return chessgame.QUEEN
}

func (player CommandLinePlayer) GetSide() chessgame.Side {
	return player.side
}

func parseMove(userInput string) (chessgame.Coordinate, chessgame.Coordinate) {
	splitCoords := strings.Split(userInput, "->")
	return parseCoord(splitCoords[0]), parseCoord(splitCoords[1])
}

func parseCoord(stringCoord string) chessgame.Coordinate {
	row, _ := strconv.Atoi(string(stringCoord[1]))
	col := letterToIndex(string(stringCoord[0]))
	return chessgame.Coordinate{row - 1, col}
}

func letterToIndex(letter string) int {
	if letter == "a" {
		return 7
	}
	if letter == "b" {
		return 6
	}
	if letter == "c" {
		return 5
	}
	if letter == "d" {
		return 4
	}
	if letter == "e" {
		return 3
	}
	if letter == "f" {
		return 2
	}
	if letter == "g" {
		return 1
	}
	return 0
}
