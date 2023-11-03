package main

import (
	"fmt"
	"testing"
)


func TestCapture( t *testing.T){
	gameBoard := GameBoard{}
	gameBoard.startingBoard()
	chess := ChessPlay{GameBoard: gameBoard, player1:Player{name:"1"}, player2: Player{"2"} }
	playerOnePiece := chess.squares[Position{x:1,y:1}].gamePiece
	result := chess.checkValidLanding(&playerOnePiece, Position{x:1, y:7})
	if !result {
		t.Errorf("Expect true to own square got %v", result)
	}
	if result {
		chess.acceptMove(Move{Position{x:1, y:7}, Position{x:1, y:7},chess.player1})
		if chess.squares[Position{x:1, y:7}].gamePiece.player.name != "2"{
			t.Errorf("Expect player1 to own square got %v", chess.squares[Position{x:1, y:7}].gamePiece.player)
		}
	}	
	playerTwoPiece := chess.squares[Position{x:8,y:8}].gamePiece
	result = chess.checkValidLanding(&playerTwoPiece, Position{x:8, y:2})
	if !result {
		t.Errorf("Expect true to own square got %v", result)
	}
	if result {
		chess.acceptMove(Move{Position{x:8,y:8}, Position{x:8, y:2} ,chess.player1})
		if chess.squares[Position{x:8, y:2}].gamePiece.player.name != "2"{
			t.Errorf("Expect player1 to own square got %v", chess.squares[Position{x:8, y:2}].gamePiece.player)
		}
	}	
}

// func TestLandingOwnPiece(t *testing.T){
// 	gameBoard := GameBoard{}
// 	gameBoard.startingBoard()
// 	chess := ChessPlay{GameBoard: gameBoard, player1:Player{name:"1"}, player2: Player{"2"}}
// 	playerOnePiece := chess.squares[Position{x:1,y:1}].gamePiece
// 	result := chess.checkValidLanding(playerOnePiece, Position{x:1,y:2})
// 	if result{
// 		t.Errorf("Expect false Landed on my pawn got %v", chess.squares[Position{x:8, y:2}].gamePiece.player)
// 	}
// }

func TestPath( t *testing.T){
	gameBoard := GameBoard{}
	gameBoard.startingBoard()
	chess := ChessPlay{GameBoard: gameBoard, player1:Player{name:"1"}, player2: Player{"2"}, playerTurn: Player{name:"1"} }
	if chess.checkPath(rook, Move{Position{1,1},Position{1,5}, chess.player1}) != false{
		t.Errorf("Expect false pawn blocking rook  got true")
	}
	if chess.checkPath(pawn, Move{Position{1,2},Position{1,3}, chess.player1}) != true{
		t.Errorf("Expect true pawn can move got false")
	}
	move := chess.player1.selectMove(Position{4,2},Position{4,3})
	chess.acceptMove(move)
	if chess.checkPath(bishop, Move{Position{3,1},Position{5,3}, chess.player1}) != true{
		chess.displayBoard()
		t.Errorf("Expect true pawn moved bishop clear got false")
	}
	if chess.checkPath(bishop, Move{Position{3,1},Position{1,3}, chess.player1}) != false{
		t.Errorf("Expect false pawn blocking bishop got true")
	}
}

func TestMovesCondition( t *testing.T){
	gameBoard := GameBoard{}
	gameBoard.startingBoard()
	chess := ChessPlay{GameBoard: gameBoard, player1:Player{name:"1"}, player2: Player{"2"} }
	pawn := chess.getPiece(Position{1,2})
	result := chess.checkCondition(pawn, MovementType{x:0,y:2, condition: Condition{name:"moved", active: false}})
	if !result {
		t.Errorf("Expect true pawn first move, two paces got %v", result)
	}
	pawn.capturing = true
	result = chess.checkCondition(pawn, MovementType{x:1,y:1, condition: Condition{name:"capturing", active: true}})
	if result {
		t.Errorf("Expect false pawn not capturing, got %v", result)
	}
}

