import { makeMove, chessState } from './index';
import { ChessColumn, Move, Piece, Player, Position } from './types';

const xNumbers: { [key: number]: string } = {
  1: 'A',
  2: 'B',
  3: 'C',
  4: 'D',
  5: 'E',
  6: 'F',
  7: 'G',
  8: 'H',
};

function createChessSquares() {
  for (let number = 9; number >= 0; number--) {
    for (const letter of ' ABCDEFGH1') {
      const div = createChessDiv(number, letter);
      addChessSquare(div);
    }
  }
}

function createChessDiv(number: number, letter: string): HTMLDivElement {
  const div = document.createElement('div');
  div.id = `${letter}${number}`;
  if ((number === 0 || number === 9 ) && letter != ' ' && letter != '1') {
    if (letter != ' ' && letter != '1'){
      div.innerText = letter;}
  } else if (letter === ' ' || letter === '1') {
    if (number != 9 && number != 0 ){ 
      div.innerText = number.toString();
    }
    if (letter === ' ' || letter === '1' ){
      div.classList.add('border-number')
    }
    if (letter === ' '){
      
      div.classList.add('left-border')
    }
  } else {
    div.classList.add('chess-square');
  }
  div.addEventListener('click', () => {
    getMove(chessState.playerTurn, div);
    checkForMove();
  });
  return div;
}

function addChessSquare(div: HTMLDivElement) {
  document.getElementsByTagName('main')[0]?.appendChild(div);
}

function addPieces(...pieces: Array<Piece>) {
  for (const piece of pieces) {
    const square = document.querySelector(
      `#${xNumbers[piece.X as number]}${piece.Y}`
    ) as HTMLDivElement;
    square.classList.add(`player-${piece.Team}`);
    square.classList.add(`${piece.Name}`);
    const img = document.createElement('img');
    console.log(`./static/images/${piece.Name}-${piece.Team == 1? "white":"black"}`)
    img.src = `./static/images/${piece.Name}-${piece.Team == 1? "white":"black"}.png`
    square.appendChild(img);
  }
}

function clearPieces() {
  const squares = document.getElementsByClassName('chess-square');
  for (const square of squares) {
    clearClasses(square);
    square.innerHTML = '';
  }
}

function clearClasses(div: Element) {
  for (let i = 0; i < div.classList.length; i++) {
    const clss = div.classList[i];
    if (clss !== 'chess-square') {
      div.classList.remove(clss!);
      i--;
    }
  }
}

function clearActiveSquare(){
  const activeSquare = document.querySelector(`.selected-square`)
  activeSquare?.classList.toggle('selected-square')
}

function checkCastle(player: Player, div: HTMLDivElement){
  if ( player.Team == 1 && div.id == 'E1'){
    if (chessState.move?.startingPosition.Y === 1 && (chessState.move?.startingPosition.X === 'A' || chessState.move?.startingPosition.X === 'H') ){
      return true
    }
  } else if (player.Team == 2 && div.id == 'E8') {
    if(chessState.move?.startingPosition.Y === 8 && (chessState.move?.startingPosition.X === 'A' || chessState.move?.startingPosition.X === 'H') ){
      return true
    }
  }
}


function getMove(player: Player, div: HTMLDivElement) {
  const squarePlayer = getPlayerFromDiv(div.classList);
  if ( player.Team.toString() === squarePlayer?.[squarePlayer.length - 1] && !checkCastle(player, div) ) {
    clearActiveSquare()
    chessState.move = {startingPosition : getPositionFromDivId(div.id), landingPosition : {X:0,Y:0}};
    div.classList.add('selected-square');
  } else if (chessState.move) {
    chessState.move.landingPosition = getPositionFromDivId(div.id);
  }
}

function checkForMove() {
  if (chessState.move?.startingPosition.X && chessState.move?.landingPosition.X) {
    makeMove()
    chessState.move.startingPosition = { X: 0, Y: 0 };
    chessState.move.landingPosition = { X: 0, Y: 0 };
    console.log(chessState);
  }
}

function getPositionFromDivId(divId: string): Position {
  return { X: divId[0] as ChessColumn, Y: parseInt(divId[1]!) };
}

function getPlayerFromDiv(divClassList: DOMTokenList): string | void {
  for (const clss of divClassList) {
    if (clss.substring(0, 6) === 'player') {
      return clss;
    }
  }
}

export {
  createChessSquares,
  createChessDiv,
  addChessSquare,
  addPieces,
  clearPieces,
};
