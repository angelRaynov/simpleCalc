package server

import (
	"calc/infrastructure/config"
	calcHttp "calc/internal/calculator/delivery/http"
	"calc/internal/calculator/repository"
	"calc/internal/calculator/usecase"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	cfg := config.New()

	cr := repository.NewCalculatorRepository()
	cu := usecase.NewCalculatorUseCase(cr)
	ch := calcHttp.NewCalculatorHandler(cu)

	router := gin.Default()
	router.GET("/errors", ch.ErrorHandler)
	router.POST("/evaluate", ch.EvaluateHandler)
	router.POST("/validate", ch.ValidateHandler)

	port := fmt.Sprintf(":%s", cfg.APIPort)

	server := &http.Server{
		Addr:    port,
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
		log.Printf("listening on port :%s", cfg.APIPort)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// Create a context with a timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server stopped gracefully")
}
