package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

		// For MVP, we'll simulate GPT-4 response
		response := OnboardingResponse{
			Checklist: []string{
				"Install VS Code",
				"Set up Git",
				"Install Docker",
				"Configure development environment",
				"Clone repository",
				"Install dependencies",
			},
			CodingChallenge: "Write a function that reverses a string and handles Unicode characters correctly",
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
