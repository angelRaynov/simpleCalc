package http

import (
	"bytes"
	"calc/internal/calculator"
	"calc/internal/calculator/repository"
	"calc/internal/calculator/usecase"
	"calc/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculatorHandler_EvaluateHandler(t *testing.T) {
	cr := repository.NewCalculatorRepository()
	uc := usecase.NewCalculatorUseCase(cr)
	h := NewCalculatorHandler(uc)
	router := gin.Default()
	router.POST("/evaluate", h.EvaluateHandler)

	tests := []struct {
		name     string
		payload  []byte
		want     string
		wantCode int
	}{
		{
			name:     "proper request",
			payload:  []byte(`{"expression":"What is 5 plus 3"}`),
			want:     "{\n    \"result\": 8\n}",
			wantCode: http.StatusOK,
		},
		{
			name:     "invalid request",
			payload:  []byte(`{"invali is 5 plus 3"}`),
			want:     "{\n    \"error\": \"invalid request\"\n}",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "eval error",
			payload:  []byte(`{"expression":"Who is she"}`),
			want:     "{\n    \"error\": \"non math question\"\n}",
			wantCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodPost, "/evaluate", bytes.NewBuffer(tt.payload))

			// Create a response recorder
			recorder := httptest.NewRecorder()
			// Execute the request
			router.ServeHTTP(recorder, req)

			// Verify the response
			assert.Equal(t, tt.wantCode, recorder.Code)
			assert.Equal(t, tt.want, recorder.Body.String())
		})
	}

}

func TestCalculatorHandler_ValidateHandler(t *testing.T) {
	cr := repository.NewCalculatorRepository()
	uc := usecase.NewCalculatorUseCase(cr)
	h := NewCalculatorHandler(uc)
	router := gin.Default()
	router.POST("/validate", h.ValidateHandler)

	tests := []struct {
		name     string
		payload  []byte
		want     string
		wantCode int
	}{
		{
			name:     "proper request",
			payload:  []byte(`{"expression":"What is 5 plus 3"}`),
			want:     "{\n    \"valid\": true\n}",
			wantCode: http.StatusOK,
		},
		{
			name:     "invalid request",
			payload:  []byte(`{"invali is 5 plus 3"}`),
			want:     "{\n    \"error\": \"invalid request\"\n}",
			wantCode: http.StatusBadRequest,
		},
		{
			name:     "eval error",
			payload:  []byte(`{"expression":"Who is she"}`),
			want:     "{\n    \"valid\": false\n}",
			wantCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodPost, "/validate", bytes.NewBuffer(tt.payload))

			// Create a response recorder
			recorder := httptest.NewRecorder()
			// Execute the request
			router.ServeHTTP(recorder, req)

			// Verify the response
			assert.Equal(t, tt.wantCode, recorder.Code)
			assert.Equal(t, tt.want, recorder.Body.String())
		})
	}

}

func TestCalculatorHandler_ErrorHandler(t *testing.T) {
	emptyStore := repository.NewCalculatorRepository()
	cr := repository.NewCalculatorRepository()

	cr.Store(&model.Error{})

	tests := []struct {
		name     string
		want     string
		wantCode int
		uc       calculator.UseCaser
	}{
		{
			name:     "no records",
			want:     "{\n    \"error\": \"no records\"\n}",
			wantCode: http.StatusNotFound,
			uc:       usecase.NewCalculatorUseCase(emptyStore),
		},
		{
			name:     "proper request",
			want:     "[\n    {\n        \"expression\": \"\",\n        \"endpoint\": \"\",\n        \"frequency\": 0,\n        \"type\": \"\"\n    }\n]",
			wantCode: http.StatusOK,
			uc:       usecase.NewCalculatorUseCase(cr),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewCalculatorHandler(tt.uc)
			router := gin.Default()
			router.GET("/errors", h.ErrorHandler)

			// Create a new HTTP request
			req := httptest.NewRequest(http.MethodGet, "/errors", nil)

			// Create a response recorder
			recorder := httptest.NewRecorder()
			// Execute the request
			router.ServeHTTP(recorder, req)

			// Verify the response
			assert.Equal(t, tt.wantCode, recorder.Code)
			assert.Equal(t, tt.want, recorder.Body.String())
		})
	}

}
