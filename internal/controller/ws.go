package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"remote/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Action    string  `json:"action"`
	DeltaX    float64 `json:"deltaX,omitempty"`
	DeltaY    float64 `json:"deltaY,omitempty"`
	Side      string  `json:"side,omitempty"`      // "left" / "right"
	Text      string  `json:"text,omitempty"`      // For keys
	Direction string  `json:"direction,omitempty"` // "up" / "down"
}

func Ws(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("failed to upgrade WebSocket: %v", err)
		return
	}
	defer conn.Close()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			log.Printf("websocket read error: %v", err)
			return
		}

		var msg Message
		if err := json.Unmarshal(data, &msg); err != nil {
			log.Printf("failed to unmarshal WebSocket message: %v", err)
			continue
		}

		switch msg.Action {
		case "move-cursor":
			if err := service.MoveCursor(msg.DeltaX, msg.DeltaY); err != nil {
				log.Printf("failed to move cursor: %v", err)
			}
		case "click":
			if err := service.Click(msg.Side == "left"); err != nil {
				log.Printf("failed mouse click: %v", err)
			}
		case "press-key":
			if err := service.PressKey(msg.Text); err != nil {
				log.Printf("failed key press: %v", err)
			}
		case "volume":
			if err := service.AdjustVolume(msg.Direction == "up"); err != nil {
				log.Printf("failed to adjust volume: %v", err)
			}
		case "toggle-mute":
			if err := service.ToggleMute(); err != nil {
				log.Printf("failed to toggle mute: %v", err)
			}
		}
	}
}
