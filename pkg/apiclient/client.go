package apiclient

import (
    "bytes"
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
)

type Client struct {
	BaseURL    string
	APIKey     string
	HTTPClient *http.Client
}

func NewClient(baseURL, apiKey string) *Client {
	return &Client{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) doRequest(method, endpoint string, body interface{}, result interface{}) error {
    url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)
    fmt.Printf("Making %s request to %s\n", method, url)

    var reqBody []byte
    var err error
    if body != nil {
        reqBody, err = json.Marshal(body)
        if err != nil {
            return fmt.Errorf("failed to marshal request body: %w", err)
        }
    }

    req, err := http.NewRequest(method, url, bytes.NewBuffer(reqBody))
    if err != nil {
        return fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-API-Key", c.APIKey)

    resp, err := c.HTTPClient.Do(req)
    if err != nil {
        return fmt.Errorf("request failed: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode >= 400 {
        return fmt.Errorf("API error: %s", resp.Status)
    }

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return fmt.Errorf("failed to read response body: %w", err)
    }

    if result != nil {
        err = json.Unmarshal(respBody, result)
        if err != nil {
            fmt.Printf("Response body: %s\n", string(respBody))
            return fmt.Errorf("failed to unmarshal response: %w", err)
        }
    }

    return nil
}