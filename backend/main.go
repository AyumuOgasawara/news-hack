package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Article Structure
type Article struct {
	Source struct {
		Name string `json:"name"`
	} `json:"source"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
	PublishedAt string `json:"publishedAt"`
}

// Response Sturucture
type NewsApiResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// get NEWS_API_KEY
	newsApiKey := os.Getenv("NEWS_API_KEY")
	if newsApiKey == "" {
		log.Fatal("NEWS_API_KEY is not set in the environment")
	}

	// Make default Gin Router
	router := gin.Default()

	// Set /health endpoint
	router.GET("/health", func(c *gin.Context) {
		// Return JSON Structure
		c.JSON(http.StatusOK, gin.H{
			"Message": "hello, okutama",
			"status":  200,
		})
	})

	// Set /articles endpoint
	router.GET("/articles", func(c *gin.Context) {
		newsURL := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&apiKey=%s", newsApiKey)

		// Send request to News API
		resp, err := http.Get(newsURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
			return
		}
		defer resp.Body.Close()

		// Decode the response
		var newsResponse NewsApiResponse
		if err := json.NewDecoder(resp.Body).Decode(&newsResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode news data"})
			return
		}

		// Return news response
		c.JSON(http.StatusOK, newsResponse)
	})

	// run server
	router.Run(":8080")
}