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

// func (c ChessPlay) getAdjacentSquares(move Move) (Square, Square){
// 	leftSquare := c.squares[Position{X: move.LandingPosition.X - 1, Y: move.LandingPosition.Y}]
// 	rightSquare := c.squares[Position{X: move.LandingPosition.X + 1, Y: move.LandingPosition.Y}]
// 	return leftSquare, rightSquare
// }

func (c *ChessPlay) checkEnPassant(move Move) {
	p := Position{X:0,Y:2}
	if p == getSpacesMoved(move){
		leftSquare := c.squares[Position{X: move.LandingPosition.X - 1, Y: move.LandingPosition.Y}]
		rightSquare := c.squares[Position{X: move.LandingPosition.X + 1, Y: move.LandingPosition.Y}]
		if leftSquare.gamePiece.Name == "pawn" && leftSquare.gamePiece.Player != move.player {
			piece := c.addEnPassant(&leftSquare.gamePiece, move)
			fmt.Println(*piece, "assinging to square left")
			leftSquare.addGamePiece(*piece)
			c.GameBoard.squares[Position{X: move.LandingPosition.X - 1, Y: move.LandingPosition.Y}] = leftSquare
			fmt.Println(c.GameBoard.squares[Position{X: move.LandingPosition.X - 1, Y: move.LandingPosition.Y}], "UPDATING SQARE left ")
		} else if rightSquare.gamePiece.Name == "pawn" && rightSquare.gamePiece.Player != move.player {
			piece := c.addEnPassant(&rightSquare.gamePiece, move)
			fmt.Println(*piece, "assinging to square right")
			rightSquare.addGamePiece(*piece)
			c.GameBoard.squares[Position{X: move.LandingPosition.X + 1, Y: move.LandingPosition.Y}] = rightSquare
			fmt.Println(c.GameBoard.squares[Position{X: move.LandingPosition.X + 1, Y: move.LandingPosition.Y}], "UPDATING SQARE right")
		}
	}
}

func (c *ChessPlay) addEnPassant(piece *GamePiece, move Move) *GamePiece{
	piece.EnPassant = move.LandingPosition
	fmt.Println(piece.Position, piece.EnPassant ," en YEAH baby")
	return piece
}

func (c *ChessPlay) clearEnPassant() {
	for position, v := range c.squares {
		v.gamePiece.EnPassant = Position{X: 0, Y: 0}
		c.GameBoard.squares[position] = v
	}
}

func (c *ChessPlay) enPassantCapture(position Position) {
	var emptySquare Square
	c.squares[position] = emptySquare
	oldSquare := c.GameBoard.squares[position]
	oldSquare.addGamePiece(GamePiece{})
}