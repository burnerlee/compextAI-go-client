package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
)

type APIClient struct {
	Client  *http.Client
	BaseURL string
	Headers http.Header
	Timeout time.Duration
	Retries int
}

func NewAPIClient(baseURL string, apiKey string) *APIClient {
	return &APIClient{
		Client:  &http.Client{},
		BaseURL: baseURL,
		Headers: http.Header{
			"Authorization": []string{fmt.Sprintf("Bearer %s", apiKey)},
			"Content-Type":  []string{"application/json"},
		},
		Timeout: 10 * time.Second,
		Retries: 3,
	}
}

type JsonResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func (c *APIClient) DoRequest(path string, method string, body interface{}) (*JsonResponse, error) {
	var jsonData io.Reader

	if body != nil {
		jsonDataBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal body: %w", err)
		}
		jsonData = bytes.NewBuffer(jsonDataBytes)
	}

	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseURL, path), jsonData)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	request.Header = c.Headers

	for i := 0; i < c.Retries; i++ {
		response, err := c.Client.Do(request)

		// check if request timed out
		if err != nil {
			if err, ok := err.(net.Error); ok && err.Timeout() {
				continue
			}
			return nil, fmt.Errorf("request failed: %w", err)
		}
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}
		// fmt.Println(string(responseBody))
		var jsonResponse interface{}
		err = json.Unmarshal(responseBody, &jsonResponse)
		if err != nil {
			return nil, fmt.Errorf("failed to decode response: %w", err)
		}

		return &JsonResponse{
			Status: response.StatusCode,
			Data:   jsonResponse,
		}, nil
	}

	return nil, fmt.Errorf("failed to make request: %w", err)
}
