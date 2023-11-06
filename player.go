package main

import (
	"fmt"
	"strconv"
	"strings"
)


type Player struct {
	Username string
	Team int
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)
	var i string
	fmt.Scan(&i)
	return i
}

func (p Player) selectMove(StartingPosition, LandingPosition Position) Move{
	fmt.Println("made it to SM")
	return Move{StartingPosition: StartingPosition, LandingPosition: LandingPosition, player: p}
}
func (p Player) selectMoveWithString(StartingPosition, LandingPosition string) Move{
	fmt.Println("attempting select")
	startingXInt := convertFirstCharacterToInt(StartingPosition[0])
	startingYInt, _ := strconv.Atoi(string(StartingPosition[1]))
	fmt.Println("attempting select after strconv")
	endingXInt := convertFirstCharacterToInt(LandingPosition[0])
	endingYInt, _ := strconv.Atoi(string(LandingPosition[1]))
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
