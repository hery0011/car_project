package ws

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Map pour stocker les connexions par commerçant
var clients = make(map[int]*websocket.Conn)
var mu sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // ⚠️ à sécuriser en prod
	},
}

// HandleWS : gérer la connexion WebSocket
// ws://localhost:8082/ws/commercant/{commercant_id}
func HandleWS(c *gin.Context) {
	commercantID := c.Param("id")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Erreur upgrade websocket:", err)
		return
	}

	// Stocker la connexion
	mu.Lock()
	var cid int
	fmt.Sscanf(commercantID, "%d", &cid)
	clients[cid] = conn
	mu.Unlock()

	fmt.Println("Commerçant connecté via WS:", cid)

	// Boucle de lecture (même si on ne reçoit rien)
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			mu.Lock()
			delete(clients, cid)
			mu.Unlock()
			fmt.Println("Commerçant déconnecté:", cid)
			break
		}
	}
}

// NotifyCommercant : envoyer un message à un commerçant connecté
func NotifyCommercant(commercantID int, message string) {
	mu.Lock()
	conn, ok := clients[commercantID]
	mu.Unlock()

	if ok {
		err := conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("Erreur envoi notif à commercant", commercantID, ":", err)
		} else {
			fmt.Println("Notification envoyée à commercant", commercantID, ":", message)
		}
	} else {
		fmt.Println("Commerçant non connecté :", commercantID)
	}
}
