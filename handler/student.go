package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golangFirstAPI/student"
	"net/http"
	"strconv"
)

type studentHandler struct {
	studentService student.Service
}

func NewStudentHandler(studentService student.Service) *studentHandler {
	return &studentHandler{studentService}
}

func convertToStudentResponse(studentStruct student.Student) student.StudentResponse {
	studentResponse := student.StudentResponse{
		ID:    studentStruct.ID,
		SID:   studentStruct.SID,
		Name:  studentStruct.Name,
		Score: studentStruct.Score,
	}

	return studentResponse
}

func (handler *studentHandler) StudentIndexHandler(context *gin.Context) {
	students, errorIndex := handler.studentService.FindAll()

	if errorIndex != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": errorIndex,
		})

		return
	}

	var studentsResponse []student.StudentResponse

	for _, row := range students {
		rowStudent := convertToStudentResponse(row)
		studentsResponse = append(studentsResponse, rowStudent)
	}

	context.JSON(http.StatusOK, gin.H{
		"data": studentsResponse,
	})
}

func (handler *studentHandler) StudentShowHandler(context *gin.Context) {
	idString := context.Param("id")
	id, _ := strconv.Atoi(idString)

	studentStruct, errorShow := handler.studentService.FindByID(id)

	if errorShow != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": errorShow,
		})

		return
	}

	studentResponse := convertToStudentResponse(studentStruct)

	context.JSON(http.StatusOK, gin.H{
		"data": studentResponse,
	})
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

func (handler *studentHandler) StudentUpdateHandler(context *gin.Context) {
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

	idString := context.Param("id")
	id, _ := strconv.Atoi(idString)

	student, errorUpdate := handler.studentService.Update(id, studentRequest)

	if errorUpdate != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": errorUpdate,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

func (handler *studentHandler) StudentDeleteHandler(context *gin.Context) {

	idString := context.Param("id")
	id, _ := strconv.Atoi(idString)

	student, errorDelete := handler.studentService.Delete(id)

	if errorDelete != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"errors": errorDelete,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}
