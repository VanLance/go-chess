package main

import (
	"fmt"
	// "strconv"
)


type Player struct {
	name string
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var i string
	fmt.Scan(&i)
	return i
}

func (p Player) selectMove(startingPosition, landingPosition Position) Move{
	return Move{startingPosition: startingPosition, endingPosition: landingPosition, player: p}
}

