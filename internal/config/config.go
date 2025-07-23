package config

import "flag"

// Config stores all user-defined settings
type Config struct {
	URL         string
	Method      string
	Payload     string
	NumRequests int
	Concurrency int
}

// Load parses CLI flags
func Load() Config {
	url := flag.String("url", "http://localhost:8080", "Target URL")
	method := flag.String("method", "GET", "HTTP method")
	payload := flag.String("data", "", "Request payload for POST/PUT")
	numRequests := flag.Int("n", 10, "Number of requests")
	concurrency := flag.Int("c", 2, "Number of concurrent workers")

	flag.Parse()

	return Config{
		URL:         *url,
		Method:      *method,
		Payload:     *payload,
		NumRequests: *numRequests,
		Concurrency: *concurrency,
	}
}
