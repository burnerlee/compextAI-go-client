package api

import (
	"fmt"
	"testing"
)

func TestDoRequest(t *testing.T) {
	client := NewAPIClient("https://api.compextai.dev", "test")

	response, err := client.DoRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("failed to do request: %v", err)
	}

	fmt.Println(response)
}
