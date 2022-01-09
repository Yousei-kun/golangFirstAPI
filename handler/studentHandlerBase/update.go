package studentHandlerBase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golangFirstAPI/student"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func StudentsUpdateHandler(context *gin.Context) {
	// DB Connection
	dsn := "host=localhost user=xcessive password=1010 dbname=student port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Taking ID from param
	id := context.Param("id")

	// Binding input JSON
	var studentRequest student.StudentRequest
	err := context.ShouldBindJSON(&studentRequest)
	if err != nil {
		var errorMessages []string

		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		context.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	// Taking record
	var studentRecord student.Student
	errorFind := db.Where("id = ?", id).First(&studentRecord).Error
	if errorFind != nil {
		fmt.Println("Error Searching Data")
		return
	}

	// Updating record
	newScore, _ := studentRequest.Score.(float64)

	studentRecord.SID = studentRequest.SID
	studentRecord.Name = studentRequest.Name
	studentRecord.Score = newScore

	errorUpdate := db.Save(&studentRecord).Error
	if errorUpdate != nil {
		fmt.Println("Error Updating Data")
		return
	}

	context.JSON(http.StatusOK, "Updated the student!")
}
