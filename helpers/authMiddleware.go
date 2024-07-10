package helpers

import (
	"net/http"
	"practice/restfulApi/global"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		// get the authorization key value from header
		authorization := c.Request.Header["Authorization"][0]
		if authorization == "" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "you forgot to add authorization header",
			})
			c.Abort()
			return
		}

		global.Logger.Print("the value is ", authorization)
		global.Logger.Print("the value is ")
		token := strings.Split(authorization, " ")[1]
		// token := "vlaue is here "

		global.Logger.Print("token the value is ", token)
		claims, err := ParseAccessToken(token)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error":  "Error occured on token",
				"detail": err.Error(),
			})
			c.Abort()
			return
		}
		if claims.Valid() != nil {

			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "claims error",
			})
			c.Abort()
			return
		}
		if claims.Type != "access" {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{
				"error": "You probably provide wrong token",
			})
			c.Abort()
			return
		}

		global.Logger.Print("the value is ", claims.Name)
		// validate json token from authorization key
		//

		c.Next()
	}
}
