package wss

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ClientID  int64
	UserID    int64
	ChannelID int64
	Conn      *websocket.Conn
	manager   *WebSocketManager
}

type WebSocketManager struct {
	clients   map[int64]map[int64]*Client
	channels  map[int64]map[int64]*Client
	clientsMu sync.Mutex
	upgrader  websocket.Upgrader
}
