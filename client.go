package main

import (
	"fmt"
	"net"
	"strings"
)

type Client struct {
	Name string
	conn net.Conn
}

func (c *Client) Read() ([]byte, error) {
	bytes := make([]byte, 1024)
	n, err := c.conn.Read(bytes)
	if err != nil {
		return []byte{}, err
	}
	return []byte(strings.TrimSpace(string(bytes[:n]))), nil
}

func (c *Client) ReadLoop(ch chan []byte) error {
	for {
		msg, err := c.Read()
		if err != nil {
			return err
		}
		ch <- []byte(fmt.Sprintf("%s (%s): %s", c.Name, c.conn.RemoteAddr(), string(msg)))
	}
}

func (c *Client) Write(message []byte) error {
	// now := time.Now().Format(time.DateTime)
	// msg := now + " -- " + string(message) + "\n"
	msg := "💬 " + string(message) + "\n"
	_, err := c.conn.Write([]byte(msg))
	return err
}