func TestValidMoves( t *testing.T){
	gameBoard := GameBoard{}
	gameBoard.startingBoard()
	chess := ChessPlay{GameBoard: gameBoard, player1:Player{name:"1"}, player2: Player{"2"}, playerTurn: Player{name:"1"} }
	if chess.checkMove(bishop, Move{Position{2,3},Position{4,5}, chess.player1}) != true {
		t.Errorf("Expect true valid bishop move got false")
	}
	chess.playerTurn = Player{name:"1"}
	if chess.checkMove(bishop, Move{Position{2,3},Position{3,4}, chess.player1}) != true {
		t.Errorf("Expect true valid bishop move got false")
	}
	chess.playerTurn = Player{name:"1"}
	if chess.checkMove(bishop, Move{Position{2,3},Position{2,2}, chess.player1}) != false {
		t.Errorf("Expect false valid bishop move got false")
	}
	chess.playerTurn = Player{name:"1"}
	fmt.Println(chess.playerTurn)
	if chess.checkMove(pawn, Move{Position{2,2},Position{2,4}, chess.player1}) != true {
		t.Errorf("Expect true valid pawn move got false")
	}
	chess.playerTurn = Player{name:"1"}
	if chess.checkMove(pawn, Move{Position{2,2},Position{3,4}, chess.player1}) != false {
		t.Errorf("Expect false invalid pawn move got true")
	}
	chess.playerTurn = Player{name:"1"}
	if chess.checkMove(pawn, Move{Position{2,3},Position{2,2}, chess.player1}) != false {
		t.Errorf("Expect false invalid pawn move got true")
	}
	chess.playerTurn = Player{name:"1"}
	if chess.checkMove(rook, Move{Position{2,3},Position{2,6}, chess.player1}) != true {
		t.Errorf("Expect true invalid rook move got false")
	}
	chess.playerTurn = Player{name:"1"}
	if chess.checkMove(rook, Move{Position{2,3},Position{6,3}, chess.player1}) != true {
		t.Errorf("Expect true invalid rook move got false")
	}
	chess.playerTurn = Player{name:"1"}
	if chess.checkMove(rook, Move{Position{2,5},Position{2,3}, chess.player1}) != true {
		t.Errorf("Expect true valid rook move got false")
	}
	chess.playerTurn = Player{name:"1"}
	if chess.checkMove(rook, Move{Position{2,3},Position{3,4}, chess.player1}) != false {
		t.Errorf("Expect false invalid rook move got true")
	}
	chess.playerTurn = Player{name:"1"}
	if chess.checkMove(knight, Move{Position{2,1},Position{3,3}, chess.player1}) != true {
		t.Errorf("Expect true invalid knight move got false")
	}
	chess.playerTurn = Player{name:"1"}
	if chess.checkMove(knight, Move{Position{2,1},Position{3,4}, chess.player1}) != false {
		t.Errorf("Expect false invalid knight move got true")
	}
}


func playerTurn(chess *ChessPlay, player Player, startingPosition, endingPosition Position){
	move := player.selectMove(startingPosition,endingPosition)
	if chess.checkPiece(move){
		piece := chess.getPiece(move.startingPosition)
		if chess.checkValidLanding(&piece, move.endingPosition) && chess.checkMove(piece, move) && chess.checkPath(piece, move) {
			chess.acceptMove(move)
			chess.moveTurn()
			chess.displayBoard()
		}
	}
}

func TestTurnChange( t *testing.T ){
	gameBoard := GameBoard{}
	gameBoard.startingBoard()
	player1 := Player{name:"1"}
	player2 := Player{name:"2"}
	chess := ChessPlay{GameBoard: gameBoard, player1: player1, player2: player2, playerTurn: player1 }
	playerTurn(&chess, player1, Position{1,2},Position{1,4} )
	if chess.playerTurn != player2 {
		t.Errorf("Expect Player2 turn got %v", chess.playerTurn)
	}
	playerTurn(&chess, player2, Position{2,7},Position{2,5} )
	if chess.playerTurn != player1 {
		t.Errorf("Expect Player1 turn got %v", chess.playerTurn)
	}
	playerTurn(&chess, player2, Position{2,7},Position{2,5} )
	if chess.playerTurn != player1 {
		t.Errorf("Expect Player1 player2 shouldn't be able to move %v", chess.playerTurn)
	}
}
