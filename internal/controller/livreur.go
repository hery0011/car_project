package controller

import (
	"car_project/internal/entities"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Structure pour réponse de l'API de géolocalisation
type GeoResponse struct {
	Query      string  `json:"query"`
	Country    string  `json:"country"`
	City       string  `json:"city"`
	RegionName string  `json:"regionName"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Status     string  `json:"status"`
	Message    string  `json:"message"`
}

func (h *livraisonHandler) GetLocation(c *gin.Context) {
	var livreurs []entities.Livreur

	// Récupérer tous les livreurs
	if err := h.db.Find(&livreurs).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if len(livreurs) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Aucun livreur trouvé",
			"data":    []entities.Livreur{},
		})
		return
	}

	// --- Récupérer position du client via IP ---
	clientIP := c.ClientIP()
	if clientIP == "::1" || clientIP == "127.0.0.1" {
		clientIP = "8.8.8.8" // test local
	}

	resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", clientIP))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible d'obtenir la localisation"})
		return
	}
	defer resp.Body.Close()

	var geo GeoResponse
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &geo); err != nil || geo.Status != "success" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossible de récupérer la position"})
		return
	}

	userLat := geo.Lat
	userLon := geo.Lon

	// --- Filtrer les livreurs dans un rayon de 1 km ---
	var livreursProches []entities.Livreur
	for _, l := range livreurs {
		distance := haversine(userLat, userLon, l.Latitude, l.Longitude)
		if distance <= 25.0 { // rayon 1 km
			livreursProches = append(livreursProches, l)
		}
	}

	// --- Retourner le JSON exactement comme demandé ---
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Liste des livreurs dans un rayon de 1 km",
		"data":    livreursProches,
	})
}

func (h *livraisonHandler) AjoutLivreur(c *gin.Context) {

}

func (h *livraisonHandler) ListLivreur(c *gin.Context) {
	var livreur []entities.Livreur

	if err := h.db.Find(&livreur).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if len(livreur) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Aucune livreur trouvée",
			"data":    []entities.Categorie{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Liste des livreurs",
		"data":    livreur,
	})
}
