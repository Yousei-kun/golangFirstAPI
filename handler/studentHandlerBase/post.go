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

func StudentsPostHandler(context *gin.Context) {

	dsn := "host=localhost user=xcessive password=1010 dbname=student port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var studentInput student.Student
	err := context.ShouldBindJSON(&studentInput)

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

	errorPost := db.Create(&studentInput).Error

	if errorPost != nil {
		fmt.Println("Error Creating Data")
		return
	}

	context.JSON(http.StatusOK, "Recorded New Student!")
}
