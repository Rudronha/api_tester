package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// SendRequest handles GET, POST, PUT, DELETE requests
func SendRequest(method, url, payload string) {
	var req *http.Request
	var err error

	// If payload is empty, don't add a body
	if strings.TrimSpace(payload) == "" {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
	}

	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}

func main() {
	// Example: GET request
	SendRequest("GET", "http://localhost:8080/0UwxV38Hg", "")

	// Example: POST request
	jsonPayload := `{"url": "https://www.linkedin.com/in/rudronha/"}`
	SendRequest("POST", "http://localhost:8080/shorten", jsonPayload)
}
