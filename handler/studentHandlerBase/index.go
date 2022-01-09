package studentHandlerBase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangFirstAPI/student"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func StudentsIndexHandler(context *gin.Context) {

	dsn := "host=localhost user=xcessive password=1010 dbname=student port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var students []student.Student

	errorIndex := db.Find(&students).Error

	if errorIndex != nil {
		fmt.Println("Error Viewing Data")
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"students": students,
	})
}

func StudentsShowHandler(context *gin.Context) {

	dsn := "host=localhost user=xcessive password=1010 dbname=student port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var firstStudent student.Student

	errorIndex := db.First(&firstStudent).Error

	if errorIndex != nil {
		fmt.Println("Error Viewing Data")
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"record": firstStudent,
	})
}
