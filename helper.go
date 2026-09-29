package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func parseLine(line string) (name string, temp float64, ok bool) {
	i := strings.IndexByte(line, ';')
	if i < 0 {
		return "", 0, false
	}
	name = line[:i]
	t, err := strconv.ParseFloat(strings.TrimSpace(line[i+1:]), 64)
	if err != nil {
		return "", 0, false
	}
	return name, t, true
}

func printReport(global map[string]*stats, workers int, elapsed time.Duration) {
	var overall stats
	for _, s := range global {
		overall.merge(*s)
	}

	fmt.Printf("Overall stats (%d stations, %d workers):\n", len(global), workers)
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("Min:     %.4f\n", overall.minimum())
	fmt.Printf("Average: %.4f\n", overall.mean())
	fmt.Printf("Max:     %.4f\n", overall.maximum())
	fmt.Printf("Sum:     %.4f\n", overall.total())
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("Time Taken: %d ms\n", elapsed.Milliseconds())
}
