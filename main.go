package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type JSONRes struct {
	PlayerOnePieces []GamePiece
	PlayerTwoPieces []GamePiece
	PlayerTurn Player
	Winner Player
	Message string
}

type MoveReq struct {
	ClientMove `json:"move"`
	PreviousState []GamePiece `json:"previousState"`
}

type ClientMove struct {
	StartingPosition string `json:"startingPosition"`
	LandingPosition string `json:"landingPosition"`
	Player int `json:"player"`
}


func handleStart(w http.ResponseWriter, r *http.Request) {
	chess := createChess()
	playerOnePieces := []GamePiece{}
	playerTwoPieces := []GamePiece{}
	for position, square := range chess.GameBoard.squares{
		if square.gamePiece.Name != ""{
			square.gamePiece.Position = position
			if square.gamePiece.Team == 1{
				playerOnePieces = append(playerOnePieces, square.gamePiece)
			} else {
				playerTwoPieces= append(playerTwoPieces, square.gamePiece)
			}
		}
	}
	jsonRes := JSONRes{PlayerOnePieces: playerOnePieces, PlayerTwoPieces: playerTwoPieces, PlayerTurn: chess.player1, Message: "Starting Pieces" }
	w.Header().Set("Content-Type","application/json")
	res, err := json.Marshal(jsonRes)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func handleMove(w http.ResponseWriter, r *http.Request){
	var m MoveReq
	
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil{
		fmt.Println(err)
	}
	chess := recreateBoard(m.PreviousState)
	fmt.Println(m.ClientMove.StartingPosition == "", m.LandingPosition == "", chess.playerTurn , "++++++=========")
	fmt.Println(m.ClientMove.StartingPosition, m.ClientMove.LandingPosition, chess.playerTurn , "++++++=========")
	move := chess.playerTurn.selectMoveWithString(m.StartingPosition , m.LandingPosition)
	// fmt.Println(move, chess.GameBoard.squares[Position{X:1, Y:2}], "===================")
	chess.makeMove(move)
	fmt.Println("MADE MOVE")
	chess.displayBoard()
}

func main() {
	http.HandleFunc("/", handleStart)
	http.HandleFunc("/make-move", handleMove)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func recreateBoard(pieces []GamePiece) ChessPlay{
	chess :=  ChessPlay{GameBoard: GameBoard{}, player1:Player{Team:1}, player2: Player{Team:2} }
	chess.addSquares()
	fmt.Println("ADDED SQUARES")
	chess.playerTurn = chess.player1
	for _, piece := range pieces {
		fmt.Println(piece)
		square := chess.GameBoard.squares[piece.Position]
		square.gamePiece = piece
		chess.GameBoard.squares[piece.Position] = square
	}
	chess.displayBoard()
	return chess
}