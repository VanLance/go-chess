package main

import (
	"fmt"
	"testing"
)


func TestCapture( t *testing.T){
	chess := createChess()
	playerOnePiece := chess.squares[Position{X:1,Y:1}]
	result := chess.checkValidLanding(&playerOnePiece, Position{X:1, Y:7})
	if !result {
		t.Errorf("Expect true to own square got %v", result)
	}
	if result {
		chess.acceptMove(Move{Position{X:1, Y:7}, Position{X:1, Y:7}, &chess.player1})
		if chess.squares[Position{X:1, Y:7}].Player.Team != 2{
			t.Errorf("Expect player1 to own square got %v", chess.squares[Position{X:1, Y:7}].Player)
		}
	}	
	playerTwoPiece := chess.squares[Position{X:8,Y:8}]
	result = chess.checkValidLanding(&playerTwoPiece, Position{X:8, Y:2})
	if !result {
		t.Errorf("Expect true to own square got %v", result)
	}
	if result {
		chess.acceptMove(Move{Position{X:8,Y:8}, Position{X:8, Y:2}, &chess.player1})
		if chess.squares[Position{X:8, Y:2}].Player.Team != 2{
			t.Errorf("Expect player1 to own square got %v", chess.squares[Position{X:8, Y:2}].Player)
		}
	}	
}

// func TestLandingOwnPiece(t *testing.T){
// 	gameBoard := GameBoard{}
// 	gameBoard.startingBoard()
// 	chess := ChessPlay{GameBoard: gameBoard, player1:Player{name:"1"}, player2: Player{"2"}}
// 	playerOnePiece := chess.squares[Position{X:1,Y:1}]
// 	result := chess.checkValidLanding(playerOnePiece, Position{X:1,Y:2})
// 	if result{
// 		t.Errorf("Expect false Landed on my pawn got %v", chess.squares[Position{X:8, Y:2}].player)
// 	}
// }

func TestPath( t *testing.T){
	
	chess := createChess()
	if chess.checkPath(rook, Move{Position{1,1},Position{1,5}, &chess.player1}) != false{
		t.Errorf("Expect false pawn blocking rook  got true")
	}
	if chess.checkPath(pawn, Move{Position{1,2},Position{1,3}, &chess.player1}) != true{
		t.Errorf("Expect true pawn can move got false")
	}
	move := chess.player1.selectMove(Position{4,2},Position{4,3})
	chess.acceptMove(move)
	if chess.checkPath(bishop, Move{Position{3,1},Position{5,3}, &chess.player1}) != true{
		chess.displayBoard()
		t.Errorf("Expect true pawn moved bishop clear got false")
	}
	if chess.checkPath(bishop, Move{Position{3,1},Position{1,3}, &chess.player1}) != false{
		t.Errorf("Expect false pawn blocking bishop got true")
	}
}

func TestMovesCondition( t *testing.T){
	chess := createChess()
	pawn := chess.getPiece(Position{1,2})
	result := chess.checkCondition(pawn, MovementType{X:0,Y:2, conditions: []Condition{{name:"moved", active: false}}}, Position{X:0,Y:4})
	if !result {
		t.Errorf("Expect true pawn first move, two paces got %v", result)
	}
	pawn.capturing = true
	result = chess.checkCondition(pawn, MovementType{X:1,Y:1, conditions: []Condition{{name:"capturing", active: true}}}, Position{X:2,Y:3})
	if result {
		t.Errorf("Expect false pawn not capturing, got %v", result)
	}
}

func TestValidMoves( t *testing.T){
	chess := createChess()
	if chess.checkMove(bishop, Move{Position{2,3},Position{4,5}, &chess.player1}) != true {
		t.Errorf("Expect true valid bishop move got false")
	}
	chess.playerTurn = &chess.player1
	if chess.checkMove(bishop, Move{Position{2,3},Position{3,4}, &chess.player1}) != true {
		t.Errorf("Expect true valid bishop move got false")
	}
	chess.playerTurn = &chess.player1
	if chess.checkMove(bishop, Move{Position{2,3},Position{2,2}, &chess.player1}) != false {
		t.Errorf("Expect false valid bishop move got false")
	}
	chess.playerTurn = &chess.player1
	fmt.Println(chess.playerTurn)
	if chess.checkMove(pawn, Move{Position{2,2},Position{2,4}, &chess.player1}) != true {
		t.Errorf("Expect true valid pawn move got false")
	}
	chess.playerTurn = &chess.player1
	if chess.checkMove(pawn, Move{Position{2,2},Position{3,4}, &chess.player1}) != false {
		t.Errorf("Expect false invalid pawn move got true")
	}
	chess.playerTurn = &chess.player1
	if chess.checkMove(pawn, Move{Position{2,3},Position{2,2}, &chess.player1}) != false {
		t.Errorf("Expect false invalid pawn move got true")
	}
	chess.playerTurn = &chess.player1
	if chess.checkMove(rook, Move{Position{2,3},Position{2,6}, &chess.player1}) != true {
		t.Errorf("Expect true invalid rook move got false")
	}
	chess.playerTurn = &chess.player1
	if chess.checkMove(rook, Move{Position{2,3},Position{6,3}, &chess.player1}) != true {
		t.Errorf("Expect true invalid rook move got false")
	}
	chess.playerTurn = &chess.player1
	if chess.checkMove(rook, Move{Position{2,5},Position{2,3}, &chess.player1}) != true {
		t.Errorf("Expect true valid rook move got false")
	}
	chess.playerTurn = &chess.player1
	if chess.checkMove(rook, Move{Position{2,3},Position{3,4}, &chess.player1}) != false {
		t.Errorf("Expect false invalid rook move got true")
	}
	chess.playerTurn = &chess.player1
	if chess.checkMove(knight, Move{Position{2,1},Position{3,3}, &chess.player1}) != true {
		t.Errorf("Expect true invalid knight move got false")
	}
	chess.playerTurn = &chess.player1
	if chess.checkMove(knight, Move{Position{2,1},Position{3,4}, &chess.player1}) != false {
		t.Errorf("Expect false invalid knight move got true")
	}
}


func playerTurn(chess *ChessPlay, player Player, StartingPosition, LandingPosition Position){
	move := player.selectMove(StartingPosition,LandingPosition)
	if chess.checkPiece(move){
		piece := chess.getPiece(move.StartingPosition)
		if chess.checkValidLanding(&piece, move.LandingPosition) && chess.checkMove(piece, move) && chess.checkPath(piece, move) {
			chess.acceptMove(move)
			chess.moveTurn()
			chess.displayBoard()
		}
	}
}

func TestTurnChange( t *testing.T ){
	chess := createChess()
	playerTurn(&chess, chess.player1, Position{1,2},Position{1,4} )
	if chess.playerTurn != &chess.player2 {
		t.Errorf("Expect Player2 turn got %v", chess.playerTurn)
	}
	playerTurn(&chess, chess.player2, Position{2,7},Position{2,5} )
	if chess.playerTurn != &chess.player1 {
		t.Errorf("Expect Player1 turn got %v", chess.playerTurn)
	}
	playerTurn(&chess, chess.player2, Position{2,7},Position{2,5} )
	if chess.playerTurn != &chess.player1 {
		t.Errorf("Expect Player1 player2 shouldn't be able to move %v", chess.playerTurn)
	}
}
