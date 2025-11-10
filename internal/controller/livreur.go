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

	// userLat := geo.Lat
	// userLon := geo.Lon

	// --- Filtrer les livreurs dans un rayon de 1 km ---
	var livreursProches []entities.Livreur
	// for _, l := range livreurs {
	// 	distance := haversine(userLat, userLon, l.Latitude, l.Longitude)
	// 	if distance <= 25.0 { // rayon 1 km
	// 		livreursProches = append(livreursProches, l)
	// 	}
	// }

	// --- Retourner le JSON exactement comme demandé ---
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Liste des livreurs dans un rayon de 1 km",
		"data":    livreursProches,
	})
}

func (h *livraisonHandler) RegisterLivreur(c *gin.Context) {
	var payload struct {
		Name        string `json:"name"`
		Lastname    string `json:"lastname"`
		Mail        string `json:"mail"`
		Contact     string `json:"contact"`
		Password    string `json:"password"`
		Adresse     string `json:"adresse"`
		LivreurData struct {
			Nom           string `json:"nom"`
			Telephone     string `json:"telephone"`
			Vehicule      string `json:"vehicule"`
			ZoneLivraison string `json:"zone_livraison"`
		} `json:"livreur_data"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data", "error": err.Error()})
		return
	}

	tx := h.db.Begin()

	// 1️⃣ Créer utilisateur
	user := entities.LoginStruct{
		Login:    payload.Mail,
		Password: payload.Password,
		Name:     payload.Name,
		Lastname: payload.Lastname,
		Type:     "livreur",
		Contact:  payload.Contact,
		Mail:     payload.Mail,
		Adresse:  payload.Adresse,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}

	// 2️⃣ Assigner profil "livreur"
	var profil struct {
		IdProfil int `gorm:"column:idProfil"`
	}
	if err := tx.Table("profil").
		Select("idProfil").
		Where("nomProfil = ?", "livreur").
		First(&profil).Error; err != nil {

		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User created but 'livreur' profil not found"})
		return
	}

	if err := tx.Table("userprofil").Create(map[string]interface{}{
		"idUser":   user.Id,
		"idProfil": profil.IdProfil,
	}).Error; err != nil {

		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User created but failed to assign profile"})
		return
	}

	livreur := entities.Livreur{
		Nom:           payload.LivreurData.Nom,
		Telephone:     payload.LivreurData.Telephone,
		Vehicule:      payload.LivreurData.Vehicule,
		ZoneLivraison: payload.LivreurData.ZoneLivraison,
		UserID:        user.Id,
	}

	if err := tx.Table("livreur").Create(&livreur).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create livreur entry"})
		return
	}

	// 4️⃣ Commit
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"message":    "Livreur created successfully",
		"user_id":    user.Id,
		"livreur_id": livreur.Livreur_id,
	})
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
