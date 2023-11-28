	package websocket

	import (
		"fmt"
	)

	type Pool struct {
		Register   chan *Client
		Unregister chan *Client
		Clients    map[*Client]bool
		Broadcast  chan Message
		Players    int
	}

	func NewPool() *Pool {
		return &Pool{
			Register:   make(chan *Client),
			Unregister: make(chan *Client),
			Clients:    make(map[*Client]bool),
			Broadcast:  make(chan Message),
		}
	}

	func (p *Pool) Start() {
		for {
			select {
			case client := <-p.Register:
				p.Players++
				p.Clients[client] = true

				// Send initial messages based on the number of players
				if p.Players == 1 {
					fmt.Println("PLAYERS", p.Players)
					// writeMessage(client, Message{Type: 1, Body: "test"})
					writeMessage(client, Message{Type: 1, Body: "player-1"})
				} else if p.Players == 2 {
					fmt.Println("PLAYERS", p.Players)
					writeMessage(client, Message{Type: 1, Body: "player-2"})
				}

				fmt.Println("Size of Connection Pool: ", len(p.Clients))

				// Notify all clients about a new user
				for c := range p.Clients {
					writeMessage(c, Message{Type: 1, Body: "New User Joined"})

					// If there are two players, notify about readiness
					if len(p.Clients) == 2 {
						writeMessage(c, Message{Type: 1, Body: "Ready"})
					}
				}

			case client := <-p.Unregister:
				delete(p.Clients, client)
				fmt.Println("Size of Connection Pool: ", len(p.Clients))

				// Notify all remaining clients about the user who dropped
				for c := range p.Clients {
					writeMessage(c, Message{Type: 1, Body: "User Dropped"})
				}

			case message := <-p.Broadcast:
				fmt.Println("Sending message to all clients")
				// Broadcast the message to all clients
				for client := range p.Clients {
					writeMessage(client, message)
				}
			}
		}
	}

	func writeMessage(client *Client, message Message) {
		if err := client.Conn.WriteJSON(message); err != nil {
			fmt.Println(err)
			// Handle the error, e.g., unregister the client or take other appropriate actions.
		}
	}
