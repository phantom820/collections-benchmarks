package main

import (
	"collections-benchmarks/benchmarks"
	"collections-benchmarks/lists"
	"collections-benchmarks/queues"
	"collections-benchmarks/sets"
	"os"
)

const (
	trials = 5
	count  = 2
)

func main() {

	timer := benchmarks.Timer{}
	setup := benchmarks.NewSetup(trials, count)
	benchmark := benchmarks.NewBenchmark(setup, &timer)

	setsResults, _ := os.Create("sets.csv")
	listsResults, _ := os.Create("lists.csv")
	queueResults, _ := os.Create("queues.csv")

	sets.Benchmark(setsResults, benchmark)
	lists.Benchmark(listsResults, benchmark)
	queues.Benchmark(queueResults, benchmark)

}
