/******/ (() => { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./src/chessSquares.ts":
/*!*****************************!*\
  !*** ./src/chessSquares.ts ***!
  \*****************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   addChessSquare: () => (/* binding */ addChessSquare),
/* harmony export */   addPieces: () => (/* binding */ addPieces),
/* harmony export */   clearPieces: () => (/* binding */ clearPieces),
/* harmony export */   createChessDiv: () => (/* binding */ createChessDiv),
/* harmony export */   createChessSquares: () => (/* binding */ createChessSquares)
/* harmony export */ });
/* harmony import */ var _index__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./index */ "./src/index.ts");
/* harmony import */ var _websocket__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./websocket */ "./src/websocket.ts");


const xNumbers = {
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
function createChessDiv(number, letter) {
    const div = document.createElement('div');
    div.id = `${letter}${number}`;
    if ((number === 0 || number === 9) && letter != ' ' && letter != '1') {
        if (letter != ' ' && letter != '1') {
            div.classList.add('letter');
            div.innerText = letter;
        }
    }
    else if (letter === ' ' || letter === '1') {
        if (number != 9 && number != 0) {
            div.innerText = number.toString();
        }
        if (letter === ' ' || letter === '1') {
            div.classList.add('border-number');
        }
        if (letter === ' ') {
            div.classList.add('left-border');
        }
    }
    else {
        if (number % 2 === 0) {
            div.classList.add('gray');
        }
        else
            div.classList.add('light-gray');
        div.classList.add('chess-square');
    }
    div.addEventListener('click', () => {
        if (_websocket__WEBPACK_IMPORTED_MODULE_1__.player) {
            getPosition(div);
        }
        else {
            getPosition(div, _index__WEBPACK_IMPORTED_MODULE_0__.chessState.playerTurn);
        }
        checkForMove();
    });
    return div;
}
function addChessSquare(div) {
    document.querySelector('.chess-board').appendChild(div);
}
function addPieces(...pieces) {
    for (const piece of pieces) {
        const square = document.querySelector(`#${xNumbers[piece.X]}${piece.Y}`);
        square.classList.add(`player-${piece.Team}`);
        const img = document.createElement('img');
        img.src = `./static/images/${piece.Name}-${piece.Team == 1 ? 'white' : 'black'}.png`;
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
    const player1Divs = document.querySelectorAll('.player-1');
    const player2Divs = document.querySelectorAll('.player-2');
    for (const div of [...player1Divs, ...player2Divs]) {
        div.classList.remove('player-1');
        div.classList.remove('player-2');
    }
    clearActiveSquare();
}
function clearActiveSquare() {
    const activeSquare = document.querySelector(`.selected-square`);
    activeSquare?.classList.toggle('selected-square');
}
function checkCastle(player, div) {
    if (player.Team == 1 && div.id == 'E1') {
        if (_index__WEBPACK_IMPORTED_MODULE_0__.chessState.move?.startingPosition.Y === 1 &&
            (_index__WEBPACK_IMPORTED_MODULE_0__.chessState.move?.startingPosition.X === 'A' ||
                _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move?.startingPosition.X === 'H')) {
            return true;
        }
    }
    else if (player.Team == 2 && div.id == 'E8') {
        if (_index__WEBPACK_IMPORTED_MODULE_0__.chessState.move?.startingPosition.Y === 8 &&
            (_index__WEBPACK_IMPORTED_MODULE_0__.chessState.move?.startingPosition.X === 'A' ||
                _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move?.startingPosition.X === 'H')) {
            return true;
        }
    }
}
function getPosition(div, checkPlayer = null) {
    if (!checkPlayer) {
        checkPlayer = _websocket__WEBPACK_IMPORTED_MODULE_1__.player;
    }
    if (checkPlayer.Team == _index__WEBPACK_IMPORTED_MODULE_0__.chessState.playerTurn.Team) {
        const squarePlayer = getPlayerFromDiv(div.classList);
        if (checkPlayer.Team.toString() === squarePlayer?.[squarePlayer.length - 1] &&
            !checkCastle(checkPlayer, div)) {
            clearActiveSquare();
            _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move = {
                startingPosition: getPositionFromDivId(div.id),
                landingPosition: { X: 0, Y: 0 },
            };
            div.classList.add('selected-square');
        }
        else if (_index__WEBPACK_IMPORTED_MODULE_0__.chessState.move) {
            _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move.landingPosition = getPositionFromDivId(div.id);
        }
    }
}
function checkForMove() {
    if (_index__WEBPACK_IMPORTED_MODULE_0__.chessState.move?.startingPosition.X &&
        _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move?.landingPosition.X) {
        if (_websocket__WEBPACK_IMPORTED_MODULE_1__.player) {
            (0,_websocket__WEBPACK_IMPORTED_MODULE_1__.sendMsg)(_index__WEBPACK_IMPORTED_MODULE_0__.chessState.move);
        }
        else {
            (0,_index__WEBPACK_IMPORTED_MODULE_0__.makeMove)();
        }
        _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move.startingPosition = { X: 0, Y: 0 };
        _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move.landingPosition = { X: 0, Y: 0 };
    }
}
function getPositionFromDivId(divId) {
    return { X: divId[0], Y: parseInt(divId[1]) };
}
function getPlayerFromDiv(divClassList) {
    for (const clss of divClassList) {
        if (clss.substring(0, 6) === 'player') {
            return clss;
        }
    }
}



/***/ }),

