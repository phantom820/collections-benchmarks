// Package sets provides benchmark tests for sets.
package sets

import (
	"collections-benchmarks/benchmarks"
	"collections-benchmarks/data"
	"fmt"
	"math/rand"
	"os"

	"github.com/phantom820/collections/sets"
	"github.com/phantom820/collections/sets/hashset"
	"github.com/phantom820/collections/sets/linkedhashset"
	"github.com/phantom820/collections/sets/treeset"

	"github.com/phantom820/collections/types"
)

func BenchmarkAdd(f *os.File, benchmark *benchmarks.Benchmark) {

	hashSet := hashset.New[types.String]()
	linkedHashSet := linkedhashset.New[types.String]()
	treeSet := treeset.New[types.String]()

	sets := map[string]sets.Set[types.String]{
		"HashSet":       hashSet,
		"LinkedHashSet": linkedHashSet,
		"TreeSet":       treeSet,
	}
	for key, set := range sets {
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				set.Add(types.String(i))
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Add", r.Duration)
	}

}

func BenchmarkContains(f *os.File, benchmark *benchmarks.Benchmark) {

	hashSet := hashset.New[types.String]()
	linkedHashSet := linkedhashset.New[types.String]()
	treeSet := treeset.New[types.String]()

	sets := map[string]sets.Set[types.String]{
		"HashSet":       hashSet,
		"LinkedHashSet": linkedHashSet,
		"TreeSet":       treeSet,
	}

	for key, set := range sets {
		data.Initialize(set, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				set.Contains(types.String(fmt.Sprint(rand.Intn(set.Len()))))
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Contains", r.Duration)
	}
}

func BenchmarkRemove(f *os.File, benchmark *benchmarks.Benchmark) {

	hashSet := hashset.New[types.String]()
	linkedHashSet := linkedhashset.New[types.String]()
	treeSet := treeset.New[types.String]()

	sets := map[string]sets.Set[types.String]{
		"HashSet":       hashSet,
		"LinkedHashSet": linkedHashSet,
		"TreeSet":       treeSet,
	}

	for key, set := range sets {
		data.Initialize(set, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				set.Remove(types.String(fmt.Sprint(rand.Intn(benchmark.Setup().Size()))))
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Remove", r.Duration)
	}
}

func BenchmarkIterator(f *os.File, benchmark *benchmarks.Benchmark) {

	hashSet := hashset.New[types.String]()
	linkedHashSet := linkedhashset.New[types.String]()
	treeSet := treeset.New[types.String]()

	sets := map[string]sets.Set[types.String]{
		"HashSet":       hashSet,
		"LinkedHashSet": linkedHashSet,
		"TreeSet":       treeSet,
	}

	for key, set := range sets {
		data.Initialize(set, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			set.Iterator()
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Iterator", r.Duration)
	}
}

func Benchmark(f *os.File, benchmark *benchmarks.Benchmark) {
	BenchmarkAdd(f, benchmark)
	BenchmarkContains(f, benchmark)
	BenchmarkRemove(f, benchmark)
	BenchmarkIterator(f, benchmark)
}
