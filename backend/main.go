package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/khalil9022/HelloWorldGoNextJS/api"
	//DATABASE
)

// type Db struct {
// 	Text string `json:"text,omitempty"`
// 	// Deskripsi string `json:"deskripsi,omitempty"`
// }

// var db []string
// type Todos struct {
// 	gorm.Model
// 	Task string `json:"task"`
// 	Done bool   `json:"done"`
// }

// type DataRequest struct {
// 	Task string `json:"task" binding:"required"`
// }

// type repositoryDb struct {
// 	db *gorm.DB
// }

// func (r *repositoryDb) handler(c *gin.Context) {
// 	var todos []Todos
// 	res := r.db.Find(&todos)
// 	if res.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"data": todos,
// 	})
// }

// func (r *repositoryDb) postHandler(c *gin.Context) {
// 	var data DataRequest

// 	if err := c.ShouldBindJSON(&data); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// todo := Todos{
// 		Task: data.Task,
// 		Done: false,
// 	}

// 	res := r.db.Create(&todo)
// 	if res.Error != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Error.Error()})
// 		return
// 	}

// 	// database.GetDB().Append(data.Text)
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "data berhasil terkirim",
// 		"data":    todo})
// }

func main() {
	db, err := api.SetupDb()
	if err != nil {
		panic(err)
	}

	server := api.MakeServer(db)
	server.RunServer()
	// var db *gorm.DB
	// var err error

	// dbUrl := os.Getenv("DATABASE_URL")

	// if os.Getenv("ENVIRONMENT") == "PROD" {
	// 	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	// } else {
	// 	db, err = gorm.Open(sqlite.Open(dbUrl), &gorm.Config{})
	// }

	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// sqlDB, err := db.DB()
	// if err != nil {
	// 	panic("failed to get database")
	// }

	// if err := sqlDB.Ping(); err != nil {
	// 	panic("failed to ping database")
	// }

	// if err := db.AutoMigrate(&Todos{}); err != nil {
	// 	panic("failed to migrate database")
	// }

	// repo := repositoryDb{
	// 	db: db,
	// }

	// r := gin.Default()
	// // database.StartDB()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"}, // untuk sementara kita meng allow kan semuanya ya
	// }))

	// r.GET("/", repo.handler)
	// r.POST("/send", repo.postHandler)
	// r.Run(":" + os.Getenv("PORT"))
}
