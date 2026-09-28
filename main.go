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

func getHTML(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>HTML test</title>
</head>
<body>
	<h1>HTML test page</h1>
	<p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.</p>
</body>
</html>`))
}

const Constant string = "Go Constant"

func main() {
	fmt.Println("go build main.go")

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
    router.GET("/html", getHTML)
	router.Run("0.0.0.0:3000")
}
