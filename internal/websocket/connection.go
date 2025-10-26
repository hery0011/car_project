package websocket

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleWS(c *gin.Context) {
	userType := c.Param("userType")
	userID, _ := strconv.Atoi(c.Param("userID"))

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("WS Upgrade Error:", err)
		return
	}

	client := &Client{ID: userID, Conn: conn, UserType: userType}
	WSHub.Mutex.Lock()
	WSHub.Clients[userID] = client
	WSHub.Mutex.Unlock()

	fmt.Printf("%s %d connecté via WS\n", userType, userID)

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			break
		}
	}

	// Déconnecter
	WSHub.Mutex.Lock()
	delete(WSHub.Clients, userID)
	WSHub.Mutex.Unlock()
	fmt.Printf("%s %d déconnecté\n", userType, userID)
}
