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
        for (const letter of '0ABCDEFGH0') {
            const div = createChessDiv(number, letter);
            addChessSquare(div);
        }
    }
}
function createChessDiv(number, letter) {
    const div = document.createElement('div');
    div.id = `${letter}${number}`;
    if (number === 0 || number === 9) {
        div.innerText = letter;
    }
    else if (letter === '0') {
        div.innerText = number.toString();
    }
    else {
        div.classList.add('chess-square');
    }
    div.addEventListener('click', () => {
        getMove(_index__WEBPACK_IMPORTED_MODULE_0__.chessState.playerTurn, div);
        checkForMove();
    });
    return div;
}
function addChessSquare(div) {
    document.getElementsByTagName('main')[0]?.appendChild(div);
}
function addPieces(...pieces) {
    for (const piece of pieces) {
        const square = document.querySelector(`#${xNumbers[piece.X]}${piece.Y}`);
        square.classList.add(`player-${piece.Team}`);
        square.classList.add(`${piece.Name}`);
        if (square) {
            const p = document.createElement('p');
            p.innerText = piece.Name + '\n' + piece.Team;
            square.appendChild(p);
        }
    }
}
function clearPieces() {
    const squares = document.getElementsByClassName('chess-square');
    for (const square of squares) {
        clearClasses(square);
        square.innerHTML = '';
    }
}
function clearClasses(div) {
    for (let i = 0; i < div.classList.length; i++) {
        const clss = div.classList[i];
        if (clss !== 'chess-square') {
            div.classList.remove(clss);
            i--;
        }
    }
}
function getMove(player, div) {
    const squarePlayer = getPlayerFromDiv(div.classList);
    if (player.Team.toString() === squarePlayer?.[squarePlayer.length - 1]) {
        _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move = { startingPosition: getPositionFromDivId(div.id), landingPosition: { X: 0, Y: 0 } };
        div.classList.add('selected-square');
    }
    else if (_index__WEBPACK_IMPORTED_MODULE_0__.chessState.move?.startingPosition.X && _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move.startingPosition.Y) {
        _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move.landingPosition = getPositionFromDivId(div.id);
    }
}
function checkForMove() {
    if (_index__WEBPACK_IMPORTED_MODULE_0__.chessState.move?.startingPosition.X && _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move?.landingPosition.X) {
        (0,_index__WEBPACK_IMPORTED_MODULE_0__.makeMove)();
        _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move.startingPosition = { X: 0, Y: 0 };
        _index__WEBPACK_IMPORTED_MODULE_0__.chessState.move.landingPosition = { X: 0, Y: 0 };
        console.log(_index__WEBPACK_IMPORTED_MODULE_0__.chessState);
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
/* harmony export */   makeMove: () => (/* binding */ makeMove)
/* harmony export */ });
/* harmony import */ var _chessSquares__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./chessSquares */ "./src/chessSquares.ts");

let chessState;
(0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.createChessSquares)();
async function getStartingPieces() {
    const res = await fetch("http://localhost:8080/");
    const data = await res.json();
    chessState = {
        playerTurn: data.PlayerTurn,
        pieces: [...data.PlayerOnePieces, ...data.PlayerTwoPieces],
    };
    console.log(chessState, 'getStarted');
    (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.clearPieces)();
    (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.addPieces)(...data.PlayerOnePieces, ...data.PlayerTwoPieces);
    console.log(chessState.playerTurn, 'getStarted');
    return data;
}
async function makeMove() {
    console.log(chessState, 'MAKE MOVE');
    const res = await fetch("http://localhost:8080/make-move", {
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
        console.log(data.PlayerTurn, "DATA FROM MOVE 2!!!!!!!!!");
        udpateData(data);
        console.log(data.PlayerTurn, "DATA FROM MOVE 2 2");
        console.log("Response Data:", data);
        // chessState.playerTurn = data.PlayerTurn
        // console.log(chessState, 'When it matters Make Moves')
        return data;
    }
    else {
        console.error('HTTP error:', res.status);
    }
}
;
function udpateData(data) {
    chessState.playerTurn = data.PlayerTurn;
    chessState.pieces = [];
    chessState.pieces.push(...data.PlayerOnePieces, ...data.PlayerTwoPieces);
    console.log(chessState, 'Wen it matters');
    (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.clearPieces)();
    (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.addPieces)(...data.PlayerOnePieces, ...data.PlayerTwoPieces);
}
(async () => { (await getStartingPieces()); })();



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