package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	const path = "weather_stations.csv"

	workers := runtime.NumCPU()
	start := time.Now()

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open %s: %v\n", path, err)
		os.Exit(1)
	}
	defer f.Close()

	global := runAggregation(f, path, workers)

	elapsed := time.Since(start)
	printReport(global, workers, elapsed)
}
