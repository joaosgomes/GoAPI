package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getGo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Gin API in Go")
}

func getRedirectGo(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://go.dev/")
}

func getStatusGo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"Status": "Ok",
	})
}

const Constant string = "Go Constant"

func main() {
	fmt.Println("go build go.go")

	var variable = "Go"
	fmt.Println(variable)

	fmt.Println(Constant)

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	router := gin.Default()
	router.GET("/", getGo)
	router.GET("/redirect", getRedirectGo)
	router.GET("/status", getStatusGo)

	router.Run("localhost:8080")
}
