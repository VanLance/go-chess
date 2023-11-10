package main

import (
	"fmt"
	// "strconv"
)

type ChessPlay struct {
	GameBoard
	player1 Player
	player2 Player
	playerTurn Player
	Winner Player
}

type Move struct{
	StartingPosition Position
	LandingPosition Position
	player Player
}

func createChess() ChessPlay{
	gameBoard := GameBoard{}
	gameBoard.startingBoard()
	gameBoard.displayBoard()
	chess :=  ChessPlay{GameBoard: gameBoard, player1:Player{Team:1}, player2: Player{Team:2} }
	chess.playerTurn = chess.player1
	return chess
}

func (c *ChessPlay) driver(){
	move := c.player1.selectMove(Position{1,2},Position{1,4})
	c.makeMove(move)
	move = c.player2.selectMove(Position{2,7},Position{2,5})
	c.makeMove(move)
	move = c.player1.selectMove(Position{1,4},Position{1,6})
	c.makeMove(move)
	move = c.player1.selectMove(Position{1,4},Position{2,5})
	c.makeMove(move)
	move = c.player2.selectMove(Position{2,5},Position{2,4})
	c.makeMove(move)
	move = c.player2.selectMoveWithString("57","55")
	c.makeMove(move)
	c.makeMove(c.player1.selectMoveWithString("52","54"))
	c.makeMove(c.player2.selectMoveWithString("55","54"))
}

func (c *ChessPlay) makeMove(move Move){
	if  c.checkPiece(move){
		piece := c.getPiece(move.StartingPosition)
		if c.checkValidLanding(&piece, move.LandingPosition) && c.checkMove(piece, move) && c.checkPath(piece, move) {
			c.acceptMove(move)
			c.moveTurn()
		}
	}
	c.displayBoard()
	if c.Winner.Team != 0 {
		fmt.Println("\n", "Winner: ", c.Winner.Team)
  } else {
		fmt.Println("\n", c.playerTurn, "'s turn")	
	}
}

func (c ChessPlay) checkPiece(move Move) bool{
	piecePlayer := c.squares[move.StartingPosition].Player
	return c.playerTurn == move.player && piecePlayer.Team == move.player.Team
}

func (c ChessPlay) getPiece(position Position) GamePiece{
	return c.GameBoard.squares[position]
}

func (c ChessPlay) checkValidLanding(piece *GamePiece, LandingPosition Position) bool{
	LandingPositionPiecePlayer := c.squares[LandingPosition].Player
	if piece.Player == LandingPositionPiecePlayer {
		if (piece.Name != "rook" && c.squares[LandingPosition].Name != "king"){
			fmt.Println("landing on owned piece")
			return false
		}
	} else if piece.Player != LandingPositionPiecePlayer && LandingPositionPiecePlayer.Team != 0 {
		capturedPiece := c.squares[LandingPosition]
		if capturedPiece.Name == "king"{
			c.Winner = piece.Player
		}
		fmt.Println("\nCaptured ", capturedPiece)
		fmt.Println("")
		piece.capturing = true
	}
	return true
}

func (c ChessPlay) checkMove(piece GamePiece, move Move) (isValidMove bool){
	spacesMoved := getSpacesMoved(move)
	if c.checkDirection(piece, move){
		for _, validMove := range MovementTypes[piece.Name] {
			if spacesMoved.X == validMove.X && spacesMoved.Y == validMove.Y{
				if len(validMove.conditions) != 0 {
					if !c.checkCondition(piece, validMove, move.LandingPosition){
						isValidMove = false
						break
					}
				}
				isValidMove = true
			} else if ( spacesMoved.X != 0 ) && ( spacesMoved.Y != 0){
				if ( validMove.X != 0 ) && ( validMove.Y != 0) {
					if spacesMoved.X / validMove.Y == spacesMoved.Y / validMove.Y && piece.Distance == true{
						isValidMove = true
					}
				}
			} else if spacesMoved.X != 0 {
				if validMove.X == 1 && validMove.Y == 0 && piece.Distance == true{
					isValidMove = true
				}
			} else {
				if validMove.X == 1 && validMove.Y == 0 && piece.Distance == true {
					isValidMove = true
				}
			} 
		}
	}
	piece.capturing = false
	return isValidMove
}

func (c ChessPlay) checkDirection(piece GamePiece, move Move) bool{
	if piece.Name == "pawn" {
		if c.playerTurn == c.player1{
			return move.StartingPosition.Y < move.LandingPosition.Y
		} else {
			return move.StartingPosition.Y > move.LandingPosition.Y
		}
	}
	return true
}

func (c ChessPlay) checkCondition(piece GamePiece, moveType MovementType, landingPosition Position) (output bool){
	for _, condition := range moveType.conditions {
		if condition.name == "en-passant"{
			if landingPosition.X == piece.EnPassant.X{
				c.enPassantCapture(piece.EnPassant)
				return true
			}
		}
		if condition.name == "moved"  && piece.Moved == false {
			output = true
		}
		if condition.name == "capture" {
			output =  piece.capturing == condition.active
		}
	}
	return output
}

func (c ChessPlay) checkPath(piece GamePiece, move Move) bool{
	if piece.Name != "knight" {
		spacesMoved := Position{ move.LandingPosition.X - move.StartingPosition.X, move.LandingPosition.Y - move.StartingPosition.Y }
		currentSquare := move.StartingPosition
		for currentSquare != move.LandingPosition{
			if c.squares[currentSquare].Player == c.playerTurn && currentSquare != move.StartingPosition {
				return false
			}
			if ( spacesMoved.X > 0 ) {
				currentSquare.X++
			} else if spacesMoved.X < 0 {
				currentSquare.X--
			}
			if ( spacesMoved.Y > 0){
				currentSquare.Y++
			} else if spacesMoved.Y < 0 {
				currentSquare.Y--
			}
		}
	}
			
	return true
}

func (c *ChessPlay) acceptMove(move Move){
	if c.checkCastle(move){
		c.castleKing(move)
		return
	}
	piece := c.squares[move.StartingPosition]
	piece.Moved = true
	piece.capturing = false
	c.squares[move.StartingPosition] = GamePiece{}
	newSquare := c.squares[move.LandingPosition] 
	newSquare = piece
	c.squares[move.LandingPosition] = newSquare
	c.clearEnPassant()
	c.checkEnPassant(move)
}


func (c *ChessPlay) moveTurn(){
	if c.playerTurn == c.player1 {
		c.playerTurn = c.player2
	} else {
		c.playerTurn = c.player1
	}
}


// func (c ChessPlay) isCheck(){
// 	var opposingTeam Player
// 	if c.playerTurn.Team == 1{
// 		opposingTeam = c.player2
// 	} else {
// 		opposingTeam = c.player1
// 	}
// 	for _, square := range c.squares{
// 		if piece.Player == opposingTeam{

// 		}
// 	}
// }
