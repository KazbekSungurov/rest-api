package client

import (
	"fmt"
	"gateway/pkg/logging"
	"gateway/pkg/rest"
	"net/http"
	"os"
	"time"
)

var (
	// AuthLoginService ...
	AuthLoginService = New(fmt.Sprintf("http://%s:%s", os.Getenv("AUTH_SERVER_DOCKER_HOST"), os.Getenv("AUTH_SERVER_PORT")), "/login")
)

// Client ...
type Client struct {
	Base     rest.BaseClient
	Resource string
}

// New ...
func New(baseURL string, resource string) *Client {
	return &Client{
		Base: rest.BaseClient{
			BaseURL: baseURL,
			HTTPClient: &http.Client{
				Timeout: 10 * time.Second,
			},
			Logger: logging.GetLogger(),
		},
		Resource: resource,
	}
}
