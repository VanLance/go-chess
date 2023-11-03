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
	winner Player
}

type Move struct{
	startingPosition Position
	endingPosition Position
	player Player
}

func (c ChessPlay) driver(){
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
	c.makeMove(c.player2.selectMoveWithString("48","84"))
	c.makeMove(c.player1.selectMoveWithString("54","55"))
	c.makeMove(c.player1.selectMoveWithString("72","83"))
	c.makeMove(c.player1.selectMoveWithString("72","73"))
	c.makeMove(c.player2.selectMoveWithString("84","83"))
	c.makeMove(c.player1.selectMoveWithString("71","83"))
	c.makeMove(c.player2.selectMoveWithString("H7","H5"))
	c.makeMove(c.player1.selectMoveWithString("D1","D3"))
	c.makeMove(c.player1.selectMoveWithString("D1","H5"))
	c.makeMove(c.player2.selectMoveWithString("b8","a6"))
	c.makeMove(c.player1.selectMoveWithString("h5","f7"))
	c.makeMove(c.player2.selectMoveWithString("e8","e7"))
	c.makeMove(c.player1.selectMoveWithString("f7","e7"))
}

func (c *ChessPlay) makeMove(move Move){
	if  c.checkPiece(move){
		piece := c.getPiece(move.startingPosition)
		if c.checkValidLanding(&piece, move.endingPosition) && c.checkMove(piece, move) && c.checkPath(piece, move) {
			c.acceptMove(move)
			c.moveTurn()
		}
	}
	c.displayBoard()
	if c.winner.name != "" {
		fmt.Println("\n", "Winner: ", c.winner.name)
  } else {
		fmt.Println("\n", c.playerTurn, "'s turn")	
	}
}

func (c ChessPlay) checkPiece(move Move) bool{
	piecePlayer := c.squares[move.startingPosition].gamePiece.player
	return c.playerTurn == move.player && piecePlayer.name == move.player.name
}

func (c ChessPlay) getPiece(position Position) GamePiece{
	return c.GameBoard.squares[position].gamePiece
}

func (c ChessPlay) checkDirection(piece GamePiece, move Move) bool{
	if piece.back == false {
		if c.playerTurn == c.player1{
			return move.startingPosition.y < move.endingPosition.y
		} else {
			return move.startingPosition.y > move.endingPosition.y
		}
	}
	return true
}

func (c ChessPlay) checkMove(piece GamePiece, move Move) (isValidMove bool){
	spacesMoved := Position{getAbsolute(move.endingPosition.x- move.startingPosition.x), getAbsolute(move.endingPosition.y - move.startingPosition.y)}
	checkCondition := false
	checkMove := MovementType{} 
	if c.checkDirection(piece, move){
		for _, validMove := range piece.movementTypes {
			if spacesMoved.x == validMove.x && spacesMoved.y == validMove.y{
				if validMove.condition.name != "" {
					checkCondition = true
					checkMove = validMove
				}
				isValidMove = true
			}
			if ( spacesMoved.x != 0 ) && ( spacesMoved.y != 0){
				if ( validMove.x != 0 ) && ( validMove.y != 0) {
					if spacesMoved.x / validMove.y == spacesMoved.y / validMove.y && piece.distance == true{
						isValidMove = true
					}
				}
			} else if spacesMoved.x != 0 {
				if validMove.x == 1 && validMove.y == 0 && piece.distance == true{
					isValidMove = true
				}
			} else {
				if validMove.x == 1 && validMove.y == 0 && piece.distance == true {
					isValidMove = true
				}
			} 
		}
	}
	
	if checkCondition && !c.checkCondition(piece, checkMove){
		isValidMove = false
	}
	piece.capturing = false
	return isValidMove
}

func (c ChessPlay) checkCondition(piece GamePiece, move MovementType) bool{
	if move.condition.name == "moved"  {
		return move.condition.active == piece.moved
	}
	if move.condition.name == "capture" {
		return piece.capturing == move.condition.active
	}
	return false
}

func (c ChessPlay) checkPath(piece GamePiece, move Move) bool{
	if piece.name != "knight" {
		spacesMoved := Position{ move.endingPosition.x - move.startingPosition.x, move.endingPosition.y - move.startingPosition.y }
		currentSquare := move.startingPosition
		for currentSquare != move.endingPosition{
			if ( spacesMoved.x > 0 ) {
				currentSquare.x++
			} else if spacesMoved.x < 0 {
				currentSquare.x--
			}
			if ( spacesMoved.y > 0){
				currentSquare.y++
			} else if spacesMoved.y < 0 {
				currentSquare.y--
			}
			if c.squares[currentSquare].gamePiece.player == c.playerTurn {
				return false
			}
		}
	}
			
	return true
}

func (c *ChessPlay) checkValidLanding(piece *GamePiece, landingPosition Position) bool{
	landingPositionPiecePlayer := c.squares[landingPosition].gamePiece.player
	if piece.player == landingPositionPiecePlayer {
		fmt.Println("landing on owned piece")
		return false
	} else if piece.player != landingPositionPiecePlayer && landingPositionPiecePlayer.name != "" {
		capturedPiece := c.squares[landingPosition].gamePiece
		if capturedPiece.name == "king"{
			c.winner = piece.player
		}
		fmt.Println("\nCaptured ", capturedPiece)
		fmt.Println("")
		piece.capturing = true
	}
	return true
}

func (c *ChessPlay) acceptMove(move Move){
	piece := c.squares[move.startingPosition].gamePiece
	piece.moved = true
	c.squares[move.startingPosition] = Square{}
	newSquare := c.squares[move.endingPosition] 
	newSquare.gamePiece = piece
	c.squares[move.endingPosition] = newSquare
}

func (c *ChessPlay) moveTurn(){
	if c.playerTurn == c.player1 {
		c.playerTurn = c.player2
	} else {
		c.playerTurn = c.player1
	}
}

func (c ChessPlay) checkLanding(){}

func getAbsolute(number int) int{
	if number >= 0{
		return number
	}
	return -number
}