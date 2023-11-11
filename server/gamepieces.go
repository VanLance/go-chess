package main

// import "fmt"


type GamePiece struct{
	Name string
	Player
	Distance bool
	Moved bool
	capturing bool
	EnPassant Position
	Position
}

func (p *GamePiece) addPlayer(position Position){
	if position.Y == 1 || position.Y == 2{
		p.Player.Team = 1
		} else {
			p.Player.Team = 2
		}
}

var pawn = GamePiece {
	Name: "pawn",
	Distance: false,
}
var rook = GamePiece{
	Name: "rook",
	Distance: true,
}
var bishop = GamePiece{
	Name: "bishop",
	Distance: true,
}
var queen = GamePiece{
	Name: "queen",
	Distance: true,
}
var king = GamePiece{
	Name: "king",
	Distance: false,
}
var knight = GamePiece{
	Name: "knight",
	Distance: false,
}

