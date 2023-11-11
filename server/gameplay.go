package main

import (
	"fmt"
	// "strconv"
)

type ChessPlay struct {
	GameBoard
	playerTurn *Player
	Winner Player
	Message string
}

type Move struct{
	StartingPosition Position
	LandingPosition Position
	player *Player
}

func createChess() ChessPlay{
	gameBoard := GameBoard{Player1: Player{Team:1}, Player2: Player{Team:2}}
	gameBoard.startingBoard()
	// gameBoard.displayBoard()
	chess :=  ChessPlay{GameBoard: gameBoard}
	chess.playerTurn = &chess.GameBoard.Player1
	return chess
}

func (c *ChessPlay) driver(){
	move := c.GameBoard.Player1.selectMove(Position{1,2},Position{1,4})
	c.makeMove(move)
	move = c.GameBoard.Player2.selectMove(Position{2,7},Position{2,5})
	c.makeMove(move)
	move = c.GameBoard.Player1.selectMove(Position{1,4},Position{1,6})
	c.makeMove(move)
	move = c.GameBoard.Player1.selectMove(Position{1,4},Position{2,5})
	c.makeMove(move)
	move = c.GameBoard.Player2.selectMove(Position{2,5},Position{2,4})
	c.makeMove(move)
	move = c.GameBoard.Player2.selectMoveWithString("57","55")
	c.makeMove(move)
	c.makeMove(c.GameBoard.Player1.selectMoveWithString("52","54"))
	c.makeMove(c.GameBoard.Player2.selectMoveWithString("55","54"))
}

func (c ChessPlay) isValidMove(move Move) bool {
	if c.checkPiece(move){
		piece := c.getPiece(move.StartingPosition)
		return c.checkValidLanding(&piece, move.LandingPosition) && c.checkMove(piece, move) && c.checkPath(piece, move)
	}	
	return false
}

func (c *ChessPlay) makeMove(move Move){
	if c.isValidMove(move){
		c.acceptMove(move)
		// c.moveTurn()
	}
	// c.displayBoard()
	if c.Winner.Team != 0 {
		fmt.Println("\n", "Winner: ", c.Winner.Team)
  } else {
		fmt.Println("\n", c.playerTurn, "'s turn")	
	}
}

func (c ChessPlay) checkPiece(move Move) bool{
	piecePlayer := c.squares[move.StartingPosition].Player
	return *c.playerTurn == *move.player && piecePlayer.Team == move.player.Team
}

func (c ChessPlay) getPiece(position Position) GamePiece{
	return c.GameBoard.squares[position]
}

func (c ChessPlay) checkValidLanding(piece *GamePiece, LandingPosition Position) bool{
	fmt.Println("CHECKING LANDING !!!!!!!!!!!!!!!!!!!")
	fmt.Println(c.squares[LandingPosition]," HELPP !!!!")
	if c.squares[LandingPosition].Name != ""{
		landingPositionPiecePlayer := c.squares[LandingPosition].Player
		fmt.Println(landingPositionPiecePlayer," HELPP !!!!")
		if piece.Player.Team ==	landingPositionPiecePlayer.Team {
			if (piece.Name != "rook" && c.squares[LandingPosition].Name != "king"){
				fmt.Println("landing on owned piece")
				return false
			}
		} else if piece.Player.Team !=	landingPositionPiecePlayer.Team &&	landingPositionPiecePlayer.Team != 0 {
			capturedPiece := c.squares[LandingPosition]
			if capturedPiece.Name == "king"{
				c.Winner = *piece.Player
			}
			fmt.Println("\nCaptured ", capturedPiece)
			fmt.Println("")
			piece.capturing = true
		}
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
		if *c.playerTurn == c.GameBoard.Player1{
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
			if c.squares[currentSquare].Player.Team == c.playerTurn.Team && currentSquare != move.StartingPosition {
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
	fmt.Println("ACCEPTING MOVE", c.playerTurn)
	if c.checkCastle(move){
		c.castleKing(move)
		return
	}
	if c.getPiece(move.StartingPosition).Name == "king"{
		fmt.Println("YEP FOUND KING")
		if c.playerTurn.Team == c.GameBoard.Player1.Team {
			c.GameBoard.Player1.king = move.LandingPosition
			fmt.Println(c.GameBoard.Player1.king , "PLAYER 1 POS")
		} else {
			c.GameBoard.Player2.king = move.LandingPosition
			fmt.Println(c.GameBoard.Player2.king , "PLAYER 2 POS")
		}
		fmt.Println("UPDATING KING POSTIOIN", c.playerTurn.king)
	} 
	piece := c.squares[move.StartingPosition]
	oldPiece := c.squares[move.LandingPosition] 
	c.squares[move.StartingPosition] = GamePiece{}
	c.squares[move.LandingPosition] = piece
	if c.isCheck(){
		c.squares[move.StartingPosition] = piece
		c.squares[move.LandingPosition] = oldPiece
		if c.getPiece(move.StartingPosition).Name == "king"{
			fmt.Println("UPDATING KING POSTIOIN", c.playerTurn.king)
			if c.playerTurn.Team == c.GameBoard.Player1.Team{
				c.GameBoard.Player1.king = move.StartingPosition
			} else {
				c.GameBoard.Player2.king = move.StartingPosition
			}
			fmt.Println("UPDATING KING POSTIOIN", c.playerTurn.king)
		} 
	} else {
		piece.Moved = true
		piece.capturing = false
		c.squares[move.LandingPosition] = piece
		c.clearEnPassant()
		c.checkEnPassant(move)
		c.moveTurn()
		fmt.Println(c.playerTurn, "AFTER NOT CHECK SHOULD SWITCH")
	}

}




func (c *ChessPlay) moveTurn(){
	if *c.playerTurn == c.GameBoard.Player1 {
		c.playerTurn = &c.GameBoard.Player2
	} else {
		c.playerTurn = &c.GameBoard.Player1
	}
}


func (c ChessPlay) isCheck() bool{
	var attackingPlayer,  defendingTeam Player
	if c.playerTurn.Team == 1{
		defendingTeam = c.GameBoard.Player1
		attackingPlayer = c.GameBoard.Player2
	} else {
		defendingTeam = c.GameBoard.Player2
		attackingPlayer = c.GameBoard.Player1
	}
	fmt.Println(defendingTeam.king, "DEFENDING King")
	fmt.Println(c.playerTurn, "FROM IN CHECK")
	c.moveTurn()
	fmt.Println(c.playerTurn, "FROM IN CHECK AFTER SWITCH")
	for _, piece := range c.squares{
		if piece.Name != ""{
			if piece.Player.Team == attackingPlayer.Team{
				if c.isValidMove(Move{StartingPosition: piece.Position, LandingPosition: defendingTeam.king, player: c.playerTurn}){
					fmt.Println("CHECK BOI==================================== ", piece, defendingTeam.king)
					return true
				}
			}
		}
	}
	return false
}

