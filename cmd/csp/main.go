package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/walteh/cloudstack-proxy/pkg/proxy"
)

// Version will be set during build
var Version = "dev"

func main() {
	fmt.Println("CloudStack Proxy - Local Development Server")
	fmt.Printf("Version: %s\n", Version)

	// Parse command line flags
	var (
		bindAddr      = flag.String("bind", "0.0.0.0:8081", "Address to bind server to")
		cloudStackURL = flag.String("cloudstack-url", "http://localhost:8080/client/api", "CloudStack API URL")
		apiKey        = flag.String("api-key", "", "CloudStack API key")
		secretKey     = flag.String("secret-key", "", "CloudStack secret key")
	)
	flag.Parse()

	// Override with environment variables if provided
	if env := os.Getenv("CSP_BIND_ADDRESS"); env != "" {
		*bindAddr = env
	}
	if env := os.Getenv("CSP_CLOUDSTACK_URL"); env != "" {
		*cloudStackURL = env
	}
	if env := os.Getenv("CSP_API_KEY"); env != "" {
		*apiKey = env
	}
	if env := os.Getenv("CSP_SECRET_KEY"); env != "" {
		*secretKey = env
	}

	// Create proxy configuration
	config := &proxy.Config{
		CloudStackURL: *cloudStackURL,
		APIKey:        *apiKey,
		SecretKey:     *secretKey,
	}

	// Create proxy instance
	p, err := proxy.NewProxy(config, Version)
	if err != nil {
		log.Fatalf("Failed to create proxy: %v", err)
	}

	// Create HTTP mux and register handlers
	mux := http.NewServeMux()
	p.RegisterHandlers(mux)

	// Additional development-specific routes can be added here
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Healthy"))
	})

	// Create server
	server := &http.Server{
		Addr:    *bindAddr,
		Handler: mux,
	}

	// Start server in a goroutine
	go func() {
		fmt.Printf("Server starting on %s\n", *bindAddr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Handle graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	fmt.Println("Server stopped")
}
