import { addPieces, clearPieces, createChessSquares } from './chessSquares'
import { ChessState } from './types'

let chessState: ChessState


createChessSquares()

async function getStartingPieces(){
  const res = await fetch("http://localhost:8080/")

  const data = await res.json()
  chessState = {
    playerTurn: data.PlayerTurn,
    pieces: [...data.PlayerOnePieces, ...data.PlayerTwoPieces],
  }
  console.log(chessState, 'getStarted')
  clearPieces()
  addPieces(...data.PlayerOnePieces, ...data.PlayerTwoPieces)
  console.log(chessState.playerTurn, 'getStarted')
  return data
}


async function makeMove() {
  console.log(chessState, 'MAKE MOVE')
  const res = await fetch("http://localhost:8080/make-move", {
    method: "POST",
    headers: {
      "Content-Type": 'application/json'
    },
    body: JSON.stringify({
      previousState: chessState.pieces,
      move : {
        startingPosition: chessState.move!.startingPosition.X + chessState.move!.startingPosition.Y.toString(),
        landingPosition: chessState.move!.landingPosition.X + chessState.move!.landingPosition.Y.toString(),
        player: chessState.playerTurn
      }
    })
  })
  if (res.ok) {
    const data = await res.json();
    console.log(data.PlayerTurn, "DATA FROM MOVE 2!!!!!!!!!")
    udpateData(data)
    console.log(data.PlayerTurn, "DATA FROM MOVE 2 2")
    console.log("Response Data:", data);
    // chessState.playerTurn = data.PlayerTurn
    // console.log(chessState, 'When it matters Make Moves')
    return data
  } else {
    console.error('HTTP error:', res.status);
  }
};

function udpateData(data : any){
    chessState.playerTurn = data.PlayerTurn
    chessState.pieces = []
    chessState.pieces.push(...data.PlayerOnePieces, ...data.PlayerTwoPieces)
    console.log(chessState, 'Wen it matters')
    clearPieces()
    addPieces(...data.PlayerOnePieces, ...data.PlayerTwoPieces)
}

(async () => { (await getStartingPieces()) })()



export {
  makeMove,
  chessState,
}