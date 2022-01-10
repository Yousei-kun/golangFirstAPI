package main

import (
	"github.com/gin-gonic/gin"
	"golangFirstAPI/handler"
	"golangFirstAPI/handler/studentHandlerBase"
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

	studentRepository := student.NewRepository(db)
	studentService := student.NewService(studentRepository)
	studentHandler := handler.NewStudentHandler(studentService)

	router := gin.Default()

	v1 := router.Group("/v1")

	//SIMPLE CRUD
	v1.POST("/students/simple/post", studentHandlerBase.StudentsPostHandler)
	v1.GET("/students/simple/index", studentHandlerBase.StudentsIndexHandler)
	v1.POST("/students/simple/update/:id", studentHandlerBase.StudentsUpdateHandler)
	v1.POST("/students/simple/destroy/:id", studentHandlerBase.StudentsDestroyHandler)

	//Create WITH HANDLER (Updating Soon!)
	v1.GET("/students", studentHandler.StudentIndexHandler)
	v1.GET("/students/show/:id", studentHandler.StudentShowHandler)
	v1.POST("/students/create", studentHandler.StudentPostHandler)
	v1.PUT("/students/update/:id", studentHandler.StudentUpdateHandler)
	v1.DELETE("/students/delete/:id", studentHandler.StudentDeleteHandler)

	router.Run("localhost:8080")
}
