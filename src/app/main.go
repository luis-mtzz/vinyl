package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Grabs all releases in a users collection
	r.GET("/:user/releases", func(c *gin.Context) {
		user := c.Param("user")
		apiURL := fmt.Sprintf("https://api.discogs.com/users/%s/collection/folders/0/releases?sort=artist", user)
		response, err := http.Get(apiURL)
		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		defer response.Body.Close()

		if response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		var v interface{}
		json.NewDecoder(response.Body).Decode(&v)
		fmt.Println("Test", v)

		c.JSON(200, v)
	})
	r.Run()
}
