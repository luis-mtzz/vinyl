package main

import (
	"github.com/gin-gonic/gin"
	"github.com/luis-mtzz/app/routes"
)

func main() {
	r := gin.Default()

	// Grabs all releases in a users collection
	r.GET("/:user/releases", routes.GetCollection)

	r.Run()
}
