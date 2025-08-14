package jwt

import (
	"car_project/internal/config"
	"car_project/internal/entities"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	SessionData entities.SessionData
	jwt.StandardClaims
}

// Générer un access token
func GenerateAccessToken(sessionData entities.SessionData) (string, error) {
	key := []byte(config.AppConfiguration.JWT.SecretKey)

	tokenExpires := getTokenExpire(config.AppConfiguration.JWT.TokenExpire, false)
	claims := &Claims{
		SessionData: sessionData,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpires).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Valider un access token
func ValidateAccessToken(tokenString string) (*Claims, error) {
	key := []byte(config.AppConfiguration.JWT.SecretKey)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || time.Now().Unix() > claims.ExpiresAt {
		return nil, errors.New("token expired")
	}

	return claims, nil
}

// Refresh access token
func RefreshAccessToken(tokenString string) (string, int, error) {
	claims, err := ValidateAccessToken(tokenString)
	if err != nil {
		return "", 401, err
	}

	newToken, err := GenerateAccessToken(claims.SessionData)
	if err != nil {
		return "", 500, err
	}

	return newToken, 200, nil
}

// Obtenir la durée du token
func getTokenExpire(tokenExpire string, refresh bool) time.Duration {
	// format: "15m", "2h", "1d"
	unit := tokenExpire[len(tokenExpire)-1:]
	val := tokenExpire[:len(tokenExpire)-1]

	duration := time.Minute * 15
	switch unit {
	case "m":
		if v, err := time.ParseDuration(val + "m"); err == nil {
			duration = v
		}
	case "h":
		if v, err := time.ParseDuration(val + "h"); err == nil {
			duration = v
		}
	case "d":
		if v, err := time.ParseDuration(val + "h"); err == nil {
			duration = v * 24
		}
	}

	return duration
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token manquant ou invalide"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := ValidateAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalide"})
			c.Abort()
			return
		}

		// Stocke les données de session dans le contexte Gin pour les handlers
		c.Set("sessionData", claims.SessionData)
		c.Next()
	}
}
