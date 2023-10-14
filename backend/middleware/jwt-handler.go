package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/kwamekyeimonies/sb-design-bk/config"
)

type JWTClaim struct {
	jwt.RegisteredClaims
	APIKey       string `json:"api_key"`
	RefreshToken string `json:"refresh_token"`
}

func LoadJWTKeys(path string) string {
	keys, _ := config.LoadInitializer("credentials")
	return keys.JWT_SECRET_KEY
}

var SecretKey = []byte(LoadJWTKeys("credentials"))

func JWTMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			ctx.Abort()
			return
		}
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			ctx.Abort()
			return
		}

		tokenString := authHeaderParts[1]
		claims := &JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return SecretKey, nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
