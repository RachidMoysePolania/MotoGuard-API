package middleware

import (
	"net/http"
	"time"

	"github.com/RachidMoysePolania/MotoGuard-API/controllers"
	"github.com/RachidMoysePolania/MotoGuard-API/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(c *gin.Context) {
	tockenstring, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := jwt.Parse(tockenstring, func(token *jwt.Token) (interface{}, error) {
		return []byte("password-for-jwt"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.Userdata
		controllers.Db.Where("email = ?", claims["user"]).First(&user)

		c.Set("user", user)
		c.Next()
	}
}
