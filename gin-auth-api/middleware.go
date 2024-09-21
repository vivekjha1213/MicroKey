package main

import (
    "encoding/base64"
    "github.com/gin-gonic/gin"
    "github.com/go-resty/resty/v2"
    "net/http"
    "strings"
)

// AuthMiddleware checks for Basic Auth and calls the Django service
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        auth := c.Request.Header.Get("Authorization")
        if auth == "" || !strings.HasPrefix(auth, "Basic ") {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        payload := strings.TrimPrefix(auth, "Basic ")
        decoded, err := base64.StdEncoding.DecodeString(payload)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header"})
            c.Abort()
            return
        }

        parts := strings.SplitN(string(decoded), ":", 2)
        if len(parts) != 2 {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
            c.Abort()
            return
        }

        username := parts[0]
        password := parts[1]

        // Call the Django authentication endpoint
        client := resty.New()
        var userDetails map[string]interface{}
        resp, err := client.R().
            SetBasicAuth(username, password).
            SetResult(&userDetails).
            Get("http://127.0.0.1:8000/auth/service/third_party")

        if err != nil || resp.StatusCode() != http.StatusOK {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
            c.Abort()
            return
        }

        c.Set("user", userDetails)
        c.Next()
    }
}
