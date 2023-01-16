package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func Routes(router *gin.Engine) {
	new_router := router.Group("/user")
	new_router.GET("/find", GetUserHandler)
}

func GetUserHandler(context *gin.Context) {
	context.JSON(http.StatusNotFound, gin.H{"error": "pelo menos bateu"})
}
