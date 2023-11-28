package main

import (
	"fmt"
	// "strconv"
)

type ChessPlay struct {
	GameBoard
	player1 Player
	player2 Player
	playerTurn *Player
	Winner Player
	Message string
}

type Move struct{
	StartingPosition Position
	LandingPosition Position
	player *Player
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
		c.acceptMove(c.squares[move.StartingPosition], move)
		// c.moveTurn()
	}
	// c.displayBoard()
	if c.Winner.Team != 0 {
		fmt.Println("\n", "Winner: ", c.Winner.Team)
  } else {
		fmt.Println("\n Player ", c.playerTurn.Team, "'s turn")	
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
	LandingPositionPiecePlayer := c.squares[LandingPosition].Player
	if piece.Player.Team == LandingPositionPiecePlayer.Team {
		if (piece.Name != "rook" && c.squares[LandingPosition].Name != "king"){
			fmt.Println("landing on owned piece")
			return false
		}
	} else if piece.Player.Team != LandingPositionPiecePlayer.Team && LandingPositionPiecePlayer.Team != 0 {
		capturedPiece := c.squares[LandingPosition]
		if capturedPiece.Name == "king"{
			c.Winner = piece.Player
		}
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
		if *c.playerTurn == c.player1{
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
			if c.squares[currentSquare].Player.Team != 0 && currentSquare != move.StartingPosition {
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

func (c *ChessPlay) acceptMove(piece GamePiece,move Move){
	if c.checkCastle(move){
		c.castleKing(move)
		c.moveTurn()
		return
	}
	if piece.Name == "king" {
		move.player.updateKingPosition(move.LandingPosition) 
	} 
	oldPiece := c.squares[move.LandingPosition] 
	c.squares[move.StartingPosition] = GamePiece{}
	c.squares[move.LandingPosition] = piece
	if c.isCheck(){
		c.squares[move.StartingPosition] = piece
		c.squares[move.LandingPosition] = oldPiece
		if c.getPiece(move.StartingPosition).Name == "king"{
			if c.playerTurn.Team == c.player1.Team{
				c.player1.king = move.StartingPosition
			} else {
				c.player2.king = move.StartingPosition
			}
			// move.player.updateKingPosition(move.StartingPosition) 
		} 
	} else {
		piece.Moved = true
		piece.capturing = false
		c.squares[move.LandingPosition] = piece
		c.clearEnPassant()
		c.checkEnPassant(move)
		c.checkPawnPosition(piece,move.LandingPosition)
		c.moveTurn()
	}

}

func (c *ChessPlay) moveTurn(){
	if c.playerTurn.Team == c.player1.Team {
		c.playerTurn = &c.player2
	} else {
		c.playerTurn = &c.player1
	}
}


func (c ChessPlay) isCheck() bool{
	var attackingPlayer,  defendingTeam Player
	if c.playerTurn.Team == 1{
		defendingTeam = c.player1
		attackingPlayer = c.player2
	} else {
		defendingTeam = c.player2
		attackingPlayer = c.player1
	}
	c.moveTurn()
	for _, piece := range c.squares{
		if piece.Name != ""{
			if piece.Player.Team == attackingPlayer.Team{
				if c.isValidMove(Move{StartingPosition: piece.Position, LandingPosition: defendingTeam.king, player: c.playerTurn}){
					fmt.Println("CHECK", piece, defendingTeam.king)
					return true
				}
			}
		}
	}
	return false
}

