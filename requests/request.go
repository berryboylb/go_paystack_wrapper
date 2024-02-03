package requests

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const baseUrl = "https://api.paystack.co"

type Request struct {
	APIKey string
}

// NewAPIClient creates a new instance of APIClient.
func NewAPIClient(apiKey string) *Request {
	return &Request{
		APIKey: apiKey,
	}
}

// Post sends a POST request to the specified endpoint with the given payload.
func (c *Request) Post(endpoint string, payload interface{}) (*http.Response, error) {
	url := baseUrl + endpoint

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := http.DefaultClient
	return client.Do(req)
}

//Get sends a get request the specified endpoint
func (c *Request) Get(endpoint string) (*http.Response, error) {
	url := baseUrl + endpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	client := http.DefaultClient
	return client.Do(req)
}
