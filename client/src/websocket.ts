import { createChessSquares } from "./chessSquares";
import { chessState, makeMove, startGame } from "./index";
import { Player } from "./types";

let player: Player
let connect: () => void
let sendMsg: (msg: any) => void
let socket: WebSocket
const main = document.getElementsByTagName('main')[0]!

const gameplayForm = document.querySelector("#gameplay-form")! as HTMLFormElement

gameplayForm.addEventListener("submit", async (e:SubmitEvent)=> {
  e.preventDefault()
  createChessSquares()
  socket = new WebSocket("ws://localhost:8080/ws");
  
  await startGame()
  main.classList.toggle('hide')
  
  connect = () => {
    console.log("Attempting Connection...");
    
    socket.onopen = () => {
      console.log("Successfully Connected");
    };
    
    socket.onmessage = event => {
      let message = JSON.parse(event.data);
      if (message.body === "player-1"){
        player = {
          username: '',
          Team: 1
        }
      } else if (message.body === "player-2"){
        player = {
          username: '',
          Team: 2
        }
      } else {
        message = JSON.parse(message.body)
        chessState.move = message
        makeMove()
      }
      console.log("Received message:", message);
      console.log(player)
    }
    
    socket.onclose = event => {
      console.log("Socket Closed Connection: ", event);
    };
    
    socket.onerror = error => {
      console.log("Socket Error: ", error);
    };
  };

  connect()
}

)
sendMsg = (msg: any) => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export {
  connect, player, sendMsg
}