package main

import (
    "github.com/gin-gonic/gin"
)
// DEPRECATED
func homeHandler(c *gin.Context) {
    c.String(200, "Welcome to our API!")
}

func helloHandler(c *gin.Context) {
    c.String(200, "Hello, World!")
}

func main() {
    router := gin.Default()

    router.GET("/", homeHandler)
    router.GET("/hello", helloHandler)

    router.Run(":8080")
}