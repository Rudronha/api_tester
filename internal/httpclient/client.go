package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"	
	"time"
)

// Response represents the result of an HTTP request
type Response struct {
	StatusCode  int           // HTTP Status like "200 OK"
	Body 		string        // Response body as string
	Duration 	time.Duration // Time taken for the request
	Success  	bool          // True if status code is 2xx
	Error    	string        // Error message if any
}

// Client wraps an http.Client for our tester
type Client struct {
	httpClient *http.Client
}

// New creates a new Client with default timeout
func New() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// Do sends a single HTTP request
func (c *Client) Do(method, url, payload string) Response {
	start := time.Now()

	var req *http.Request
	var err error
	if payload == "" {
		req, err = http.NewRequest(method, url, nil)
	} else {
		req, err = http.NewRequest(method, url, bytes.NewBuffer([]byte(payload)))
		req.Header.Set("Content-Type", "application/json")
	}
	if err != nil {
		return Response{Success: false, Error: err.Error()}
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Response{Success: false, Error: err.Error()}
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{Success: false, Error: err.Error()}
	}

	return Response{
		StatusCode:		resp.StatusCode,
		Body:         	string(bodyBytes),
		Duration: 		time.Since(start),
		Success:  		resp.StatusCode >= 200 && resp.StatusCode < 300,
	}
}
