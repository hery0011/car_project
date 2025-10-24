package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID       int
	Conn     *websocket.Conn
	UserType string // client, merchant, admin
}

type Hub struct {
	Clients map[int]*Client
	Mutex   sync.RWMutex
}

var WSHub = Hub{
	Clients: make(map[int]*Client),
}
