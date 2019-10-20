package pivotal

import (
	"errors"
	"net/http"
	"net/url"
)

const defaultBaseURL = "https://www.pivotaltracker.com/services/v5"

type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	Token      string
}

func NewClient(token string) (*Client, error) {
	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		return nil, err
	}
	if len(token) == 0 {
		return nil, errors.New("token is required")
	}

	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
		Token:      token,
	}, nil
}
