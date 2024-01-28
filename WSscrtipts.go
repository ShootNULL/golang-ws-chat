package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool)
var clientsNames = []string{}
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
	clientsNames = append(clientsNames, user.Name)

	for {
		var msg Message
		err := conn.ReadJSON(&msg)

		if err != nil {
			fmt.Println(err)
			delete(clients, conn)

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

			if clientsNames[iterator] != msg.UserNameTo {
				iterator++
			} else {
				err := client.WriteJSON(msg.Message)
				//log.Print(clientsNames[iterator], "", msg.UserNameTo, "", clientsNames)

				if err != nil {
					fmt.Println(err)
					client.Close()
					delete(clients, client)
				}
			}

		}
	}
}
