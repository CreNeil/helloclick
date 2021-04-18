package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var ApiToken = "Authorized"

type Hello struct {
	Name    string `form:"name" `
	Address string `form:"address"`
}

func helloGet(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	//stand for c.Request.URL.Query().Get("lastname")
	lastname := c.Query("lastname")
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}

func helloGetPath(c *gin.Context) {
	path := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"name": path})
}

func helloPost(c *gin.Context) {
	// get post parameters from form
	//message := c.Request.Form.Get("message")
	var hello Hello

	if err := c.ShouldBindJSON(&hello); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":    hello.Name,
		"address": hello.Address})
}

func helloPut(c *gin.Context) {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick", "anonymous")
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"nick":    nick})
}

func helloDelete(c *gin.Context) {
	deleteId := c.Request.Form.Get("id")
	c.JSON(http.StatusOK, gin.H{"message": "delete" + deleteId})
}

func helloPatch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func helloHead(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func helloOptions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func middleWare() gin.HandlerFunc {
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

func main() {
	//default mode , include Logger Recovery middle ware
	r := gin.Default()
	r.Use(middleWare())

	r.GET("/get", helloGet)
	r.GET("/get/:name", helloGetPath)
	r.POST("/post", helloPost)
	r.PUT("/put", helloPut)
	r.DELETE("/delete", helloDelete)
	r.PATCH("/patch", helloPatch)
	r.HEAD("/head", helloHead)
	r.OPTIONS("/options", helloOptions)

	_ = r.Run(":2466")
}
