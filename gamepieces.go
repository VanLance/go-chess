package main

// import "fmt"

type Condition struct{
	conditionName string
	moved bool
	capture string
}

type MovementType struct{
	x int
	y int
	condition string
}

type GamePiece struct{
	name string
	player Player
	distance bool
	movementTypes []MovementType
	back bool
	moved bool
	capturing bool
}

func (p *GamePiece) addPlayer(position Position){
	if position.y == 1 || position.y == 2{
		p.player.name = "1"
		} else {
			p.player.name = "2"
		}
}

var pawn = GamePiece{
	name: "pawn",
	movementTypes: []MovementType{
		{
			x:0,
			y:1,
		},
		{
			x:0,
			y:2,
			condition: "moved",
		},
		{
			x:1,
			y:1,
			condition: "capture",
		},
	},
	distance: false,
	back: false,
}
var rook = GamePiece{
	name: "rook",
	movementTypes: []MovementType{
		{
			x:0,
			y:1,
		},
		{
			x:1,
			y:0,
		},
	},
	distance: true,
	back: true,
}
var bishop = GamePiece{
	name: "bishop",
	movementTypes: []MovementType{
		{
			x:1,
			y:1,
		},
	},
	back: true,
	distance: true,
}
var queen = GamePiece{
	name: "queen",
	movementTypes: []MovementType{
		{
			x:1,
			y:1,
		},
		{
			x:1,
			y:1,
		},
		{
			x:0,
			y:1,
		},
		{
			x:1,
			y:0,
		},
	},
	distance: true,
	back: true,
}
var king = GamePiece{
	name: "king",
	movementTypes: []MovementType{
		{
			x:1,
			y:1,
		},
		{
			x:1,
			y:1,
		},
		{
			x:0,
			y:1,
		},
		{
			x:1,
			y:0,
		},
	},
	distance: false,
	back: true,
}
var knight = GamePiece{
	name: "knight",
	movementTypes: []MovementType{
		{
			x:2,
			y:1,
		},
		{
			x:1,
			y:2,
		},
	},
	distance: false,
	back: true,
}

// func createPieces()(pieces []GamePiece){
// 	var pawn = GamePiece{
// 		name: "pawn",
// 		captured: false,
// 		Movement: Movement{},
// 	}
// 	var rook = GamePiece{
// 		name: "rook",
// 		captured: false,
// 		Movement: Movement{},
// 	}
// 	var bishop = GamePiece{
// 		name: "bishop",
// 		captured: false,
// 		Movement: Movement{},
// 	}
// 	var queen = GamePiece{
// 		name: "queen",
// 		captured: false,
// 		Movement: Movement{},
// 	}
// 	var king = GamePiece{
// 		name: "king",
// 		captured: false,
// 		Movement: Movement{},
// 	}
// 	pieces = append(pieces, rook, bishop, queen, king, rook, bishop)
// 	for num := 0; num < 8; num++ {
// 		pieces = append(pieces,pawn)
// 	}
// 	return pieces
// }