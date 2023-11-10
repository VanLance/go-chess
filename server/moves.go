package main



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

func (c ChessPlay) getAdjacentSquares(move Move) (Square, Square){
	leftSquare := c.squares[Position{X: move.LandingPosition.X - 1, Y: move.LandingPosition.Y}]
	rightSquare := c.squares[Position{X: move.LandingPosition.X + 1, Y: move.LandingPosition.Y}]
	return leftSquare, rightSquare
}

func (c *ChessPlay) checkEnPassant(move Move) {
	p := Position{X:0,Y:2}
	if p == getSpacesMoved(move){
		leftSquare, rightSquare := c.getAdjacentSquares(move) 
		if leftSquare.gamePiece.Name == "pawn" && leftSquare.gamePiece.Player != move.player {
			piece := c.addEnPassant(leftSquare.gamePiece, move)
			leftSquare.addGamePiece(piece)
			c.GameBoard.squares[Position{X: move.LandingPosition.X - 1, Y: move.LandingPosition.Y}] = leftSquare
		} else if rightSquare.gamePiece.Name == "pawn" && rightSquare.gamePiece.Player != move.player {
			piece := c.addEnPassant(rightSquare.gamePiece, move)
			rightSquare.addGamePiece(piece)
			c.GameBoard.squares[Position{X: move.LandingPosition.X + 1, Y: move.LandingPosition.Y}] = rightSquare
		}
	}
}

func (c ChessPlay) addEnPassant(piece GamePiece, move Move) GamePiece{
	piece.EnPassant = move.LandingPosition
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

func (c ChessPlay) checkCastle(move Move) (bool){
	piece := c.squares[move.StartingPosition].gamePiece
	landingPiece := c.squares[move.LandingPosition].gamePiece
	if (piece.Name == "rook" && landingPiece.Name == "king" && piece.Moved == false && landingPiece.Moved == false){
		return true
	} 
	return false
} 

func (c ChessPlay) castleKing(move Move){
	var rookXLanding, kingXLanding, yLanding int
	if move.StartingPosition.X == 1 {
			rookXLanding = 4
			kingXLanding = 3
		} else {
			rookXLanding = 6
			kingXLanding = 7
	}
	if move.player.Team == 1 {
		yLanding = 1
	} else {
		yLanding = 8
	}
	c.acceptMove(Move{move.StartingPosition, Position{X:rookXLanding, Y:yLanding}, move.player})
	c.acceptMove(Move{move.LandingPosition, Position{X:kingXLanding, Y:yLanding}, move.player})
}