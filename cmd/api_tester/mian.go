package main

import (
	"api_tester/internal/config"
	"api_tester/internal/runner"
	"fmt"
)

func main() {
	// Load configuration (flags or defaults)
	cfg := config.Load()
	fmt.Println("config: ", cfg)
	// Run the traffic generator
	results := runner.Run(cfg)

	// Print final summary
	results.Print()
}
