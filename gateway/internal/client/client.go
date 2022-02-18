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
	// AuthLogoutService ...
	AuthLogoutService = New(fmt.Sprintf("http://%s:%s", os.Getenv("AUTH_SERVER_DOCKER_HOST"), os.Getenv("AUTH_SERVER_PORT")), "/logout")
	// AuthRegistrationService ...
	AuthRegistrationService = New(fmt.Sprintf("http://%s:%s", os.Getenv("AUTH_SERVER_DOCKER_HOST"), os.Getenv("AUTH_SERVER_PORT")), "/registration")
	// AuthRefreshService ...
	AuthRefreshService = New(fmt.Sprintf("http://%s:%s", os.Getenv("AUTH_SERVER_DOCKER_HOST"), os.Getenv("AUTH_SERVER_PORT")), "/refresh")
)

// CtxKey ...
type CtxKey int8

const (
	// AccessTokenCtxKey ...
	AccessTokenCtxKey CtxKey = 1
	// RefreshTokenCtxKey ...
	RefreshTokenCtxKey CtxKey = 2
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
