package users

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"

	"github.com/kwamekyeimonies/sb-design-bk/middleware"
	"github.com/kwamekyeimonies/sb-design-bk/models"
)

func (userControl *UserController) LoginUserController(ctx *gin.Context) {
	if ctx.Request.ContentLength == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid/empty body request",
		})
		return
	}

	var usr *models.User

	if err := ctx.ShouldBindJSON(&usr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
	}

	if usr.Email == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email/password"})
	}

	if usr.Password == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email/password"})
	}

	response, err := userControl.userCntrl.LoginUserAccount(usr, ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	expirationTime := time.Now().Add(99999 * time.Hour)
	claims := &middleware.JWTClaim{
		APIKey: response.ApiKey,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.SecretKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		ctx.Abort()
		return
	}

	cookiesExpiration := time.Now().Add(7 * 24 * time.Hour).Unix()
	ctx.Header("Authorization", "Bearer "+tokenString)
	ctx.SetCookie("token", tokenString, int(cookiesExpiration-time.Now().Unix()), "/", "", false, false)

	ctx.JSON(200, gin.H{
		"message":  "Login successful",
		"response": response,
		"token":    tokenString,
	})

}
