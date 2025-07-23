package runner

import (
	"fmt"
	"sync"
	"time"

	"api_tester/internal/config"
	"api_tester/internal/httpclient"
	"api_tester/internal/stats"
)

// Run executes the load test and returns collected stats
func Run(cfg config.Config) *stats.Stats {
	fmt.Printf("Running %d requests with concurrency=%d\n", cfg.NumRequests, cfg.Concurrency)

	results := stats.NewStats()
	var wg sync.WaitGroup
	tasks := make(chan int, cfg.NumRequests)

	// Launch workers
	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			client := httpclient.New()

			for range tasks {
				start := time.Now()
				resp := client.Do(cfg.Method, cfg.URL, cfg.Payload)
				duration := time.Since(start).Milliseconds()

				result := stats.Result{
					StatusCode: resp.StatusCode,
					DurationMs: duration,
				}
				if resp.Error != "" {
					result.Error = resp.Error
				}
				results.Add(result)
			}
		}()
	}

	// Distribute tasks
	for i := 0; i < cfg.NumRequests; i++ {
		tasks <- i
	}
	close(tasks)

	wg.Wait()
	return results
}
