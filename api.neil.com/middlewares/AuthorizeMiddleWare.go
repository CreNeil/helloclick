package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ApiToken = "Authorized"

func AuthorizeMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("it is my middle Ware!")
		contentType := c.Request.Header.Get("Content-Type")
		ContentLength := c.Request.Header.Get("Content-Length")
		UserAgent := c.Request.Header.Get("User-Agent")
		token := c.Request.Header.Get("token")
		fmt.Println(contentType)
		fmt.Println(ContentLength)
		fmt.Println(UserAgent)
		fmt.Println(token)
		// find api tokens in redis
		// yum redis in vmware
		if token != ApiToken {
			c.JSON(http.StatusNonAuthoritativeInfo,
				gin.H{"message": "no authorize"})
			c.Abort()
		}
	}
}
