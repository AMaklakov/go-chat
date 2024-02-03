package main

import (
	"log"
	"net"
	"sync"
)

type Room struct {
	ID       string
	Clients  map[net.Conn]*Client
	messages chan []byte
	lock     sync.RWMutex
}

func (r *Room) AddClient(client *Client) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.Clients[client.conn] = client
	log.Printf("Client %s added to room %s\n", client.Name, r.ID)
}

func (r *Room) DelClient(conn net.Conn) {
	r.lock.Lock()
	defer r.lock.Unlock()
	c := r.Clients[conn]
	delete(r.Clients, conn)
	log.Printf("Client %s removed from room %s\n", c.Name, r.ID)
}

func (r *Room) ServeMessages() error {
	log.Printf("Serving messages for room %s\n", r.ID)
	defer log.Printf("Ended serving messages for room %s\n", r.ID)
	for message := range r.messages {
		for _, client := range r.Clients {
			if err := client.Write(message); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *Room) Destroy() {
	close(r.messages)
}
