package main

import (
	"fmt"
	"strconv"
)

type Movement struct{
	x int
	y int
	direction string
}

type GamePiece struct{
	name string
	captured bool
	Movement
}

type Square struct{
	gamePiece *GamePiece
}

func (s Square) addGamePiece(gamepiece *GamePiece){
	s.gamePiece = gamepiece
}

func createBoard() (map[string]Square) {
	gameboard := make(map[string]Square)
	for _, letter := range "ABCDEFGH" {
		for num := 1; num <= 8; num++ {
			fmt.Println(num,string(letter))
			gameboard[string(letter)+strconv.Itoa(num)] = Square{}
		}
	}
	return gameboard
}