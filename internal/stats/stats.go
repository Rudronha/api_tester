package stats

import (
	"fmt"
	"sync"
)

// Result stores the outcome of one HTTP request
type Result struct {
	StatusCode int
	DurationMs int64
	Error      string
}

// Stats aggregates results
type Stats struct {
	mu      sync.Mutex
	results []Result
}

func NewStats() *Stats {
	return &Stats{
		results: []Result{},
	}
}

// Add a new result to the stats
func (s *Stats) Add(r Result) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.results = append(s.results, r)
}

// Print final statistics
func (s *Stats) Print() {
	s.mu.Lock()
	defer s.mu.Unlock()

	total := len(s.results)
	if total == 0 {
		fmt.Println("No results collected.")
		return
	}

	var totalDuration int64
	statusCount := make(map[int]int)
	errorCount := 0

	for _, r := range s.results {
		totalDuration += r.DurationMs
		statusCount[r.StatusCode]++
		if r.Error != "" {
			errorCount++
		}
	}

	fmt.Printf("\n--- Test Results ---\n")
	fmt.Printf("Total Requests: %d\n", total)
	fmt.Printf("Errors: %d\n", errorCount)
	fmt.Printf("Average Response Time: %dms\n", totalDuration/int64(total))
	fmt.Printf("Status Codes: %v\n", statusCount)
}
