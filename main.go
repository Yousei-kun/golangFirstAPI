package main

import (
	"github.com/gin-gonic/gin"
	"golangFirstAPI/handler"
	"golangFirstAPI/student"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {

	dsn := "host=localhost user=xcessive password=1010 dbname=student port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DB ERROR!")
	}

	db.AutoMigrate(&student.Student{})

	router := gin.Default()

	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler)
	v1.GET("/jessica", handler.SecondHandler)
	v1.GET("/students/:id/:name", handler.StudentsHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/students/post/old", handler.StudentsPostHandlerOld)
	v1.POST("/students/post", handler.StudentsPostHandler)

	router.Run("localhost:8080")
}

func connectDB() {

}