/***/ "./src/index.ts":
/*!**********************!*\
  !*** ./src/index.ts ***!
  \**********************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   chessState: () => (/* binding */ chessState),
/* harmony export */   makeMove: () => (/* binding */ makeMove),
/* harmony export */   startGame: () => (/* binding */ startGame)
/* harmony export */ });
/* harmony import */ var _chessSquares__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./chessSquares */ "./src/chessSquares.ts");

let chessState;
async function startGame() {
    const res = await fetch("https://go-chess.onrender.com");
    if (res.ok) {
        const data = await res.json();
        chessState = {
            playerTurn: data.PlayerTurn,
            pieces: [...data.PlayerOnePieces, ...data.PlayerTwoPieces],
        };
        updateChessState(data);
        console.log("Response Data:", data);
        return data;
    }
    else {
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
            move: {
                startingPosition: chessState.move.startingPosition.X + chessState.move.startingPosition.Y.toString(),
                landingPosition: chessState.move.landingPosition.X + chessState.move.landingPosition.Y.toString(),
                player: chessState.playerTurn
            }
        })
    });
    if (res.ok) {
        const data = await res.json();
        updateChessState(data);
        console.log("Response Data:", data);
        return data;
    }
    else {
        console.error('HTTP error:', res.status);
    }
}
;
function updateChessState(data) {
    chessState.playerTurn = data.PlayerTurn;
    chessState.pieces = [];
    chessState.pieces.push(...data.PlayerOnePieces, ...data.PlayerTwoPieces);
    updatePlayerTurn();
    (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.clearPieces)();
    (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.addPieces)(...data.PlayerOnePieces, ...data.PlayerTwoPieces);
}
function updatePlayerTurn() {
    let playerTurnP = document.getElementById('player-turn');
    playerTurnP.innerText = playerTurnP?.innerText.substring(0, playerTurnP.innerHTML?.length - 2) + ' ' + chessState.playerTurn.Team;
}



/***/ }),

/***/ "./src/uiMessage.ts":
/*!**************************!*\
  !*** ./src/uiMessage.ts ***!
  \**************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   updatePTag: () => (/* binding */ updatePTag)
/* harmony export */ });
const pTag = document.querySelector('#player-turn');
function updatePTag() {
    //  pTag.innerText = chessState.playing == false ? "Waiting on Player 2" : pTag.innerText 
    pTag.innerText = "Waiting on Player 2";
}



/***/ }),

/***/ "./src/websocket.ts":
/*!**************************!*\
  !*** ./src/websocket.ts ***!
  \**************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   connect: () => (/* binding */ connect),
/* harmony export */   gameplay: () => (/* binding */ gameplay),
/* harmony export */   player: () => (/* binding */ player),
/* harmony export */   sendMsg: () => (/* binding */ sendMsg)
/* harmony export */ });
/* harmony import */ var _chessSquares__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./chessSquares */ "./src/chessSquares.ts");
/* harmony import */ var _index__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ./index */ "./src/index.ts");
/* harmony import */ var _uiMessage__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./uiMessage */ "./src/uiMessage.ts");



