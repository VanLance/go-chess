import { addPieces, clearPieces, createChessSquares } from './chessSquares'
import { Move, Piece, Player } from './types'


let playerTurn: Player
console.log("TEST")

createChessSquares()

async function testApi(){
  const res = await fetch("http://localhost:8080/")

  const data = await res.json()
  playerTurn = data.PlayerTurn
  clearPieces()
  addPieces(data.PlayerOnePieces)
  addPieces(data.PlayerTwoPieces)
  return data
}




async function makeMove(pieces: Array<Piece>, move: Move) {
  // const pieces = await testApi()
  const res = await fetch("http://localhost:8080/make-move", {
    method: "POST",
    headers: {
      "Content-Type": 'application/json'
    },
    body: JSON.stringify({
      previousState: pieces,
      move : {
        startingPosition: move.startingPosition.X + move.startingPosition.Y.toString(),
        landingPosition: move.landingPosition.X + move.landingPosition.Y.toString(),
        player: 1
      }
    })
  })
  if (res.ok) {
    const data = await res.json();
    playerTurn = data.PlayerTurn
    clearPieces()
    addPieces(data.PlayerOnePieces)
    addPieces(data.PlayerTwoPieces)
    pieces.length = 0
    pieces.push(...data.PlayerOnePieces, ...data.PlayerTwoPieces)
    console.log("Response Data:", data);
    return data
  } else {
    console.error('HTTP error:', res.status);
  }
};


(async () => { (await testApi()) })()

const data = await testApi()
console.log(data, "DATA===================")
const pieces = [...data.PlayerOnePieces, ...data.PlayerTwoPieces]

export {
  makeMove,
  playerTurn,
  pieces
}