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
		"message": "Hello world 2",
	})
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // untuk sementara kita meng allow kan semuanya ya
	}))
	r.GET("/", handler)
	r.Run(":" + os.Getenv("PORT"))
}
