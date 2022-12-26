package middleware

import (
	"errors"
	"github/criotech/resturant-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			res := utils.NewHTTPResponse(http.StatusUnauthorized, errors.New("auth failed"))
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(clientToken)
		if err != nil {
			res := utils.NewHTTPResponse(http.StatusUnauthorized, err)
			c.JSON(http.StatusUnauthorized, res)
			c.Abort()
			return
		}
		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_name)
		c.Set("last_name", claims.Last_name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)
		c.Next()
	}
}
