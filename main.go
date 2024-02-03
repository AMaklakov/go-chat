package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection acceptance error:", err)
			continue
		}
		go func() {
			if err := HandleConnect(conn); err != nil {
				fmt.Println("Error: ", err)
			}
		}()
	}
}

func HandleConnect(conn net.Conn) error {
	client := &Client{Name: "", conn: conn}

	// Read client's name
	if err := client.Write([]byte("Please, enter your name: ")); err != nil {
		return err
	}
	if name, err := client.Read(); err == nil {
		client.Name = string(name)
		log.Println("User added: ", client.Name)
	} else {
		return err
	}

	// Reading Room id
	err := client.Write([]byte("Type in room number: "))
	if err != nil {
		fmt.Println("Error room number: ", err)
	}
	roomId, err := client.Read()
	if err != nil {
		fmt.Println("Error reading room id", err)
	}

	// Getting room
	room := state.GetOrAddRoom(string(roomId))
	room.AddClient(client)
	room.messages <- []byte(fmt.Sprintf("*** A new user \"%s\" joined the room \"%s\" ***", client.Name, room.ID))

	clients := []string{}
	for _, c := range room.Clients {
		clients = append(clients, c.Name)
	}
	room.messages <- []byte(fmt.Sprintf("*** Now in room: %s ***", strings.Join(clients, ", ")))

	// Reading messages
	if err := client.ReadLoop(room.messages); err != nil {
		if err != io.EOF {
			fmt.Println("Error in read loop: ", err)
		}
		room.DelClient(conn)
		room.messages <- []byte(fmt.Sprintf("*** User \"%s\" left the room \"%s\" ***", client.Name, room.ID))
		state.TryDelRoom(room.ID)
	}
	return nil
}
