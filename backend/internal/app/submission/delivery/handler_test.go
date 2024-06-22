package delivery_test

import (
	"backend/internal/app/submission/delivery"
	"backend/internal/app/submission/usecase"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSubmissionRegisterHandler_ValidRequest(t *testing.T) {
	// Create a new instance of the mock service
	mockSubmissionService := new(usecase.MockSubmissionService)

	// Mock the Register method to return nil (indicating success)
	mockSubmissionService.On("Register", mock.Anything, mock.Anything).Return(nil)

	// Setup Gin router with the handler
	router := gin.Default()
	router.POST("/api/submission/register", func(c *gin.Context) {
		delivery.SubmissionRegisterHandler(c, mockSubmissionService)
	})

	// Prepare JSON payload for request
	jsonPayload := `{"name": "John Doe", "identity_number": "123456789", "email": "john.doe@example.com", "date_of_birth": "2000-01-01"}`

	// Create a mock HTTP request
	req, err := http.NewRequest("POST", "/api/submission/register", strings.NewReader(jsonPayload))
	assert.NoError(t, err)

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(rr, req)

	// Assert HTTP status code 200 (OK)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert response body contains "Submission created successfully"
	assert.Contains(t, rr.Body.String(), "Submission created successfully")

	// Assert that the Register method on mockSubmissionService was called once
	mockSubmissionService.AssertCalled(t, "Register", mock.Anything, mock.Anything)
}

func TestSubmissionRegisterHandler_ServiceError(t *testing.T) {
	// Create a new instance of the mock service
	mockSubmissionService := new(usecase.MockSubmissionService)

	// Setup Gin router with the handler
	router := gin.Default()
	router.POST("/api/submission/register", func(c *gin.Context) {
		mockSubmissionService.On("Register", mock.Anything, mock.Anything).Return(errors.New("mock error"))
		delivery.SubmissionRegisterHandler(c, mockSubmissionService)
	})

	// Prepare JSON payload for request
	jsonPayload := `{"name": "John Doe", "identity_number": "123456789", "email": "john.doe@example.com", "date_of_birth": "2000-01-01"}`

	// Create a mock HTTP request
	req, err := http.NewRequest("POST", "/api/submission/register", strings.NewReader(jsonPayload))
	assert.NoError(t, err)

	// Set Content-Type header
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(rr, req)

	// Assert HTTP status code 500 (Internal Server Error)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Assert response body contains "mock error"
	assert.Contains(t, rr.Body.String(), "mock error")
}
