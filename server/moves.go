package main

import "fmt"

type Condition struct {
	name   string
	active bool
}

type MovementType struct {
	X int
	Y int
	conditions []Condition
}

var MovementTypes = map[string][]MovementType{
	"pawn": {
		{
			X:         0,
			Y:         1,
			conditions: []Condition{{name: "capture", active: false}},
		},
		{
			X:         0,
			Y:         2,
			conditions: []Condition{{name: "moved", active: false}},
		},
		{
			X:         1,
			Y:         1,
			conditions: []Condition{{name: "capture", active: true}, {name:"en-passant", active: true}},
		},

	},
	"rook": {
		{
			X: 0,
			Y: 1,
		},
		{
			X: 1,
			Y: 0,
		},
	},
	"bishop": {
		{
			X: 1,
			Y: 1,
		},
	},
	"queen": {
		{
			X: 1,
			Y: 1,
		},
		{
			X: 1,
			Y: 1,
		},
		{
			X: 0,
			Y: 1,
		},
		{
			X: 1,
			Y: 0,
		},
	},
	"king": {
		{
			X: 1,
			Y: 1,
		},
		{
			X: 1,
			Y: 1,
		},
		{
			X: 0,
			Y: 1,
		},
		{
			X: 1,
			Y: 0,
		},
	},
	"knight": {
		{
			X: 2,
			Y: 1,
		},
		{
			X: 1,
			Y: 2,
		},
	},
}

type EnPassant struct {
	target Position
}

func (c *ChessPlay) checkEnPassant(move Move) {
	leftSquare := c.squares[Position{X: move.LandingPosition.X - 1, Y: move.LandingPosition.Y}]
	rightSquare := c.squares[Position{X: move.LandingPosition.X + 1, Y: move.LandingPosition.Y}]
	if leftSquare.gamePiece.Name == "pawn" && leftSquare.gamePiece.Player != move.player {
		piece := c.addEnPassant(&leftSquare.gamePiece, move)
		leftSquare.gamePiece = *piece
	} else if rightSquare.gamePiece.Name == "pawn" && rightSquare.gamePiece.Player != move.player {
		piece := c.addEnPassant(&rightSquare.gamePiece, move)
		rightSquare.gamePiece = *piece
	} 
	
	fmt.Println(leftSquare.gamePiece, "FROM CHEC EN P")
	fmt.Println(rightSquare.gamePiece, "FROM CHEC EN P")
}

func (c *ChessPlay) addEnPassant(piece *GamePiece, move Move) *GamePiece{
	piece.enPassant = move.LandingPosition
	fmt.Println(piece.Position, piece.enPassant ," en YEAH baby")
	return piece
}

func (c *ChessPlay) clearEnPassant() {
	for _, v := range c.squares {
		v.gamePiece.enPassant = Position{X: 0, Y: 0}
	}
}

func (c *ChessPlay) enPassantCapture(target GamePiece) {
	var emptySquare Square
	c.squares[target.Position] = emptySquare
}