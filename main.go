package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/", rootHandler)
	router.GET("/jessica", secondHandler)

	router.GET("/students/:id/:name", studentsHandler)
	router.GET("/query", queryHandler)

	router.POST("/students/post", studentsPostHandler)

	router.Run("localhost:8080")
}

func rootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name": "Ivan Budianto",
		"desc": "Machine Learning Engineer",
	})
}

func secondHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name":  "Jessica Maya",
		"hobby": "Making new things!",
	})
}

func studentsHandler(context *gin.Context) {
	id := context.Param("id")
	name := context.Param("name")

	context.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func queryHandler(context *gin.Context) {
	id := context.Query("id")
	name := context.Query("name")

	context.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

type StudentInput struct {
	ID    string      `json:"id" binding:"required"`
	Name  string      `json:"name" binding:"required"`
	Email json.Number `json:"email" binding:"required,number"`
}

func studentsPostHandler(context *gin.Context) {
	var studentInput StudentInput

	errors := context.ShouldBindJSON(&studentInput)
	if errors != nil {
		for _, err := range errors.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field: %s, condition: %s", err.Field(), err.ActualTag())
			context.JSON(http.StatusBadRequest, errorMessage)
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"id":    studentInput.ID,
		"name":  studentInput.Name,
		"email": studentInput.Email,
	})
}
