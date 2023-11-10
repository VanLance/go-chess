package main

func getAbsolute(number int) int{
	if number >= 0{
		return number
	}
	return -number
}

func getSpacesMoved(move Move) Position{
	return Position{getAbsolute(move.LandingPosition.X - move.StartingPosition.X), getAbsolute(move.LandingPosition.Y - move.StartingPosition.Y)}
}