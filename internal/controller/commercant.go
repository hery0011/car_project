package controller

import (
	"car_project/internal/entities"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

const marchantProfilID = 2 // idProfil pour les commerçants

// Haversine pour calculer la distance en km
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Rayon de la Terre en km
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLon := (lon2 - lon1) * math.Pi / 180.0

	lat1 = lat1 * math.Pi / 180.0
	lat2 = lat2 * math.Pi / 180.0

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}

// Cherche les commerçants proches (1 km max)
func (h *livraisonHandler) ChercheCommercant(c *gin.Context) {
	var payload entities.PayloadCommercant
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	// Sélection avec jointure
	var commercants []entities.LoginStruct
	if err := h.db.Table("user").
		Select("user.*").
		Joins("join userProfil on user.id = userProfil.idUser").
		Where("userProfil.idProfil = ?", marchantProfilID).
		Find(&commercants).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}

	// Filtrer par distance
	var proches []entities.LoginStruct
	for _, cmt := range commercants {
		dist := haversine(payload.Latitude, payload.Longitude, cmt.Latitude, cmt.Longitude)
		if dist <= 1.0 { // rayon de 1 km
			proches = append(proches, cmt)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"client_position": payload,
		"commercants":     proches,
	})
}
