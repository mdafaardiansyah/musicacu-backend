package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mdafaardiansyah/musicacu-backend/internal/configs"
	"github.com/mdafaardiansyah/musicacu-backend/pkg/jwt"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Service.SecretKey
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")

		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
			return
		}

		userID, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userID", userID)
		c.Set("username", username)
		c.Next()
	}
}

//func AuthRefreshMiddleware() gin.HandlerFunc {
//	secretKey := configs.Get().Service.SecretKey
//	return func(c *gin.Context) {
//		header := c.Request.Header.Get("Authorization")
//
//		header = strings.TrimSpace(header)
//		if header == "" {
//			c.AbortWithError(http.StatusUnauthorized, errors.New("missing token"))
//			return
//		}
//
//		userID, username, err := jwt.ValidateTokenWithoutExpiry(header, secretKey)
//		if err != nil {
//			c.AbortWithError(http.StatusUnauthorized, err)
//			return
//		}
//		c.Set("userID", userID)
//		c.Set("username", username)
//		c.Next()
//	}
//
//}
