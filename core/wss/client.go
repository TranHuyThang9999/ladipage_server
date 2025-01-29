package wss

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func (m *WebSocketManager) removeChannelClient(client *Client, channelID int64) {
	m.clientsMu.Lock()
	defer m.clientsMu.Unlock()

	if channelClients, exists := m.channels[channelID]; exists {
		delete(channelClients, client.ClientID)
		if len(channelClients) == 0 {
			delete(m.channels, channelID)
		}
	}
}

func (m *WebSocketManager) registerClient(client *Client) {
	m.clientsMu.Lock()
	defer m.clientsMu.Unlock()

	if m.clients[client.UserID] == nil {
		m.clients[client.UserID] = make(map[int64]*Client)
	}
	m.clients[client.UserID][client.ClientID] = client
}
func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		clients:  make(map[int64]map[int64]*Client),
		channels: make(map[int64]map[int64]*Client),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (m *WebSocketManager) removeClient(client *Client) {
	m.clientsMu.Lock()
	defer m.clientsMu.Unlock()

	delete(m.clients[client.UserID], client.ClientID)
	if len(m.clients[client.UserID]) == 0 {
		delete(m.clients, client.UserID)
	}
}
