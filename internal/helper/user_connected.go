package helper

import (
	"net/http"

	"car_project/internal/entities"

	"github.com/gin-gonic/gin"
)

// GetUserID récupère l'ID de l'utilisateur depuis le contexte Gin.
// Retourne un int et une erreur si l'utilisateur n'est pas authentifié.
func GetUserID(c *gin.Context) (int, error) {
	sessionInterface, exists := c.Get("sessionData")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Utilisateur non authentifié"})
		return 0, http.ErrNoCookie // ou une autre erreur personnalisée
	}

	sessionData, ok := sessionInterface.(entities.SessionData)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer la session utilisateur"})
		return 0, http.ErrNoCookie // ou une autre erreur personnalisée
	}

	return sessionData.User.Id, nil
}
