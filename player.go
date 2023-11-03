package main

import (
	"fmt"
	"strconv"
	"strings"
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
func (p Player) selectMoveWithString(startingPosition, landingPosition string) Move{
	startingXInt := convertFirstCharacterToInt(startingPosition[0])
	startingYInt, _ := strconv.Atoi(string(startingPosition[1]))
	endingXInt := convertFirstCharacterToInt(landingPosition[0])
	endingYInt, _ := strconv.Atoi(string(landingPosition[1]))
	return p.selectMove(Position{int(startingXInt),startingYInt}, Position{endingXInt,endingYInt})
}

func convertFirstCharacterToInt(char byte) int{
	for i, letter := range "ABCDEFGH"{
		if strings.ToUpper(string(char)) == string(letter) {
			return i + 1
		}
	}
	int, _ := strconv.Atoi(string(char))
	return int
}
