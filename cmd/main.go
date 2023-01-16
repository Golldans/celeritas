package main

import (
	"celeritas/net/user"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	user.Routes(router)

	err := router.Run(":8083")

	if err != nil {
		log.Fatalf("impossible to start server: %s", err)
	}
}
