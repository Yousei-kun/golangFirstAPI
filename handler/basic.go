package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RootHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name": "Ivan Budianto",
		"desc": "Machine Learning Engineer",
	})
}

func SecondHandler(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"name":  "Jessica Maya",
		"hobby": "Making new things!",
	})
}

func StudentsHandler(context *gin.Context) {
	id := context.Param("id")
	name := context.Param("name")

	context.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}

func QueryHandler(context *gin.Context) {
	id := context.Query("id")
	name := context.Query("name")

	context.JSON(http.StatusOK, gin.H{
		"id":   id,
		"name": name,
	})
}
