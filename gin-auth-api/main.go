package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()

    router.GET("/auth/service/middleware", AuthMiddleware(), func(c *gin.Context) {
        userDetails, _ := c.Get("user") 
        c.JSON(http.StatusOK, userDetails)
    })

    router.Run(":9000")
}
