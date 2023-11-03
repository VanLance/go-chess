package main

func main() {
	gameBoard := GameBoard{}
	gameBoard.startingBoard()
	gameBoard.displayBoard()
	chess := ChessPlay{GameBoard: gameBoard, player1:Player{name:"1"}, player2: Player{"2"}}
	chess.playerTurn = chess.player1
	chess.driver()
}
