package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golangFirstAPI/student"
	"net/http"
)

type studentHandler struct {
	studentService student.Service
}

func NewStudentHandler(studentService student.Service) *studentHandler {
	return &studentHandler{studentService}
}

func (handler *studentHandler) StudentPostHandler(context *gin.Context) {
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

	student, errorPost := handler.studentService.Create(studentRequest)

	if errorPost != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": err,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}
