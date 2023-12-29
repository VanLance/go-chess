import { chessState } from "./index"

const pTag = document.querySelector('#player-turn')! as HTMLParagraphElement

function updatePTag(){
//  pTag.innerText = chessState.playing == false ? "Waiting on Player 2" : pTag.innerText 
 pTag.innerText = "Waiting on Player 2" 
}

export {
  updatePTag
}