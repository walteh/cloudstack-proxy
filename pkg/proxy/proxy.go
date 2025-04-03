package proxy

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"connectrpc.com/connect"
	corev1 "github.com/walteh/cloudstack-proxy/gen/proto/cloudstack/core/v1"
	"github.com/walteh/cloudstack-proxy/gen/proto/cloudstack/core/v1/corev1connect"
)

// Proxy is the main CloudStack proxy that implements all services
type Proxy struct {
	// Configuration options
	config *Config

	// CloudStack API client
	csClient CloudStackClient

	// Server startup time
	startTime time.Time

	// Version of the proxy
	version string
}

// Config contains proxy configuration
type Config struct {
	// CloudStack API URL
	CloudStackURL string

	// API key for CloudStack
	APIKey string

	// Secret key for CloudStack
	SecretKey string
}

// CloudStackClient defines the interface for interacting with CloudStack
type CloudStackClient interface {
	// ExecuteRequest sends a request to the CloudStack API
	ExecuteRequest(ctx context.Context, apiName string, params map[string]string) ([]byte, error)
}

// NewProxy creates a new CloudStack proxy with the given configuration
func NewProxy(config *Config, version string) (*Proxy, error) {
	// Create CloudStack client (implementation to be added)
	csClient := &defaultCloudStackClient{
		baseURL:   config.CloudStackURL,
		apiKey:    config.APIKey,
		secretKey: config.SecretKey,
	}

	return &Proxy{
		config:    config,
		csClient:  csClient,
		startTime: time.Now(),
		version:   version,
	}, nil
}

// RegisterHandlers registers all service handlers with the given mux
func (p *Proxy) RegisterHandlers(mux *http.ServeMux) {
	// Register health service
	healthPath, healthHandler := corev1connect.NewHealthServiceHandler(p)
	mux.Handle(healthPath, healthHandler)

	// Register additional services here as they are implemented
}

// Check implements the health check service
func (p *Proxy) Check(
	ctx context.Context,
	req *connect.Request[corev1.HealthCheckRequest],
) (*connect.Response[corev1.HealthCheckResponse], error) {
	// Check CloudStack connectivity
	csStatus := corev1.HealthCheckResponse_SERVING
	message := "Healthy"

	// Try to ping CloudStack API
	_, err := p.csClient.ExecuteRequest(ctx, "ping", nil)
	if err != nil {
		csStatus = corev1.HealthCheckResponse_NOT_SERVING
		message = fmt.Sprintf("CloudStack API unavailable: %v", err)
	}

	res := connect.NewResponse(&corev1.HealthCheckResponse{
		Status:  csStatus,
		Version: p.version,
		Message: message,
	})

	return res, nil
}

// defaultCloudStackClient implements the CloudStackClient interface
type defaultCloudStackClient struct {
	baseURL   string
	apiKey    string
	secretKey string
}

// ExecuteRequest implements CloudStackClient.ExecuteRequest
func (c *defaultCloudStackClient) ExecuteRequest(
	ctx context.Context,
	apiName string,
	params map[string]string,
) ([]byte, error) {
	// This is a stub implementation - actual implementation will use CloudStack API

	// TODO:
	// 1. Create request with required parameters
	// 2. Add authentication (API key + signature)
	// 3. Send HTTP request to CloudStack API
	// 4. Process response

	return []byte(`{}`), nil
}
