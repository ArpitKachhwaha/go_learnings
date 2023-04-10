package main

import (
	"net/http"
	"time"

	"github.com/arpitkachhwaha/go-api/packages/common/db"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": time.Now()})
	})

	r.Run()
}
