package studentHandlerBase

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golangFirstAPI/student"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func StudentsDestroyHandler(context *gin.Context) {
	// DB Connection
	dsn := "host=localhost user=xcessive password=1010 dbname=student port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Taking ID from param
	id := context.Param("id")

	// Taking record
	var studentRecord student.Student
	errorFind := db.Where("id = ?", id).First(&studentRecord).Error
	if errorFind != nil {
		fmt.Println("Error Searching Data")
		return
	}

	errorUpdate := db.Delete(&studentRecord).Error
	if errorUpdate != nil {
		fmt.Println("Error Deleting Data")
		return
	}

	context.JSON(http.StatusOK, "Deleted the student!")
}
