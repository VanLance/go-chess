package main

import "fmt"

type Condition struct {
	name   string
	active bool
}

type MovementType struct {
	X int
	Y int
	Condition
}

var MovementTypes = map[string][]MovementType{
	"pawn": {
		{
			X:         0,
			Y:         1,
			Condition: Condition{name: "capture", active: false},
		},
		{
			X:         0,
			Y:         2,
			Condition: Condition{name: "moved", active: false},
		},
		{
			X:         1,
			Y:         1,
			Condition: Condition{name: "capture", active: true},
		},
		// {
		// 	X:1,
		// 	Y:1,
		// 	Condition: Condition{name:"en-passant", active: true},
		// },
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
		c.addEnPassant(&leftSquare.gamePiece, c.squares[move.StartingPosition].gamePiece)
	} else if rightSquare.gamePiece.Name == "pawn" && rightSquare.gamePiece.Player != move.player {
		c.addEnPassant(&rightSquare.gamePiece, c.squares[move.StartingPosition].gamePiece)
	}
}

func (c *ChessPlay) addEnPassant(piece *GamePiece, targetPiece GamePiece) {
	piece.enPassant = targetPiece.Position
	fmt.Println(piece ," en YEAH baby")
}

func (c *ChessPlay) clearEnPassant() {
	for _, v := range c.squares {
		v.gamePiece.enPassant = Position{X: 0, Y: 0}
	}
}

func (c *ChessPlay) enPassantCapture(target GamePiece) {
	captureSquare := c.squares[target.Position]
	for _, v := range c.squares {
		if v.gamePiece.Name == "" {
			captureSquare.gamePiece = v.gamePiece
		}
	}
}