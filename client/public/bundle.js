/******/ (() => { // webpackBootstrap
/******/ 	"use strict";
/******/ 	var __webpack_modules__ = ({

/***/ "./src/chessSquares.ts":
/*!*****************************!*\
  !*** ./src/chessSquares.ts ***!
  \*****************************/
/***/ ((module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.a(module, async (__webpack_handle_async_dependencies__, __webpack_async_result__) => { try {
__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   addChessSquare: () => (/* binding */ addChessSquare),
/* harmony export */   addPieces: () => (/* binding */ addPieces),
/* harmony export */   clearPieces: () => (/* binding */ clearPieces),
/* harmony export */   createChessDiv: () => (/* binding */ createChessDiv),
/* harmony export */   createChessSquares: () => (/* binding */ createChessSquares)
/* harmony export */ });
/* harmony import */ var _index__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./index */ "./src/index.ts");
var __webpack_async_dependencies__ = __webpack_handle_async_dependencies__([_index__WEBPACK_IMPORTED_MODULE_0__]);
_index__WEBPACK_IMPORTED_MODULE_0__ = (__webpack_async_dependencies__.then ? (await __webpack_async_dependencies__)() : __webpack_async_dependencies__)[0];

let startingPosition;
let landingPosition;
const xNumbers = {
    1: 'A',
    2: 'B',
    3: 'C',
    4: 'D',
    5: 'E',
    6: 'F',
    7: 'G',
    8: 'H'
};
function createChessSquares() {
    for (let number = 9; number >= 0; number--) {
        for (const letter of "0ABCDEFGH0") {
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
        getMove({ Team: 1 }, div);
        checkForMove();
    });
    return div;
}
function addChessSquare(div) {
    document.getElementsByTagName('main')[0]?.appendChild(div);
}
function addPieces(pieces) {
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
    console.log(squares);
    for (const square of squares) {
        square.innerHTML = '';
    }
}
function getMove(player, div) {
    console.log(div.id);
    const squarePlayer = getPlayerFromDiv(div.classList);
    console.log(squarePlayer?.charAt(-1), player.Team.toString());
    if (player.Team.toString() === squarePlayer?.[squarePlayer.length - 1]) {
        console.log("TEST");
        startingPosition = getPositionFromDivId(div.id);
        div.classList.add("selected-square");
    }
    else if (startingPosition.X && startingPosition.Y) {
        landingPosition = getPositionFromDivId(div.id);
    }
    console.log(startingPosition, landingPosition);
}
function checkForMove() {
    if (startingPosition.X && landingPosition.X) {
        (0,_index__WEBPACK_IMPORTED_MODULE_0__.makeMove)(_index__WEBPACK_IMPORTED_MODULE_0__.pieces, { startingPosition, landingPosition, playerTurn: _index__WEBPACK_IMPORTED_MODULE_0__.playerTurn });
        startingPosition = { X: 0, Y: 0 };
        landingPosition = { X: 0, Y: 0 };
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


__webpack_async_result__();
} catch(e) { __webpack_async_result__(e); } });

/***/ }),

/***/ "./src/index.ts":
/*!**********************!*\
  !*** ./src/index.ts ***!
  \**********************/
/***/ ((module, __webpack_exports__, __webpack_require__) => {

__webpack_require__.a(module, async (__webpack_handle_async_dependencies__, __webpack_async_result__) => { try {
__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   makeMove: () => (/* binding */ makeMove),
/* harmony export */   pieces: () => (/* binding */ pieces),
/* harmony export */   playerTurn: () => (/* binding */ playerTurn)
/* harmony export */ });
/* harmony import */ var _chessSquares__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! ./chessSquares */ "./src/chessSquares.ts");
var __webpack_async_dependencies__ = __webpack_handle_async_dependencies__([_chessSquares__WEBPACK_IMPORTED_MODULE_0__]);
_chessSquares__WEBPACK_IMPORTED_MODULE_0__ = (__webpack_async_dependencies__.then ? (await __webpack_async_dependencies__)() : __webpack_async_dependencies__)[0];

let playerTurn;
console.log("TEST");
(0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.createChessSquares)();
async function testApi() {
    const res = await fetch("http://localhost:8080/");
    const data = await res.json();
    playerTurn = data.PlayerTurn;
    (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.clearPieces)();
    (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.addPieces)(data.PlayerOnePieces);
    (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.addPieces)(data.PlayerTwoPieces);
    return data;
}
async function makeMove(pieces, move) {
    // const pieces = await testApi()
    const res = await fetch("http://localhost:8080/make-move", {
        method: "POST",
        headers: {
            "Content-Type": 'application/json'
        },
        body: JSON.stringify({
            previousState: pieces,
            move: {
                startingPosition: move.startingPosition.X + move.startingPosition.Y.toString(),
                landingPosition: move.landingPosition.X + move.landingPosition.Y.toString(),
                player: 1
            }
        })
    });
    if (res.ok) {
        const data = await res.json();
        playerTurn = data.PlayerTurn;
        (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.clearPieces)();
        (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.addPieces)(data.PlayerOnePieces);
        (0,_chessSquares__WEBPACK_IMPORTED_MODULE_0__.addPieces)(data.PlayerTwoPieces);
        pieces.length = 0;
        pieces.push(...data.PlayerOnePieces, ...data.PlayerTwoPieces);
        console.log("Response Data:", data);
        return data;
    }
    else {
        console.error('HTTP error:', res.status);
    }
}
;
(async () => { (await testApi()); })();
const data = await testApi();
console.log(data, "DATA===================");
const pieces = [...data.PlayerOnePieces, ...data.PlayerTwoPieces];


__webpack_async_result__();
} catch(e) { __webpack_async_result__(e); } }, 1);

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
/******/ 	/* webpack/runtime/async module */
/******/ 	(() => {
/******/ 		var webpackQueues = typeof Symbol === "function" ? Symbol("webpack queues") : "__webpack_queues__";
/******/ 		var webpackExports = typeof Symbol === "function" ? Symbol("webpack exports") : "__webpack_exports__";
/******/ 		var webpackError = typeof Symbol === "function" ? Symbol("webpack error") : "__webpack_error__";
/******/ 		var resolveQueue = (queue) => {
/******/ 			if(queue && queue.d < 1) {
/******/ 				queue.d = 1;
/******/ 				queue.forEach((fn) => (fn.r--));
/******/ 				queue.forEach((fn) => (fn.r-- ? fn.r++ : fn()));
/******/ 			}
/******/ 		}
/******/ 		var wrapDeps = (deps) => (deps.map((dep) => {
/******/ 			if(dep !== null && typeof dep === "object") {
/******/ 				if(dep[webpackQueues]) return dep;
/******/ 				if(dep.then) {
/******/ 					var queue = [];
/******/ 					queue.d = 0;
/******/ 					dep.then((r) => {
/******/ 						obj[webpackExports] = r;
/******/ 						resolveQueue(queue);
/******/ 					}, (e) => {
/******/ 						obj[webpackError] = e;
/******/ 						resolveQueue(queue);
/******/ 					});
/******/ 					var obj = {};
/******/ 					obj[webpackQueues] = (fn) => (fn(queue));
/******/ 					return obj;
/******/ 				}
/******/ 			}
/******/ 			var ret = {};
/******/ 			ret[webpackQueues] = x => {};
/******/ 			ret[webpackExports] = dep;
/******/ 			return ret;
/******/ 		}));
/******/ 		__webpack_require__.a = (module, body, hasAwait) => {
/******/ 			var queue;
/******/ 			hasAwait && ((queue = []).d = -1);
/******/ 			var depQueues = new Set();
/******/ 			var exports = module.exports;
/******/ 			var currentDeps;
/******/ 			var outerResolve;
/******/ 			var reject;
/******/ 			var promise = new Promise((resolve, rej) => {
/******/ 				reject = rej;
/******/ 				outerResolve = resolve;
/******/ 			});
/******/ 			promise[webpackExports] = exports;
/******/ 			promise[webpackQueues] = (fn) => (queue && fn(queue), depQueues.forEach(fn), promise["catch"](x => {}));
/******/ 			module.exports = promise;
/******/ 			body((deps) => {
/******/ 				currentDeps = wrapDeps(deps);
/******/ 				var fn;
/******/ 				var getResult = () => (currentDeps.map((d) => {
/******/ 					if(d[webpackError]) throw d[webpackError];
/******/ 					return d[webpackExports];
/******/ 				}))
/******/ 				var promise = new Promise((resolve) => {
/******/ 					fn = () => (resolve(getResult));
/******/ 					fn.r = 0;
/******/ 					var fnQueue = (q) => (q !== queue && !depQueues.has(q) && (depQueues.add(q), q && !q.d && (fn.r++, q.push(fn))));
/******/ 					currentDeps.map((dep) => (dep[webpackQueues](fnQueue)));
/******/ 				});
/******/ 				return fn.r ? promise : getResult();
/******/ 			}, (err) => ((err ? reject(promise[webpackError] = err) : outerResolve(exports)), resolveQueue(queue)));
/******/ 			queue && queue.d < 0 && (queue.d = 0);
/******/ 		};
/******/ 	})();
/******/ 	
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