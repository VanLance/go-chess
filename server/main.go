package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"go-chess/pkg/websocket"
)

func createBoardRes(chess ChessPlay, message string) JSONRes{
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
	return JSONRes { 
		PlayerOnePieces: playerOnePieces, 
		PlayerTwoPieces: playerTwoPieces,
		PlayerTurn: *chess.playerTurn,
		Message: message, 
	}
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	chess := createChess()
	jsonRes := createBoardRes(chess, "starting pieces")
	w.Header().Set("Content-Type","application/json")
	res, err := json.Marshal(jsonRes)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func handleMove( w http.ResponseWriter, r *http.Request){
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
	jsonRes := createBoardRes(chess, "move accepted")
	w.Header().Set("Content-Type","application/json")
	res, err := json.Marshal(jsonRes)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w,r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	
	client.Read()
}

var pool *websocket.Pool = websocket.NewPool()

func setupRoutes(){
	pool = websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		serveWs(pool, w, r)
	})
	http.HandleFunc("/", handleStart)
	http.HandleFunc("/make-move", handleMove)
}

func main() {
	setupRoutes()
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

func handleLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add your logging logic here
		println("Request received:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
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

