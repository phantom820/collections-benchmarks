// Package lists provides benchmark tests for lists.
package lists

import (
	"collections-benchmarks/benchmarks"
	"collections-benchmarks/data"
	"fmt"
	"math/rand"
	"os"

	"github.com/phantom820/collections/lists"
	"github.com/phantom820/collections/lists/forwardlist"
	"github.com/phantom820/collections/lists/list"
	"github.com/phantom820/collections/lists/vector"
	"github.com/phantom820/collections/sort"

	"github.com/phantom820/collections/types"
)

func BenchmarkAdd(f *os.File, benchmark *benchmarks.Benchmark) {

	forwardList := forwardlist.New[types.String]()
	list := list.New[types.String]()
	vector := vector.New[types.String]()

	lists := map[string]lists.List[types.String]{
		"ForwardList": forwardList,
		"List":        list,
		"Vector":      vector,
	}

	for key, list := range lists {
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				list.Add(types.String(fmt.Sprint(i)))
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Add", r.Duration)
	}

}

func BenchmarkContains(f *os.File, benchmark *benchmarks.Benchmark) {

	forwardList := forwardlist.New[types.String]()
	list := list.New[types.String]()
	vector := vector.New[types.String]()

	lists := map[string]lists.List[types.String]{
		"ForwardList": forwardList,
		"List":        list,
		"Vector":      vector,
	}

	for key, list := range lists {
		data.Initialize(list, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				list.Contains(types.String(fmt.Sprint(rand.Intn(list.Len()))))
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Contains", r.Duration)
	}
}

func BenchmarkRemove(f *os.File, benchmark *benchmarks.Benchmark) {

	forwardList := forwardlist.New[types.String]()
	list := list.New[types.String]()
	vector := vector.New[types.String]()

	lists := map[string]lists.List[types.String]{
		"ForwardList": forwardList,
		"List":        list,
		"Vector":      vector,
	}

	for key, list := range lists {
		data.Initialize(list, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				list.Remove(types.String(fmt.Sprint(rand.Intn(benchmark.Setup().Size()))))
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Remove", r.Duration)
	}
}

func BenchmarkIterator(f *os.File, benchmark *benchmarks.Benchmark) {

	forwardList := forwardlist.New[types.String]()
	list := list.New[types.String]()
	vector := vector.New[types.String]()

	lists := map[string]lists.List[types.String]{
		"ForwardList": forwardList,
		"List":        list,
		"Vector":      vector,
	}

	for key, list := range lists {
		data.Initialize(list, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			list.Iterator()
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Iterator", r.Duration)
	}

}

func BenchmarkReverse(f *os.File, benchmark *benchmarks.Benchmark) {

	forwardList := forwardlist.New[types.String]()
	list := list.New[types.String]()
	vector := vector.New[types.String]()

	lists := map[string]lists.List[types.String]{
		"ForwardList": forwardList,
		"List":        list,
		"Vector":      vector,
	}

	for key, list := range lists {
		data.Initialize(list, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			list.Reverse()
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Reverse", r.Duration)
	}
}

func BenchmarkRemoveBack(f *os.File, benchmark *benchmarks.Benchmark) {

	forwardList := forwardlist.New[types.String]()
	list := list.New[types.String]()
	vector := vector.New[types.String]()

	lists := map[string]lists.List[types.String]{
		"ForwardList": forwardList,
		"List":        list,
		"Vector":      vector,
	}
	for key, list := range lists {
		data.Initialize(list, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				list.RemoveBack()
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "RemoveBack", r.Duration)
	}
}

func BenchmarkRemoveAt(f *os.File, benchmark *benchmarks.Benchmark) {

	forwardList := forwardlist.New[types.String]()
	list := list.New[types.String]()
	vector := vector.New[types.String]()

	lists := map[string]lists.List[types.String]{
		"ForwardList": forwardList,
		"List":        list,
		"Vector":      vector,
	}

	for key, list := range lists {
		data.Initialize(list, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				list.RemoveAt(rand.Intn(list.Len()))
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "RemoveAt", r.Duration)
	}

}

func BenchmarkSort(f *os.File, benchmark *benchmarks.Benchmark) {

	forwardList := forwardlist.New[types.String]()
	list := list.New[types.String]()
	vector := vector.New[types.String]()

	lists := map[string]lists.List[types.String]{
		"ForwardList": forwardList,
		"List":        list,
		"Vector":      vector,
	}
	for key, list := range lists {
		data.InitializeReverse(list, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			sort.Sort[types.String](list)
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Sort", r.Duration)
	}
}

func Benchmark(f *os.File, benchmark *benchmarks.Benchmark) {

	BenchmarkAdd(f, benchmark)
	BenchmarkContains(f, benchmark)
	BenchmarkRemove(f, benchmark)
	BenchmarkRemoveBack(f, benchmark)
	BenchmarkRemoveAt(f, benchmark)
	BenchmarkIterator(f, benchmark)
	BenchmarkReverse(f, benchmark)
	BenchmarkSort(f, benchmark)

}
