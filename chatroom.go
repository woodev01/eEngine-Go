package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

type building struct {
	roomSize    int
	rooms       []chatroom
	currentRoom chatroom
}

type chatroom struct {
	name     string
	roomType string
	sockets  []*websocket.Conn
}

func serveRoom(port, name, roomType string) {
	var office building

	var initialSlice = make([]chatroom)
	office.rooms = initialSlice

	newMux := http.NewServeMux()
	newMux.Handle("/", websocket.Handler(room.sockHandler))

	log.Fatal(http.ListenAndServe(port, newMux))
}

func (b *building) makeRoom(name, roomType string) {
	var room chatroom

	var initialSlice = make([]*websocket.Conn, 1)
	room.name = name
	room.roomType = roomType
	room.sockets = initialSlice

	b.rooms = append(b.rooms, room)
}

func (c *chatroom) sockHandler(ws *websocket.Conn) {
	var err error

	c.sockets = append(c.sockets, ws)
	fmt.Println("Sockets: ", c.name, c.sockets)
	websocket.Message.Send(ws, "You've Joined the Chatroom")

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can't receive")
			break
		}

		fmt.Println("Received back from client: " + reply)

		msg := reply
		fmt.Println("Sending to client: " + msg)

		for i := 1; i < len(c.sockets); i++ {
			if err = websocket.Message.Send(c.sockets[i], msg); err != nil {
				fmt.Println("Can't send")
				break
			}
		}
	}
}
