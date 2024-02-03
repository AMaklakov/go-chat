package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

var state *State = &State{
	rooms: make(map[string]*Room),
	lock:  sync.RWMutex{},
}

type State struct {
	rooms map[string]*Room
	lock  sync.RWMutex
}

func (s *State) AddRoom(id string) *Room {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.rooms[id] = &Room{
		ID:       id,
		Clients:  make(map[net.Conn]*Client),
		messages: make(chan []byte, 10),
		lock:     sync.RWMutex{},
	}
	return s.rooms[id]
}

func (s *State) TryDelRoom(id string) {
	s.lock.RLock()
	room, ok := s.rooms[id]
	s.lock.RUnlock()

	if ok && len(room.Clients) == 0 {
		s.delRoom(id)
	}
}

func (s *State) delRoom(id string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if room, ok := s.rooms[id]; ok {
		room.Destroy()
		delete(s.rooms, id)
		log.Printf("Room %s was deleted!", id)
	}
}

func (s *State) Room(id string) *Room {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.rooms[id]
}

func (s *State) GetOrAddRoom(id string) *Room {
	if room := s.Room(id); room != nil {
		return room
	}

	room := state.AddRoom(id)
	go func() {
		if err := room.ServeMessages(); err != nil {
			fmt.Println("Error serving messages: ", err)
		}
	}()
	return room
}
