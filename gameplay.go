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
		piece := c.getPiece(move.StartingPosition)
		if c.checkValidLanding(&piece, move.LandingPosition) && c.checkMove(piece, move) && c.checkPath(piece, move) {
			c.acceptMove(move)
			c.moveTurn()
		}
	}
	c.displayBoard()
	if c.winner.Team != 0 {
		fmt.Println("\n", "Winner: ", c.winner.Team)
  } else {
		fmt.Println("\n", c.playerTurn, "'s turn")	
	}
}

func (c ChessPlay) checkPiece(move Move) bool{
	piecePlayer := c.squares[move.StartingPosition].gamePiece.Player
	return c.playerTurn == move.player && piecePlayer.Team == move.player.Team
}

func (c ChessPlay) getPiece(position Position) GamePiece{
	return c.GameBoard.squares[position].gamePiece
}

func (c *ChessPlay) checkValidLanding(piece *GamePiece, LandingPosition Position) bool{
	LandingPositionPiecePlayer := c.squares[LandingPosition].gamePiece.Player
	fmt.Println(LandingPositionPiecePlayer)
	if piece.Player == LandingPositionPiecePlayer {
		fmt.Println(piece.Name, c.squares[LandingPosition].gamePiece.Name, "pieces !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		if (piece.Name != "rook" && c.squares[LandingPosition].gamePiece.Name != "king"){
			fmt.Println("landing on owned piece")
			return false
		}
	} else if piece.Player != LandingPositionPiecePlayer && LandingPositionPiecePlayer.Team != 0 {
		capturedPiece := c.squares[LandingPosition].gamePiece
		if capturedPiece.Name == "king"{
			c.winner = piece.Player
		}
		fmt.Println("\nCaptured ", capturedPiece)
		fmt.Println("")
		piece.capturing = true
	}
	return true
}

func (c ChessPlay) checkMove(piece GamePiece, move Move) (isValidMove bool){
	fmt.Println("CHECKING MOVE")
	spacesMoved := Position{getAbsolute(move.LandingPosition.X - move.StartingPosition.X), getAbsolute(move.LandingPosition.Y - move.StartingPosition.Y)}
	checkCondition := false
	checkMove := MovementType{} 
	if c.checkDirection(piece, move){
		for _, validMove := range MovementTypes[piece.Name] {
			if spacesMoved.X == validMove.X && spacesMoved.Y == validMove.Y{
				if validMove.Condition.name != "" {
					checkCondition = true
					checkMove = validMove
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
	if checkCondition && !c.checkCondition(piece, checkMove){
		isValidMove = false
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

func (c ChessPlay) checkCondition(piece GamePiece, move MovementType) bool{
	if move.Condition.name == "moved"  {
		return move.Condition.active == piece.Moved
	}
	if move.Condition.name == "capture" {
		return piece.capturing == move.Condition.active
	}
	return false
}

func (c ChessPlay) checkPath(piece GamePiece, move Move) bool{
	fmt.Println("CHECKING PATH!!!")
	if piece.Name != "knight" {
		spacesMoved := Position{ move.LandingPosition.X - move.StartingPosition.X, move.LandingPosition.Y - move.StartingPosition.Y }
		currentSquare := move.StartingPosition
		for currentSquare != move.LandingPosition{
			if c.squares[currentSquare].gamePiece.Player == c.playerTurn && currentSquare != move.StartingPosition {
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
	fmt.Println("ACCEPTING !!!!!!!!!!!!!!!!", move)
	if c.checkCastle(move){
		return
	}
	piece := c.squares[move.StartingPosition].gamePiece
	piece.Moved = true
	piece.capturing = false
	c.squares[move.StartingPosition] = Square{}
	newSquare := c.squares[move.LandingPosition] 
	newSquare.gamePiece = piece
	c.squares[move.LandingPosition] = newSquare
}

func (c *ChessPlay) checkCastle(move Move) (bool){
	piece := c.squares[move.StartingPosition].gamePiece
	landingPiece := c.squares[move.LandingPosition].gamePiece
	if (piece.Name == "rook" && landingPiece.Name == "king" && piece.Moved == false && landingPiece.Moved == false){
		if piece.Player.Team == 1 {
			c.acceptMove(Move{move.StartingPosition, Position{X:4, Y:move.StartingPosition.Y}, move.player})
			c.acceptMove(Move{move.LandingPosition, Position{X:3, Y:move.StartingPosition.Y}, move.player})
		} else {
			c.acceptMove(Move{move.StartingPosition, Position{X:6, Y:move.StartingPosition.Y}, move.player})
			c.acceptMove(Move{move.LandingPosition, Position{X:7, Y:move.StartingPosition.Y}, move.player})
		}
		return true
	} 
	return false
} 

func (c *ChessPlay) moveTurn(){
	if c.playerTurn == c.player1 {
		c.playerTurn = c.player2
	} else {
		c.playerTurn = c.player1
	}
}

func getAbsolute(number int) int{
	if number >= 0{
		return number
	}
	return -number
}