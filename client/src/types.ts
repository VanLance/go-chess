type Player = {
  username?: string
  Team: number
}

type Piece = Position & {
  Name: string
  Username?: string
  Team: number
  MovementTypes: Array<Movement>
  
}

type Position = {
  X: number | ChessColumn
  Y: number
}

type Movement = Position & Condition

type Condition = {
  name: string
  active: boolean
}

type Move = {
  startingPosition: Position
  landingPosition: Position
  playerTurn: Player
}

type ChessColumn = "A" | "B" | "C" | "D" | "E" | "F" | "G" | "H"

export {
  Player,
  Piece,
  Position,
  Movement,
  Condition,
  ChessColumn,
  Move
}