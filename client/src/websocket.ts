import { chessState, makeMove } from "./index";
import { Player } from "./types";

var socket = new WebSocket("ws://localhost:8080/ws");
let player: Player

let connect = () => {
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

let sendMsg = (msg: any)=> {
  console.log("sending msg: ", msg);
  socket.send(msg);
};


document.getElementById('webpack-broadcast')?.addEventListener('click',()=>{sendMsg("TEST")})

export {
  connect, sendMsg, player
}