package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golangFirstAPI/student"
	"net/http"
)

func StudentsPostHandlerOld(context *gin.Context) {
	var studentInput student.StudentInput

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

	context.JSON(http.StatusOK, gin.H{
		"id":    studentInput.ID,
		"name":  studentInput.Name,
		"score": studentInput.Score,
	})
}
