package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool)
var clientsNames = make(map[*websocket.Conn]string)
var broadcast = make(chan Message)

func handleWsConnections(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)

	user := DBloginUser(r.Header.Get("name"), r.Header.Get("pass"))

	if user == nil {
		conn.WriteJSON("Invalid name/password")
		conn.Close()
	}

	if err != nil {
		fmt.Println(err, "ERROR")
		return
	}
	defer conn.Close()

	clients[conn] = true
	clientsNames[conn] = user.Name

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		msg.UserFromName = clientsNames[conn]
		if err != nil {
			fmt.Println(err)
			delete(clients, conn)
			delete(clientsNames, conn)
			return
		}

		broadcast <- msg
	}

}

func handleMessages() {
	for {
		msg := <-broadcast

		iterator := 0
		for client := range clients {

			if clientsNames[client] != msg.UserNameTo {
				iterator++
			} else {
				err := client.WriteJSON(msg)

				if err != nil {
					fmt.Println(err)
					client.Close()
					delete(clients, client)
					delete(clientsNames, client)
				}
			}

		}
	}
}
