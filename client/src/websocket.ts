import { createChessSquares } from "./chessSquares";
import { chessState, makeMove, startGame } from "./index";
import { Player } from "./types";

let player: Player
let connect: () => void
let sendMsg: (msg: any) => void
let socket: WebSocket
let gameplay: 'online' | 'local'


const main = document.getElementsByTagName('main')[0]!

const gameplayForm = document.querySelector("#gameplay-form")! as HTMLFormElement

gameplayForm.addEventListener("submit", async (e:SubmitEvent)=> {
  e.preventDefault()

  createChessSquares()

  main.classList.toggle('hide')

  const selectGameplay = document.querySelector('#game-play')! as HTMLSelectElement

  gameplay = selectGameplay.value as 'online' | 'local'
  
  if (gameplay == 'online'){
    socket = new WebSocket("ws://localhost:8080/ws");
    connect = () => {
      console.log("Attempting Connection...");
      
      socket.onopen = () => {
        console.log("Successfully Connected");
      };
      
      socket.onmessage = async event => {
        let message = JSON.parse(event.data);
        if (message.body === "player-1"){
          player = {
            username: '',
            Team: 1
          }
          chessState.playing = false
        } else if (message.body === "player-2"){
          player = {
            username: '',
            Team: 2
          }
          chessState.playing = true
        } else {
          message = JSON.parse(message.body)
          chessState.move = message
          makeMove()
        }
        console.log("Received message:", message);
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
  await startGame()
})

sendMsg = (msg: any) => {
  console.log("sending msg: ", msg);
  socket.send(msg);
};

export {
  connect, player, sendMsg, gameplay
}