package auth

import (
	"context"
	"gateway/internal/apperror"
	"gateway/internal/client"
	"gateway/pkg/response"
	"io"
)

// Registration ...
func Registration(ctx context.Context, c *client.Client, body io.Reader) (*response.Service, error) {
	url, err := c.Base.BuildURL(c.Resource, nil)
	if err != nil {
		return nil, err
	}

	req, err := c.Base.CreateRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	resp, err := c.Base.SendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	if resp.IsOk {
		resp, err := c.Base.ReadResponse(resp, nil)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, apperror.APIError(resp.Error.ErrorCode, resp.Error.Message, resp.Error.DeveloperMessage)
}
