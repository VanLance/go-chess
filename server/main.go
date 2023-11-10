package main

import (
	"encoding/json"
	"fmt"
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
	Player `json:"player"`
}


func handleStart(w http.ResponseWriter, r *http.Request) {
	chess := createChess()
	chess.playerTurn = chess.player1
	jsonRes := createBoardRes(chess)
	w.Header().Set("Content-Type","application/json")
	res, err := json.Marshal(jsonRes)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func createBoardRes(chess ChessPlay) JSONRes{
	playerOnePieces := []GamePiece{}
	playerTwoPieces := []GamePiece{}
	for position, square := range chess.GameBoard.squares{
		if square.Name != ""{
			square.Position = position
			if square.Team == 1{
				playerOnePieces = append(playerOnePieces, square)
			} else {
				playerTwoPieces= append(playerTwoPieces, square)
			}
		}
	}
	return JSONRes{PlayerOnePieces: playerOnePieces, PlayerTwoPieces: playerTwoPieces, PlayerTurn: chess.playerTurn, Message: "Starting Pieces" }
}

func handleMove(w http.ResponseWriter, r *http.Request){
	var m MoveReq
	
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil{
		fmt.Println("Error decoding:", err)
	}
	fmt.Printf("Received data: %+v\n", m)
	
	fmt.Printf("StartingPosition: %s, LandingPosition: %s, PlayerTurn: %+v\n", m.ClientMove.StartingPosition, m.ClientMove.LandingPosition, m.Player)
	chess := recreateBoard(m.PreviousState, m.Player)

	move := chess.playerTurn.selectMoveWithString(m.ClientMove.StartingPosition, m.ClientMove.LandingPosition)
	chess.makeMove(move)
	chess.displayBoard()
	jsonRes := createBoardRes(chess)
	w.Header().Set("Content-Type","application/json")
	res, err := json.Marshal(jsonRes)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func main() {
	http.HandleFunc("/", handleStart)
	http.HandleFunc("/make-move", handleMove)
	// Set up the middleware
	handler := handleLogging(handleCORS(http.DefaultServeMux))

	// Use the handler for server
	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}
  if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func recreateBoard(pieces []GamePiece, player Player) ChessPlay{
	chess :=  ChessPlay{GameBoard: GameBoard{}, player1:Player{Team:1}, player2: Player{Team:2} }
	chess.addSquares()
	chess.playerTurn = player
	for _, piece := range pieces {
		square := chess.GameBoard.squares[piece.Position]
		square = piece
		chess.GameBoard.squares[piece.Position] = square
		if piece.Name == "king"{
			piece.Player.king = piece.Position
		}
	}
	chess.displayBoard()
	return chess
}

func handleCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func handleLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add your logging logic here
		println("Request received:", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}