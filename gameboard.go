package main

import (
	"fmt"
	"strconv"
)

type Position struct{
	X int
	Y int
}

type Square struct{
	gamePiece GamePiece
}

type GameBoard struct{ 
	squares map[Position]Square
}



func (s *Square) addGamePiece(gamepiece GamePiece){
	s.gamePiece = gamepiece
}

func (g *GameBoard) addSquares() {
	g.squares= make(map[Position]Square)
	for numX := 1; numX <= 8; numX++ {
		for numY := 1; numY <= 8; numY++ {
			g.squares[Position{X:numX, Y:numY}] = Square{}
		}
	}
}

func (g *GameBoard) addPawns() {
	for numX := 1; numX <= 8; numX++ {
		square1 := g.squares[Position{X:numX, Y:2}]
		square2 := g.squares[Position{X:numX, Y:7}]
		pawn1 := pawn
		pawn1.addPlayer(Position{X:numX, Y:2})
		pawn2 := pawn
		pawn2.addPlayer(Position{X:numX, Y:7})
		square1.addGamePiece(pawn1)
		square2.addGamePiece(pawn2)
		
		g.squares[Position{X:numX, Y:2}] = square1
		g.squares[Position{X:numX, Y:7}] = square2
	}
}


func (g *GameBoard) addPieces(piecePositions []Position, piece GamePiece) {
	for _, position := range piecePositions {
		piece.addPlayer(position)
			square := g.squares[position]
			square.gamePiece = piece
			g.squares[position] = square
		}
	}

func (g *GameBoard) addRooks() {
	rookPostions := []Position{
		{X:1,Y:1},
		{X:8,Y:1},
		{X:8,Y:1}, 
		{X:8,Y:1}, 
		{X:1,Y:8},
		{X:8,Y:8}}
	g.addPieces(rookPostions, rook)
	
}

func (g *GameBoard) addKnights() {
	knightPostions := []Position{{X:2 ,Y: 1},{X:7 ,Y:1},{X:2 ,Y: 8},{X:7 ,Y:8 }}
	g.addPieces(knightPostions, knight)
	
}

func (g *GameBoard) addBishops() {
	bishopPositions := []Position{{X:3,Y:1},{X:6, Y:1 },{X:3,Y:8},{X:6 ,Y:8 }}
	g.addPieces(bishopPositions,bishop )
	
}

func (g *GameBoard) addKings() {
	kingPositions := []Position{{X:5,Y:1},{X:5,Y:8}}
	g.addPieces(kingPositions, king)
	
}

func (g *GameBoard) addQueens() {
	queenPositions := []Position{{X:4,Y:1},{X:4,Y:8}}
	g.addPieces(queenPositions, queen)
	
}


func (g *GameBoard) startingBoard() {
	g.addSquares()
	g.addPawns()
	g.addRooks()
	g.addKnights()
	g.addBishops()
	g.addKings()
	g.addQueens()
}

func (g GameBoard) displayBoard(){
	letters := []string{"  |    A     |","   B     |","   C     |","   D     |","   E     |","   F     |","   G     |","   H     |  "}
	fmt.Println(letters)
	fmt.Println("   _______________________________________________________________________________________")
	for numberY := 8; numberY > 0; numberY-- {
		var row []string
		row = append(row, strconv.Itoa(numberY),"|")
		for numberX:= 1; numberX <= 8; numberX++{
			piece := g.squares[Position{X:numberX,Y: numberY}].gamePiece
			pieceName := evenCells(piece.Name)
			row = append(row, pieceName,piece.Player.Username, "|")
		}
		row = append(row, strconv.Itoa(numberY))
		fmt.Println(row)
		fmt.Println("   _______________________________________________________________________________________")
	}
	fmt.Println(letters)
}

func evenCells(word string) string{
	remainingLength := 6 - len(word)
	if word == ""{
		remainingLength ++
	}
	for i:=0; i<remainingLength; i++{
		word += " "
	}
	return word
}