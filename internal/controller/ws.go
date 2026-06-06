package controller

import (
	"encoding/json"
	"fmt"
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
		fmt.Println(err)
		return
	}
	defer conn.Close()

	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		var msg Message
		if err := json.Unmarshal(data, &msg); err != nil {
			fmt.Println(err)
			continue
		}

		switch msg.Action {
		case "move-cursor":
			_ = service.MoveCursor(msg.DeltaX, msg.DeltaY)
		case "click":
			_ = service.Click(msg.Side == "left")
		case "press-key":
			_ = service.PressKey(msg.Text)
		case "volume":
			_ = service.AdjustVolume(msg.Direction == "up")
		case "toggle-mute":
			_ = service.ToggleMute()
		}
	}
}
