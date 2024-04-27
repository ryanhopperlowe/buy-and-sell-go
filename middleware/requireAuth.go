package middleware

import (
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/ryanhopperlowe/buy-and-sell-go/initializers"
	"github.com/ryanhopperlowe/buy-and-sell-go/model"
)

func RequireAuth(ctx *gin.Context) {
	// Get token from cookie
	tokenString, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusAccepted, gin.H{"error": "No token provided"})
		return
	}

	// Decode/Validate Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrInvalidKeyType
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// pull claims from token
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		err = claims.Valid()

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: " + err.Error()})
			return
		}

		// Find user in database
		id, err := strconv.ParseUint(claims["sub"].(string), 10, 64)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token: invalid subject id"})
			return
		}

		var user model.User
		result := initializers.DB.First(&user, model.Model{ID: model.Identifier(id)})

		if result.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}

		// Set user in context
		ctx.Set("user", user)

		// Continue
		ctx.Next()
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid claim map"})
		return
	}
}
