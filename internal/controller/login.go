package controller

import (
	"car_project/internal/entities"
	"car_project/internal/jwt"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Login authenticates a user based on provided login credentials.
// If the credentials are valid, returns a JWT access token.
//
//	@Summary		User login
//	@Description	Authenticates a user using their login and password. Returns a JWT access token if successful.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			loginData	body		entities.LoginData	true	"Login payload"
//	@Success		200	{object}	map[string]interface{}	"Login successful"
//	@Failure		400	{object}	map[string]interface{}	"Invalid login data or missing fields"
//	@Failure		401	{object}	map[string]interface{}	"Invalid login credentials"
//	@Failure		500	{object}	map[string]interface{}	"Internal server error"
//	@Router			/auth/login [post]
func (h *livraisonHandler) Login(c *gin.Context) {
	var loginData entities.LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		fmt.Printf("Error parsing request body: %s\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid login data",
		})
		return
	}

	if loginData.Login == "" || loginData.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Login and password are required",
		})
		return
	}

	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Database connection not initialized",
		})
		return
	}

	var user entities.LoginStruct
	if err := h.db.Where("mail = ?", loginData.Login).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Invalid login credentials",
		})
		return
	}

	accessToken, err := jwt.GenerateAccessToken(entities.SessionData{User: entities.LoginStruct(user)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Failed to generate access token",
		})
		return
	}

	userResp := entities.UserResponse{
		Id:        user.Id,
		Login:     user.Login,
		Name:      user.Name,
		Lastname:  user.Lastname,
		Type:      user.Type,
		Contact:   user.Contact,
		Mail:      user.Mail,
		Adresse:   user.Adresse,
		Latitude:  user.Latitude,
		Longitude: user.Longitude,
	}

	response := entities.LoginResponse{
		AccessToken: accessToken,
		User:        userResp,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Login successful",
		"data":    response,
	})
}

// Logout handles user logout by validating the provided Bearer token.
// If the token is valid, the user is considered logged out (client-side token should be discarded).
//
//	@Summary		Logout a user
//	@Description	Validates the Bearer token and logs the user out. The token should be removed on the client side.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer access token"
//	@Success		200	{object}	map[string]interface{}	"Déconnexion réussie"
//	@Failure		401	{object}	map[string]interface{}	"Token manquant ou invalide"
//	@Router			/auth/logout [post]
func (h *livraisonHandler) Logout(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Token manquant ou invalide",
		})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	_, err := jwt.ValidateAccessToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Token invalide",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Déconnexion réussie",
	})
}

// Refresh generates a new access token using a valid refresh token.
// The request must include a Bearer token in the Authorization header.
//
//	@Summary		Refresh access token
//	@Description	Generates a new access token if the provided Bearer token is valid.
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			Authorization	header		string	true	"Bearer refresh token"
//	@Success		200	{object}	map[string]interface{}	"Access token refreshed successfully"
//	@Failure		401	{object}	map[string]interface{}	"Token manquant ou invalide"
//	@Failure		500	{object}	map[string]interface{}	"Internal server error"
//	@Router			/auth/refresh [post]
func (h *livraisonHandler) Refresh(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Token manquant ou invalide",
		})
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	newAccessToken, code, err := jwt.RefreshAccessToken(tokenString)
	if err != nil {
		c.JSON(code, gin.H{
			"status":  "failed",
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(code, gin.H{
		"status":  "success",
		"message": "Access token refreshed successfully",
		"data":    gin.H{"access_token": newAccessToken},
	})
}
