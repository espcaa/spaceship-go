package spaceship

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	baseURL    string
	apiKey     string
	apiSecret  string
	httpClient *http.Client
}

var defaultBaseURL = "https://spaceship.dev/api/v1"

func NewClient(apiKey, apiSecret string) *Client {
	return &Client{
		baseURL:    defaultBaseURL,
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		httpClient: http.DefaultClient,
	}
}

func NewCustomClient(apiKey, apiSecret string, httpClient *http.Client, baseURL ...string) *Client {
	return &Client{
		baseURL:    baseURL[0],
		apiKey:     apiKey,
		apiSecret:  apiSecret,
		httpClient: httpClient,
	}
}

func (c *Client) newRequest(method, path string, body any) (*http.Request, error) {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, c.baseURL+path, &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Key", c.apiKey)
	req.Header.Set("X-Api-Secret", c.apiSecret)
	return req, nil
}

func (c *Client) do(req *http.Request, out any) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %s", resp.Status)
	}

	if out != nil {
		return json.NewDecoder(resp.Body).Decode(out)
	}

	return nil
}
