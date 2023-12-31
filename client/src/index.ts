import { addPieces, clearPieces } from './chessSquares'
import { ChessState } from './types'

let chessState: ChessState 

async function startGame(){
  const res = await fetch("https://go-chess.onrender.com")
  if (res.ok) {
    const data = await res.json()
    chessState = {
      playerTurn: data.PlayerTurn,
      pieces: [...data.PlayerOnePieces, ...data.PlayerTwoPieces],
    }
    updateChessState(data)
    console.log("Response Data:", data);
    return data
  } else {
    console.error('HTTP error:', res.status);
  }
}

async function makeMove() {
  const res = await fetch("https://go-chess.onrender.com/make-move", {
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
    updateChessState(data)
    console.log("Response Data:", data);
    return data
  } else {
    console.error('HTTP error:', res.status);
  }
};

function updateChessState(data: any){
    chessState.playerTurn = data.PlayerTurn
    chessState.pieces = []
    chessState.pieces.push(...data.PlayerOnePieces, ...data.PlayerTwoPieces)
    updatePlayerTurn()
    clearPieces()
    addPieces(...data.PlayerOnePieces, ...data.PlayerTwoPieces)
}

function updatePlayerTurn(){
  let playerTurnP = document.getElementById('player-turn')!
  playerTurnP.innerText = playerTurnP?.innerText.substring(0, playerTurnP.innerHTML?.length -2) + ' ' + chessState.playerTurn.Team
}

export {
  makeMove,
  chessState,
  startGame
}

