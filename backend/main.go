package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/khalil9022/HelloWorldGoNextJS/database"
)

func handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": database.GetDB(),
	})
}

// type Db struct {
// 	Text string `json:"text,omitempty"`
// 	// Deskripsi string `json:"deskripsi,omitempty"`
// }

// var db []string

type DataRequest struct {
	Text string `json:"text"`
}

func postHandler(c *gin.Context) {
	var data DataRequest

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.GetDB().Append(data.Text)
	c.JSON(http.StatusOK, gin.H{
		"message": "data berhasil terkirim",
		"data":    data.Text})
}

func main() {
	database.StartDB()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // untuk sementara kita meng allow kan semuanya ya
	}))

	r.GET("/", handler)
	r.POST("/send", postHandler)
	r.Run(":" + os.Getenv("PORT"))
}
