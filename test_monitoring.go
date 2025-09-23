package main

import (
	"fmt"
	"log"
	"time"

	"github.com/achmdndy/safa-life-api/src/infrastructure/monitoring"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🧪 Testing monitoring integration...")

	// Initialize Jaeger tracing
	fmt.Println("🔍 Initializing Jaeger tracing...")
	jaegerCleanup := monitoring.InitJaeger(monitoring.JaegerConfig{
		ServiceName: "safa-life-api-test",
		Endpoint:    "http://localhost:14268/api/traces",
	})
	defer jaegerCleanup()

	// Create Gin router
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// Setup monitoring middleware
	fmt.Println("📊 Setting up monitoring...")
	monitoringMiddleware := monitoring.NewMonitoringMiddleware()
	monitoringMiddleware.Setup(router, "safa-life-api-test")

	// Add test endpoints
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Monitoring test endpoint",
			"status":  "ok",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
			"time":   time.Now(),
		})
	})

	// Start server
	serverAddr := ":8080"
	fmt.Printf("🌙 Test server running on http://localhost%s\n", serverAddr)
	fmt.Printf("📊 Prometheus metrics available at http://localhost%s/metrics\n", serverAddr)
	fmt.Println("🧪 Test endpoints:")
	fmt.Println("   - GET http://localhost:8080/")
	fmt.Println("   - GET http://localhost:8080/health")
	fmt.Println("   - GET http://localhost:8080/metrics")

	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start test server: %v", err)
	}
}