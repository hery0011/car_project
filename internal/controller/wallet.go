package controller

import (
	"car_project/internal/entities"
	"car_project/internal/helper"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *livraisonHandler) GetMontantWallet(c *gin.Context) {
	userID, err := helper.GetUserID(c)
	if err != nil {
		// l'erreur a déjà été gérée dans GetUserID, on stoppe le handler
		return
	}

	// 1. Déclarer une variable pour stocker le Wallet
	var wallet entities.Wallet

	// 2. Interroger la base de données
	// Assurez-vous que 'h.DB' est bien votre instance de *gorm.DB
	result := h.db.Where("user_id = ?", userID).First(&wallet)

	// 3. Gérer les erreurs de la base de données
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Aucun portefeuille trouvé pour cet utilisateur
			c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found for this user"})
			return
		}
		// Autres erreurs de base de données
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error while fetching wallet"})
		return
	}

	// 4. Retourner le solde (Balance)
	c.JSON(http.StatusOK, gin.H{
		"user_id":  userID,
		"balance":  wallet.Balance,
		"currency": wallet.Currency,
	})
}
