import { makeMove, chessState } from './index';
import { ChessColumn, Move, Piece, Player, Position } from './types';
import { player, sendMsg } from "./websocket"

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
  if ((number === 0 || number === 9) && letter != ' ' && letter != '1') {
    if (letter != ' ' && letter != '1') {
      div.classList.add('letter')
      div.innerText = letter;
    }
  } else if (letter === ' ' || letter === '1') {
    if (number != 9 && number != 0) {
      div.innerText = number.toString();
    } 
    if (letter === ' ' || letter === '1') {
      div.classList.add('border-number');
    }
    if (letter === ' ') {
      div.classList.add('left-border');
    }
  } else {
    if (number % 2 === 0) {
      div.classList.add('gray');
    } else div.classList.add('light-gray');
    div.classList.add('chess-square');
  }
  div.addEventListener('click', () => {
    if (player) {
      getPosition(div);
    } else {
      getPosition(div, chessState.playerTurn);
    }
    checkForMove();
  });
  return div;
}

function addChessSquare(div: HTMLDivElement) {
  document.querySelector('.chess-board')!.appendChild(div);
}

function addPieces(...pieces: Array<Piece>) {
  for (const piece of pieces) {
    const square = document.querySelector(
      `#${xNumbers[piece.X as number]}${piece.Y}`
    ) as HTMLDivElement;
    square.classList.add(`player-${piece.Team}`);
    const img = document.createElement('img');
    img.src = `./static/images/${piece.Name}-${
      piece.Team == 1 ? 'white' : 'black'
    }.png`;
    square.appendChild(img);
  }
}

function clearPieces() {
  const squares = document.getElementsByClassName('chess-square');
  for (const square of squares) {
    clearClasses();
    square.innerHTML = '';
  }
}

function clearClasses() {
  const player1Divs = document.querySelectorAll('.player-1')
  const player2Divs = document.querySelectorAll('.player-2')
  for (const div of [...player1Divs, ...player2Divs]) {
    div.classList.remove('player-1')
    div.classList.remove('player-2')
  }
  clearActiveSquare();
}

function clearActiveSquare() {
  const activeSquare = document.querySelector(`.selected-square`);
  activeSquare?.classList.toggle('selected-square');
}

function checkCastle(player: Player, div: HTMLDivElement) {
  if (player.Team == 1 && div.id == 'E1') {
    if (
      chessState.move?.startingPosition.Y === 1 &&
      (chessState.move?.startingPosition.X === 'A' ||
        chessState.move?.startingPosition.X === 'H')
    ) {
      return true;
    }
  } else if (player.Team == 2 && div.id == 'E8') {
    if (
      chessState.move?.startingPosition.Y === 8 &&
      (chessState.move?.startingPosition.X === 'A' ||
        chessState.move?.startingPosition.X === 'H')
    ) {
      return true;
    }
  }
}

function getPosition( div: HTMLDivElement, checkPlayer: Player | null = null) {
  if ( !checkPlayer ) {
    checkPlayer = player
  }
  if ( checkPlayer.Team == chessState.playerTurn.Team ){
    const squarePlayer = getPlayerFromDiv(div.classList);
    if (
      checkPlayer.Team.toString() === squarePlayer?.[squarePlayer.length - 1] &&
      !checkCastle(checkPlayer, div)
    ) {
      clearActiveSquare();
      chessState.move = {
        startingPosition: getPositionFromDivId(div.id),
        landingPosition: { X: 0, Y: 0 },
      };
      div.classList.add('selected-square');
    } else if (chessState.move) {
      chessState.move.landingPosition = getPositionFromDivId(div.id);
    }
  }
}

function checkForMove() {
  if (
    chessState.move?.startingPosition.X &&
    chessState.move?.landingPosition.X
    ) {
      if (player) {
        sendMsg(JSON.stringify(chessState.move))
      } else {
        makeMove()
      }
    chessState.move.startingPosition = { X: 0, Y: 0 };
    chessState.move.landingPosition = { X: 0, Y: 0 };
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
