package websocket

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

// Hub manages all connected WebSocket clients and broadcasts messages.
type Hub struct {
	mu      sync.Mutex
	clients map[*websocket.Conn]bool
}

var DefaultHub = &Hub{clients: make(map[*websocket.Conn]bool)}

// HandleConnection registers a WebSocket client and keeps the connection alive.
func (h *Hub) HandleConnection(c *websocket.Conn) {
	h.mu.Lock()
	h.clients[c] = true
	h.mu.Unlock()

	log.Printf("WebSocket client connected (%d total)", len(h.clients))

	// Keep connection alive; on disconnect, remove the client.
	for {
		if _, _, err := c.ReadMessage(); err != nil {
			h.mu.Lock()
			delete(h.clients, c)
			h.mu.Unlock()
			log.Printf("WebSocket client disconnected (%d total)", len(h.clients))
			return
		}
	}
}

// broadcast sends a message to all connected clients.
func (h *Hub) broadcast(message []byte) {
	h.mu.Lock()
	defer h.mu.Unlock()

	for conn := range h.clients {
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("WebSocket write error: %v", err)
			conn.Close()
			delete(h.clients, conn)
		}
	}
}

// NotifyChange broadcasts a "data_changed" event after any mutation.
func NotifyChange() {
	DefaultHub.broadcast([]byte(`{"event":"data_changed"}`))
}
