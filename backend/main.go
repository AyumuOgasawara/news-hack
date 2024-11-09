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

// POST Request Structure
type ArticleRequest struct {
	Keywords    string `json:"keyword" binding:"required"`
	NumArticles int    `json:"numArticles" binding:"required"`
	PhoneNumber string `json:"phoneNumber" binding:"required"`
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

	// Set / endpoint
	router.GET("/", func(c *gin.Context) {
		// Return JSON Structure
		c.JSON(http.StatusOK, gin.H{
			"Title":       "Welcome to NEWS HACK",
			"Discription": "This is the API for NEWS HACK",
		})
	})

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

	// POST /articles endpoint
	router.POST("/articles", func(c *gin.Context) {
		// Bind to ArticleRequest Structure
		var req ArticleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}

		// make News API url
		newsURL := fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&pageSize=%d&apiKey=%s",
			req.Keywords, req.NumArticles, newsApiKey)

		resp, err := http.Get(newsURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch news"})
			return
		}
		defer resp.Body.Close()

		// Decode Response
		var newsResponse NewsApiResponse
		if err := json.NewDecoder(resp.Body).Decode(&newsResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode news data"})
			return
		}

		for i := 0; i < req.NumArticles; i++ {
			fmt.Printf("This is article title %s", newsResponse.Articles[i].Title)
		}
	})

	// run server
	router.Run(":8080")
}
