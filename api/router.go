package api

import (
    "github.com/gin-gonic/gin"
    "healthcare-blockchain/auth"
    "healthcare-blockchain/blockchain"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    auth.RegisterRoutes(r)
    blockchain.RegisterRoutes(r)

    r.GET("/healthcheck", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy"})
    })

    return r
}
