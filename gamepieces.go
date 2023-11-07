package main

// import "fmt"

type Condition struct{
	name string
	active bool
}

type MovementType struct{
	X int
	Y int
	Condition
}

type GamePiece struct{
	Name string
	Player
	Distance bool
	MovementTypes []MovementType
	Back bool
	Moved bool
	capturing bool
	Position
}

func (p *GamePiece) addPlayer(position Position){
	if position.Y == 1 || position.Y == 2{
		p.Player.Team = 1
		} else {
			p.Player.Team = 2
		}
}

var MovementTypes = map[string][]MovementType{
	"pawn" : {
		{
			X:0,
			Y:1,
			Condition: Condition{name:"capture", active: false},
		},
		{
			X:0,
			Y:2,
			Condition: Condition{name:"moved", active: false},
		},
		{
			X:1,
			Y:1,
			Condition: Condition{name:"capture", active: true},
		},
	},
	"rook": {
		{
			X:0,
			Y:1,
		},
		{
			X:1,
			Y:0,
		},
	},
	"bishop": {
		{
			X:1,
			Y:1,
		},
	},
	"queen": {
		{
			X:1,
			Y:1,
		},
		{
			X:1,
			Y:1,
		},
		{
			X:0,
			Y:1,
		},
		{
			X:1,
			Y:0,
		},
	},
	"king": {
		{
			X:1,
			Y:1,
		},
		{
			X:1,
			Y:1,
		},
		{
			X:0,
			Y:1,
		},
		{
			X:1,
			Y:0,
		},
	},
	"knight" : {
		{
			X:2,
			Y:1,
		},
		{
			X:1,
			Y:2,
		},
	},
}

var pawn = GamePiece{
	Name: "pawn",
	MovementTypes: []MovementType{
		{
			X:0,
			Y:1,
			Condition: Condition{name:"capture", active: false},
		},
		{
			X:0,
			Y:2,
			Condition: Condition{name:"moved", active: false},
		},
		{
			X:1,
			Y:1,
			Condition: Condition{name:"capture", active: true},
		},
	},
	Distance: false,
	Back: false,
}
var rook = GamePiece{
	Name: "rook",
	MovementTypes: []MovementType{
		{
			X:0,
			Y:1,
		},
		{
			X:1,
			Y:0,
		},
	},
	Distance: true,
	Back: true,
}
var bishop = GamePiece{
	Name: "bishop",
	MovementTypes: []MovementType{
		{
			X:1,
			Y:1,
		},
	},
	Back: true,
	Distance: true,
}
var queen = GamePiece{
	Name: "queen",
	MovementTypes: []MovementType{
		{
			X:1,
			Y:1,
		},
		{
			X:1,
			Y:1,
		},
		{
			X:0,
			Y:1,
		},
		{
			X:1,
			Y:0,
		},
	},
	Distance: true,
	Back: true,
}
var king = GamePiece{
	Name: "king",
	MovementTypes: []MovementType{
		{
			X:1,
			Y:1,
		},
		{
			X:1,
			Y:1,
		},
		{
			X:0,
			Y:1,
		},
		{
			X:1,
			Y:0,
		},
	},
	Distance: false,
	Back: true,
}
var knight = GamePiece{
	Name: "knight",
	MovementTypes: []MovementType{
		{
			X:2,
			Y:1,
		},
		{
			X:1,
			Y:2,
		},
	},
	Distance: false,
	Back: true,
}

