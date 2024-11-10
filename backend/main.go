package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

type WhatsappMessage struct {
	MessagingProduct string `json:"messaging_product"`
	To               string `json:"to"`
	Type             string `json:"type"`
	Text             struct {
		Body string `json:"body"`
	} `json:"text"`
}

func sendMessageToWhatsapp(article Article, phoneNumber string) error {
	whatsappToken := os.Getenv("WHATSAPP_TOKEN")
	if whatsappToken == "" {
		return fmt.Errorf("WhatsApp token is missing")
	}
	// メッセージ本文を作成
	messageBody := fmt.Sprintf("Source: %s\nTitle: %s\nDescription: %s\nURL: %s\nPublished At: %s",
		article.Source.Name, article.Title, article.Description, article.Url, article.PublishedAt)

	// WhatsApp API用のメッセージ構造体を作成
	message := WhatsappMessage{
		MessagingProduct: "whatsapp",
		To:               phoneNumber,
		Type:             "text",
	}
	message.Text.Body = messageBody

	// JSON形式にエンコード
	messageData, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	phoneNumberId := os.Getenv("PHONE_NUMBER_ID")
	if phoneNumberId == "" {
		return fmt.Errorf("WhatsApp token is missing")
	}

	requestURI := fmt.Sprintf("https://graph.facebook.com/v20.0/%s/messages", phoneNumberId)

	// WhatsApp APIにPOSTリクエストを送信
	req, err := http.NewRequest("POST", requestURI, bytes.NewBuffer(messageData))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+whatsappToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	// レスポンスをログに出力
	respBody, _ := io.ReadAll(resp.Body)
	log.Printf("Response Status: %s, Body: %s\n", resp.Status, respBody)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-OK HTTP status: %s", resp.Status)
	}

	return nil
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

		for _, article := range newsResponse.Articles {
			err := sendMessageToWhatsapp(article, req.PhoneNumber)
			if err != nil {
				log.Printf("Failed to send message to WhatsApp: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to send Message to WhatsApp"})
			}
		}

		c.JSON(http.StatusOK, gin.H{"message": "Articles sent to WhatsApp"})
	})

	// run server
	router.Run(":8080")
}
