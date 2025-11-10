package controller

import (
	"car_project/internal/entities"
	"log"
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

func (h *livraisonHandler) RegisterCommercant(c *gin.Context) {
	var payload struct {
		Name           string `json:"name"`
		Lastname       string `json:"lastname"`
		Mail           string `json:"mail"`
		Contact        string `json:"contact"`
		Password       string `json:"password"`
		Adresse        string `json:"adresse"`
		CommercantData struct {
			Nom         string `json:"nom"`
			Adresse     string `json:"adresse"`
			Telephone   string `json:"telephone"`
			Email       string `json:"email"`
			Description string `json:"description"`
		} `json:"commercant_data"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	tx := h.db.Begin()

	// 1️⃣ Créer utilisateur
	user := entities.LoginStruct{
		Login:    payload.Mail,
		Password: payload.Password,
		Name:     payload.Name,
		Lastname: payload.Lastname,
		Type:     "commercant",
		Contact:  payload.Contact,
		Mail:     payload.Mail,
		Adresse:  payload.Adresse,
	}

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		log.Printf("[ERROR] Failed to create user. Data: %+v | DB Error: %v", user, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create user"})
		return
	}

	// 2️⃣ Assigner profil "commercant"
	var profil struct {
		IdProfil int `gorm:"column:idProfil"`
	}
	if err := tx.Table("profil").
		Select("idProfil").
		Where("nomProfil = ?", "commercant").
		First(&profil).Error; err != nil {

		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "User created but 'commercant' profil not found"})
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

	// 3️⃣ Créer commerçant et récupérer ID
	type Commercant struct {
		CommercantID int    `gorm:"primaryKey;column:commercant_id"`
		Nom          string `gorm:"column:nom"`
		Adresse      string `gorm:"column:adresse"`
		Telephone    string `gorm:"column:telephone"`
		Email        string `gorm:"column:email"`
		Description  string `gorm:"column:description"`
	}

	commercant := Commercant{
		Nom:         payload.CommercantData.Nom,
		Adresse:     payload.CommercantData.Adresse,
		Telephone:   payload.CommercantData.Telephone,
		Email:       payload.CommercantData.Email,
		Description: payload.CommercantData.Description,
	}

	if err := tx.Table("commercant").Create(&commercant).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create commercant entry"})
		return
	}

	// 4️⃣ Mettre à jour user avec commercant_id
	if err := tx.Model(&user).Update("commercant_id", commercant.CommercantID).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to link user to commercant"})
		return
	}

	// 5️⃣ Commit (UN SEUL)
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to commit transaction"})
		return
	}

	// 6️⃣ (Optionnel) créer wallet
	go h.createWalletForUser(user.Id)

	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusOK,
		"message":       "Commercant created successfully",
		"user_id":       user.Id,
		"commercant_id": commercant.CommercantID,
	})
}
