package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": Dbs,
	})
}

type Db struct {
	Text string `json:"text,omitempty"`
	// Deskripsi string `json:"deskripsi,omitempty"`
}

var Dbs []Db

type DataRequest struct {
	Text string `json:"text"`
}

func postHandler(c *gin.Context) {
	var newDb Db
	if err := c.ShouldBindJSON(&newDb); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Dbs = append(Dbs, newDb)
	c.JSON(http.StatusOK, gin.H{
		"message": "data berhasil terkirim",
		"data":    newDb})
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // untuk sementara kita meng allow kan semuanya ya
	}))

	r.GET("/", handler)
	r.POST("/send", postHandler)
	r.Run(":" + os.Getenv("PORT"))
}
