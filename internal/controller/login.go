package controller

import (
	"car_project/internal/entities"
	"car_project/internal/jwt"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

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
	if err := h.db.Where("login = ?", loginData.Login).First(&user).Error; err != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Login successful",
		"data":    gin.H{"access_token": accessToken},
	})
}
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
