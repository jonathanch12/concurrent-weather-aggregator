# Concurrent Weather Aggregator

A small Go program that processes a large weather dataset concurrently. It reads
`weather_stations.csv` (1,000,000 lines of `Station;Temperature`), splits the
work into batches, and processes those batches in parallel using a pool of
worker goroutines. Each worker keeps its own local statistics, so there is no
shared mutable state on the hot path and no locking is required. The per-worker
results are merged at the end, and the program prints the overall minimum,
average, maximum, and total sum of all temperatures, along with how long the
run took.

## Project Structure

```
concurrent-weather-aggregator/
├── main.go              # Entry point: opens the CSV, runs aggregation, prints the report
├── workers.go           # Concurrency: worker pool, batching, reader goroutine, merge
├── stats.go             # The stats struct (min, max, sum, count)
├── calc.go              # Calculations on stats: add, merge, total, mean, minimum, maximum
├── helper.go            # Helpers: parseLine and printReport
├── go.mod               # Go module definition
├── weather_stations.csv # Input data (Station;Temperature per line)
```

## Requirements

- Go 1.27 or newer (see `go.mod`)

## How to Run

From the project directory:

```powershell
go run .
```

`go run .` compiles every `.go` file in the package together. Use it instead of
`go run main.go`, which compiles only a single file and fails because the code
is split across multiple files.

### Alternative: build then run

If `go run .` is blocked by a Windows Application Control policy (Smart App
Control, WDAC, or AppLocker), build a binary into the project folder and run it
from there:

```powershell
go build -o aggregator.exe .
.\aggregator.exe
```

## Example Output

```
Overall stats (41343 stations, 16 workers):
----------------------------------------
Min:     -99.9000
Average: 49.9000
Max:     99.9000
Sum:     49912345.6000
----------------------------------------
Time Taken: 221 ms
```

The number of workers matches the machine's CPU count (`runtime.NumCPU()`), so
it varies by machine.
