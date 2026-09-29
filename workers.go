package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

const batchSize = 8192

func runAggregation(f *os.File, path string, workers int) map[string]*stats {
	batches := make(chan []string, workers)
	results := make(chan map[string]*stats, workers)

	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			local := make(map[string]*stats)
			for batch := range batches {
				for _, line := range batch {
					name, temp, ok := parseLine(line)
					if !ok {
						continue
					}
					s := local[name]
					if s == nil {
						s = &stats{}
						local[name] = s
					}
					s.add(temp)
				}
			}
			results <- local
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		scanner := bufio.NewScanner(f)
		scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
		batch := make([]string, 0, batchSize)
		for scanner.Scan() {
			batch = append(batch, scanner.Text())
			if len(batch) == batchSize {
				batches <- batch
				batch = make([]string, 0, batchSize)
			}
		}
		if len(batch) > 0 {
			batches <- batch
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "read %s: %v\n", path, err)
		}
		close(batches)
	}()

	global := make(map[string]*stats)
	for local := range results {
		for name, s := range local {
			g := global[name]
			if g == nil {
				g = &stats{}
				global[name] = g
			}
			g.merge(*s)
		}
	}

	return global
}
