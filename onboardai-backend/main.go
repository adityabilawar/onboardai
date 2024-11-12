package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"` // Added role field
}

type OnboardingRequest struct {
	EngineerEmail string `json:"engineerEmail"`
	Documentation string `json:"documentation"`
}

type OnboardingResponse struct {
	Checklist       []string `json:"checklist"`
	CodingChallenge string   `json:"codingChallenge"`
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		// Log error but don't fail - env vars might be set another way
		println("Error loading .env file:", err)
	}

	r := gin.Default()

	// Configure CORS with specific options
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"} // Vue dev server default port
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Type"}
	r.Use(cors.New(config))

	// Simple in-memory storage for MVP
	users := make(map[string]User)
	onboardingData := make(map[string]OnboardingResponse)

	// Auth endpoints
	r.POST("/api/signup", func(c *gin.Context) {
		var user User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if user already exists
		if _, exists := users[user.Email]; exists {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}

		users[user.Email] = user
		c.JSON(http.StatusOK, gin.H{
			"message": "User registered",
			"user": gin.H{
				"email": user.Email,
				"role":  user.Role,
			},
		})
	})

	// Add login endpoint
	r.POST("/api/login", func(c *gin.Context) {
		var credentials User
		if err := c.BindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if user, exists := users[credentials.Email]; exists {
			if user.Password == credentials.Password {
				c.JSON(http.StatusOK, gin.H{
					"message": "Login successful",
					"user": gin.H{
						"email": user.Email,
						"role":  user.Role,
					},
				})
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	})

	// Generate onboarding checklist
	r.POST("/api/generate-onboarding", func(c *gin.Context) {
		var req OnboardingRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call GPT-4 API
		client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: "gpt-4",
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    "user",
						Content: "Using this documentation of a codebase, create a checklist for new hired dev engineer along with a coding challenge.\n\n" + req.Documentation,
					},
				},
			},
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate onboarding"})
			return
		}

		// Parse GPT response - assuming format: checklist items followed by coding challenge
		gptResponse := resp.Choices[0].Message.Content
		parts := strings.Split(gptResponse, "\nCoding Challenge:")

		checklist := strings.Split(strings.TrimSpace(parts[0]), "\n")
		codingChallenge := ""
		if len(parts) > 1 {
			codingChallenge = strings.TrimSpace(parts[1])
		}

		response := OnboardingResponse{
			Checklist:       checklist,
			CodingChallenge: codingChallenge,
		}

		onboardingData[req.EngineerEmail] = response
		c.JSON(http.StatusOK, response)
	})

	r.GET("/api/onboarding/:email", func(c *gin.Context) {
		email := c.Param("email")
		if data, exists := onboardingData[email]; exists {
			c.JSON(http.StatusOK, data)
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "No onboarding data found"})
		}
	})

	r.Run(":8080")
}
