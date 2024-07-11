package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luis-mtzz/app/models"
)

func GetCollection(c *gin.Context) {
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

		var result models.CollectionItems
		if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "can not decode JSON",
			})
			return
		}

		// Extract the desired fields from each release
		var simplifiedReleases []models.Release
		for _, release := range result.Releases {
			simplifiedRelease := models.Release{
				ArtistName: release.BasicInformation.Artists[0].Name,
				Genres:     release.BasicInformation.Genres,
				Title:      release.BasicInformation.Title,
			}
			simplifiedReleases = append(simplifiedReleases, simplifiedRelease)
		}

		// Print the simplified releases
		for _, sr := range simplifiedReleases {
			fmt.Printf("Artist: %s, Title: %s, Genres: %v\n", sr.ArtistName, sr.Title, sr.Genres)
		}

		// Convert the simplified releases to JSON
		formattedJSON, err := json.MarshalIndent(simplifiedReleases, "", "  ")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "can't format JSON"})
			return
		}

		c.Data(http.StatusOK, "application/json", formattedJSON)
	}
