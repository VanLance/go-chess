package main

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
