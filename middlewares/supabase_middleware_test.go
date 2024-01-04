package middlewares

import (
	"inlove-app-server/constants/environments"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestJWTAuthMiddleware(t *testing.T) {
	const secret = "secret"
	// Create a new Gin router
	router := gin.Default()

	// Add the JWTAuthMiddleware to the router
	router.Use(JWTAuthMiddleware())

	// Define a simple GET endpoint for testing
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	environments.JwtSecret = secret
	// generate a valid JWT token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": "John Doe",
		"sub":  "123",
	}).SignedString([]byte(secret))

	if err != nil {
		t.Fatal(err)
	}

	invalidToken, err := jwt.NewWithClaims(jwt.SigningMethodHS384, jwt.MapClaims{
		"name": "John Doe",
		"sub":  "123",
	}).SignedString([]byte("invalid-secret"))

	if err != nil {
		t.Fatal(err)
	}

	// Test with a valid JWT token
	t.Run("valid token", func(t *testing.T) {
		// Create a new HTTP request
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		// Create a new HTTP response recorder
		resp := httptest.NewRecorder()

		// Serve the HTTP request
		router.ServeHTTP(resp, req)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusOK, resp.Code)
	})

	// Test with an invalid JWT token
	t.Run("invalid token", func(t *testing.T) {
		// Create a new HTTP request
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer invalid-token")

		// Create a new HTTP response recorder
		resp := httptest.NewRecorder()

		// Serve the HTTP request
		router.ServeHTTP(resp, req)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	// Test with no Authorization header
	t.Run("no authorization header", func(t *testing.T) {
		// Create a new HTTP request
		req, _ := http.NewRequest("GET", "/test", nil)

		// Create a new HTTP response recorder
		resp := httptest.NewRecorder()

		// Serve the HTTP request
		router.ServeHTTP(resp, req)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	// Test with invalid signing algorithm
	t.Run("invalid signing algorithm", func(t *testing.T) {
		// Create a new HTTP request
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+invalidToken)

		// Create a new HTTP response recorder
		resp := httptest.NewRecorder()

		// Serve the HTTP request
		router.ServeHTTP(resp, req)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})

	// Test with expired token
	t.Run("expired token", func(t *testing.T) {
		// Create a new expired JWT token
		expiredToken, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name": "John Doe",
			"sub":  "123",
			"exp":  time.Now().Add(-time.Hour).Unix(),
		}).SignedString([]byte(secret))

		// Create a new HTTP request
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+expiredToken)

		// Create a new HTTP response recorder
		resp := httptest.NewRecorder()

		// Serve the HTTP request
		router.ServeHTTP(resp, req)

		// Assert the HTTP status code
		assert.Equal(t, http.StatusUnauthorized, resp.Code)
	})
}
