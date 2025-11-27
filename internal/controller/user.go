package controller

import (
	"car_project/internal/entities"
	"car_project/internal/helper"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatUser handles the creation of a new user within a transactional context.
// It processes the incoming HTTP request, validates the input, and creates a user record in the database.
// If any error occurs during the process, the transaction is rolled back and an appropriate error response is returned.
// This handler expects a JSON payload representing the user to be created.
//
//	@Summary		Create a new user
//	@Description	Creates a new user in the system within a transactional context.
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		entities.LoginStruct	true	"User creation payload"
//	@Success		200		{object}	map[string]interface{}	"User created successfully"
//	@Failure		400		{object}	map[string]interface{}	"Invalid request data"
//	@Failure		500		{object}	map[string]interface{}	"Internal server error"
//	@Router			/admin/user/creatUser [post]
func (h *livraisonHandler) CreatUser(c *gin.Context) {
	var user entities.LoginStruct

	// Validate incoming JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request data",
		})
		return
	}

	user.Login = user.Mail

	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Database connection not initialized",
		})
		return
	}

	tx := h.db.Begin()

	// 1. Create user
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to create user",
		})
		return
	}

	// 2. Get the profil "client" dynamically
	var profil struct {
		IdProfil int `gorm:"column:idProfil"`
	}

	if err := tx.Table("profil").
		Select("idProfil").
		Where("nomProfil = ?", "client").
		First(&profil).Error; err != nil {

		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "User created but failed to assign default profile: 'client' profil not found",
		})
		return
	}

	// 3. Insert user-profil entry
	userProfil := map[string]interface{}{
		"idUser":   user.Id,
		"idProfil": profil.IdProfil,
	}

	if err := tx.Table("userprofil").Create(&userProfil).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "User created but failed to assign profile",
		})
		return
	}

	// 4. Commit
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to commit transaction",
		})
		return
	}

	go h.createWalletForUser(user.Id)

	// Response
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User created successfully and default 'client' profile assigned",
		"data":    user,
	})
}

func (h *livraisonHandler) createWalletForUser(userID int) {
	// Création de l'entité Wallet
	wallet := entities.Wallet{
		UserID:   userID,
		Balance:  1000000.00,
		Currency: "AR",
	}

	if err := h.db.Create(&wallet).Error; err != nil {
		log.Printf("ERROR: Failed to create wallet for UserID %d: %v", userID, err)
	} else {
		log.Printf("INFO: Wallet created successfully for UserID %d", userID)
	}
}

// DeleteUser supprime un utilisateur existant par son ID.
//
// @Summary      Supprimer un utilisateur
// @Description  Supprime un utilisateur en fonction de son identifiant (ID).
// @Tags         users
// @Param        idUser   path      int  true  "ID de l'utilisateur"
// @Produce      json
// @Success      200  {object}  map[string]interface{}  "Suppression réussie"
// @Failure      400  {object}  map[string]interface{}  "ID invalide"
// @Failure      200  {object}  map[string]interface{}  "Utilisateur introuvable"
// @Failure      500  {object}  map[string]interface{}  "Erreur interne du serveur"
// @Router       /admin/user/{idUser}/delete [delete]
func (h *livraisonHandler) DeleteUser(c *gin.Context) {
	var userEntite entities.LoginStruct
	var idUser = c.Param("idUser")
	id, err := strconv.Atoi(idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID invalide",
		})
		return
	}

	if err := h.db.First(&userEntite, id).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Utilisateur introuvable",
		})
		return
	}

	if err := h.db.Where("id = ?", id).Delete(&userEntite).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "suppression user avec succès",
		"data":    userEntite,
	})
	fmt.Println(idUser)
}

// UpdateUser godoc
// @Summary Met à jour un utilisateur
// @Description Met à jour les informations d'un utilisateur existant
// @Tags users
// @Accept json
// @Produce json
// @Param payload body entities.LoginStruct true "Utilisateur à mettre à jour"
// @Success 200 {object} map[string]interface{} "Utilisateur mis à jour avec succès"
// @Failure 400 {object} map[string]interface{} "Erreur lors de la mise à jour"
// @Router /admin/user/updateUser [put]
func (h *livraisonHandler) UpdateUser(c *gin.Context) {
	var payload entities.LoginStruct

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid request data",
		})
		return
	}

	if payload.Id == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ID utilisateur manquant",
		})
		return
	}

	if err := h.db.Where("id = ?", payload.Id).Updates(payload).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Utilisateur mis à jour avec succès",
		"data":    payload,
	})
}

func (h *livraisonHandler) GetUserMenu(c *gin.Context) {
	// Récupérer l'ID de l'utilisateur depuis la session
	userID, err := helper.GetUserID(c)
	if err != nil {
		// erreur déjà gérée dans GetUserID
		return
	}

	// Récupérer le rôle/profil de l'utilisateur depuis la table userprofil et profil
	var role string
	err = h.db.Table("userprofil").
		Select("profil.nomProfil").
		Joins("join profil on userprofil.idProfil = profil.idProfil").
		Where("userprofil.idUser = ?", userID).
		Limit(1). // si un user peut avoir plusieurs profils, on prend le premier
		Scan(&role).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Impossible de récupérer le profil de l'utilisateur"})
		return
	}

	// Récupérer les menus associés au rôle
	var menus []struct {
		Label string `json:"label"`
		Icon  string `json:"icon"`
		Link  string `json:"link"`
	}

	err = h.db.Table("menu").
		Select("menu.label, menu.icon, menu.link").
		Joins("join menu_roles on menu.id = menu_roles.menu_id").
		Where("menu_roles.role = ?", role).
		Scan(&menus).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Impossible de récupérer le menu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   menus,
	})
}

func (h *livraisonHandler) GetProfile(c *gin.Context) {
	// 👀 Ici on simule que userID = 1 pour exemple
	var user entities.User
	if err := h.db.First(&user, 1).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Utilisateur non trouvé"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": user})
}
