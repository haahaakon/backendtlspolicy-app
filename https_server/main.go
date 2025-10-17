package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	server := gin.Default()

	server.GET("/", func(req *gin.Context) {
		req.String(http.StatusOK, "pong!")
	})

	// Look for certificates in mounted volume or fallback to local files
	certFile := getEnvOrDefault("TLS_CERT_FILE", "/etc/tls/tls.crt")
	keyFile := getEnvOrDefault("TLS_KEY_FILE", "/etc/tls/tls.key")

	// Fallback to local files for development
	if _, err := os.Stat(certFile); os.IsNotExist(err) {
		log.Printf("Certificate not found at %s, falling back to cert.pem", certFile)
		certFile = "cert.pem"
	}
	if _, err := os.Stat(keyFile); os.IsNotExist(err) {
		log.Printf("Private key not found at %s, falling back to key.pem", keyFile)
		keyFile = "key.pem"
	}

	log.Printf("Starting HTTPS server on :8888 with cert: %s, key: %s", certFile, keyFile)
	server.RunTLS(":8888", certFile, keyFile)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
