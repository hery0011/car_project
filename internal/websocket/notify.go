package websocket

import "car_project/internal/entities"

type Notification struct {
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NotifyClient(userID int, commande *entities.Order) {
	WSHub.Mutex.RLock()
	defer WSHub.Mutex.RUnlock()
	if c, ok := WSHub.Clients[userID]; ok && c.UserType == "client" {
		c.Conn.WriteJSON(Notification{"commande", "Votre commande a été enregistrée", commande})
	}
}

func NotifyMerchant(merchantID int, commande *entities.Order) {
	WSHub.Mutex.RLock()
	defer WSHub.Mutex.RUnlock()
	if c, ok := WSHub.Clients[merchantID]; ok && c.UserType == "merchant" {
		c.Conn.WriteJSON(Notification{"commande", "Nouvelle commande pour vos produits", commande})
	}
}

func NotifyAdmin(commande *entities.Order) {
	WSHub.Mutex.RLock()
	defer WSHub.Mutex.RUnlock()
	for _, c := range WSHub.Clients {
		if c.UserType == "admin" {
			c.Conn.WriteJSON(Notification{"commande", "Nouvelle commande enregistrée", commande})
		}
	}
}
