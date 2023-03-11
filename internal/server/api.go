package server

import (
	"github.com/gin-gonic/gin"
	"github.com/pawlh/unero/internal/dal/mariadb"
)

func registerRoutes(router *gin.Engine) {
	router.POST("/api/v1/transaction", createTransaction)
}

func createTransaction(c *gin.Context) {
	// get the transaction from the request body

	transaction := mariadb.Transaction{}

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	db, err := mariadb.OpenDB()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	transaction, err = db.Create(&transaction)

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