let player;
let connect;
let sendMsg;
let socket;
let gameplay;
const main = document.getElementsByTagName('main')[0];
const gameplayForm = document.querySelector("#gameplay-form");
const selectGameplay = document.querySelector('#game-play');
gameplayForm.addEventListener("submit", async (e) => {
    e.preventDefault();
    (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.createChessSquares)();
    await (0,_index__WEBPACK_IMPORTED_MODULE_1__.startGame)();
    main.classList.toggle('hide');
    gameplayForm.classList.toggle('hide');
    gameplay = selectGameplay.value;
    if (gameplay == 'online') {
        socket = new WebSocket("ws://localhost:8080/ws");
        connect = () => {
            console.log("Attempting Connection...");
            socket.onopen = () => {
                console.log("Successfully Connected");
            };
            socket.onmessage = async (event) => {
                let message = JSON.parse(event.data);
                if (message.body === "player-1") {
                    player = {
                        username: '',
                        Team: 1
                    };
                    _index__WEBPACK_IMPORTED_MODULE_1__.chessState.playing = false;
                }
                else if (message.body === "player-2") {
                    player = {
                        username: '',
                        Team: 2
                    };
                    sendMsg("Player-2 Joined");
                    _index__WEBPACK_IMPORTED_MODULE_1__.chessState.playing = true;
                }
                else if (message.body === "Player-2 Joined") {
                    _index__WEBPACK_IMPORTED_MODULE_1__.chessState.playing = true;
                }
                else {
                    message = JSON.parse(message.body);
                    if (message.hasOwnProperty('startingPosition') && message.hasOwnProperty('landingPosition')) {
                        _index__WEBPACK_IMPORTED_MODULE_1__.chessState.move = message;
                        (0,_index__WEBPACK_IMPORTED_MODULE_1__.makeMove)();
                    }
                }
                console.log("Received message:", message);
            };
            socket.onclose = event => {
                console.log("Socket Closed Connection: ", event);
            };
            socket.onerror = error => {
                console.log("Socket Error: ", error);
            };
        };
        connect();
    }
    (0,_uiMessage__WEBPACK_IMPORTED_MODULE_2__.updatePTag)();
});
sendMsg = (msg) => {
    console.log("sending msg: ", msg);
    socket.send(JSON.stringify(msg));
};



/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		var cachedModule = __webpack_module_cache__[moduleId];
/******/ 		if (cachedModule !== undefined) {
/******/ 			return cachedModule.exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			// no module.id needed
/******/ 			// no module.loaded needed
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		__webpack_modules__[moduleId](module, module.exports, __webpack_require__);
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/ 	
/************************************************************************/
/******/ 	/* webpack/runtime/define property getters */
/******/ 	(() => {
/******/ 		// define getter functions for harmony exports
/******/ 		__webpack_require__.d = (exports, definition) => {
/******/ 			for(var key in definition) {
/******/ 				if(__webpack_require__.o(definition, key) && !__webpack_require__.o(exports, key)) {
/******/ 					Object.defineProperty(exports, key, { enumerable: true, get: definition[key] });
/******/ 				}
/******/ 			}
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/hasOwnProperty shorthand */
/******/ 	(() => {
/******/ 		__webpack_require__.o = (obj, prop) => (Object.prototype.hasOwnProperty.call(obj, prop))
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/make namespace object */
/******/ 	(() => {
/******/ 		// define __esModule on exports
/******/ 		__webpack_require__.r = (exports) => {
/******/ 			if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 				Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 			}
/******/ 			Object.defineProperty(exports, '__esModule', { value: true });
/******/ 		};
/******/ 	})();
/******/ 	
/************************************************************************/
/******/ 	
/******/ 	// startup
/******/ 	// Load entry module and return exports
/******/ 	// This entry module is referenced by other modules so it can't be inlined
/******/ 	var __webpack_exports__ = __webpack_require__("./src/index.ts");
/******/ 	
/******/ })()
;
//# sourceMappingURL=bundle.js.map