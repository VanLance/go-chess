import { makeMove, pieces, playerTurn } from "./index"
import { ChessColumn, Move, Piece, Player, Position } from "./types"

let startingPosition: Position
let landingPosition: Position

const xNumbers: {[key:number]:string} = {
  1: 'A',
  2: 'B',
  3: 'C',
  4: 'D',
  5: 'E',
  6: 'F',
  7: 'G',
  8: 'H'
}

function createChessSquares(){
  for(let number=9; number >= 0; number--){
    for (const letter of "0ABCDEFGH0"){
      const div = createChessDiv(number, letter)
      addChessSquare(div)
    }
  }
}

function createChessDiv(number: number, letter: string): HTMLDivElement {
  const div = document.createElement('div')
  div.id = `${letter}${number}`
  if (number === 0 || number === 9){
    div.innerText = letter
  } else if ( letter === '0'){
    div.innerText = number.toString()
  } else {
    div.classList.add('chess-square')
  }
  div.addEventListener('click', ()=>{
    getMove({Team:1},div)
    checkForMove()
  })
  return div
}

function addChessSquare(div: HTMLDivElement){
  document.getElementsByTagName('main')[0]?.appendChild(div)
}

function addPieces(pieces: Array<Piece>){
  for (const piece of pieces){
  const square = document.querySelector(`#${xNumbers[piece.X as number]}${piece.Y}`) as HTMLDivElement
  square.classList.add(`player-${piece.Team}`)
  square.classList.add(`${piece.Name}`)
  if(square){
    const p = document.createElement('p')
    p.innerText = piece.Name + '\n' + piece.Team
    square.appendChild(p)
  }}
}

function clearPieces(){
  const squares = document.getElementsByClassName('chess-square')
  console.log(squares)
  for (const square of squares){
    square.innerHTML = ''
  }
}

function getMove(player: Player, div: HTMLDivElement){
  console.log(div.id)
  const squarePlayer = getPlayerFromDiv(div.classList)
  console.log(squarePlayer?.charAt(-1), player.Team.toString())
  if ( player.Team.toString() === squarePlayer?.[squarePlayer.length -1 ]){
    console.log("TEST")
    startingPosition = getPositionFromDivId(div.id)
    div.classList.add("selected-square")
  } else if (startingPosition.X && startingPosition.Y) {
    landingPosition =getPositionFromDivId(div.id)
  }
  console.log(startingPosition, landingPosition)
  
}

function checkForMove(){
  if( startingPosition.X && landingPosition.X ){
    makeMove(pieces, {startingPosition, landingPosition, playerTurn})
    startingPosition = {X:0,Y:0}
    landingPosition = {X:0,Y:0}
  }
}

function getPositionFromDivId(divId: string): Position{
  return {X:divId[0] as ChessColumn, Y: parseInt(divId[1]!)}
}

function getPlayerFromDiv(divClassList: DOMTokenList): string | void{
  for (const clss of divClassList){
    if (clss.substring(0,6) === 'player'){
      return clss
    }
  }
}

export {
  createChessSquares,
  createChessDiv,
  addChessSquare,
  addPieces,
  clearPieces
}