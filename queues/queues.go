package queues

import (
	"collections-benchmarks/benchmarks"
	"collections-benchmarks/data"
	"fmt"
	"os"

	"github.com/phantom820/collections/queues"
	"github.com/phantom820/collections/queues/listdequeue"
	"github.com/phantom820/collections/queues/listqueue"
	"github.com/phantom820/collections/queues/slicedequeue"
	"github.com/phantom820/collections/queues/slicequeue"
	"github.com/phantom820/collections/types"
)

func BenchmarkFront(f *os.File, benchmark *benchmarks.Benchmark) {

	listQueue := listqueue.New[types.String]()
	sliceQueue := slicequeue.New[types.String]()
	listDequeue := listdequeue.New[types.String]()
	sliceDequeue := slicedequeue.New[types.String]()

	queues := map[string]queues.Queue[types.String]{
		"ListQueue":    listQueue,
		"SliceQueue":   sliceQueue,
		"ListDequeue":  listDequeue,
		"SliceDequeue": sliceDequeue,
	}

	for key, queue := range queues {
		data.Initialize(queue, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				queue.Front()
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Front", r.Duration)
	}

}

func BenchmarkRemoveFront(f *os.File, benchmark *benchmarks.Benchmark) {

	listQueue := listqueue.New[types.String]()
	sliceQueue := slicequeue.New[types.String]()
	listDequeue := listdequeue.New[types.String]()
	sliceDequeue := slicedequeue.New[types.String]()

	queues := map[string]queues.Queue[types.String]{
		"ListQueue":    listQueue,
		"SliceQueue":   sliceQueue,
		"ListDequeue":  listDequeue,
		"SliceDequeue": sliceDequeue,
	}

	for key, queue := range queues {
		data.Initialize(queue, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < queue.Len(); i++ {
				queue.RemoveFront()
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "RemoveFront", r.Duration)
	}

}

func BenchmarkBack(f *os.File, benchmark *benchmarks.Benchmark) {

	listDequeue := listdequeue.New[types.String]()
	sliceDequeue := slicedequeue.New[types.String]()

	queues := map[string]queues.Dequeue[types.String]{
		"ListDequeue":  listDequeue,
		"SliceDequeue": sliceDequeue,
	}

	for key, queue := range queues {
		data.Initialize(queue, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				queue.Back()
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Back", r.Duration)
	}
}

func BenchmarkRemoveBack(f *os.File, benchmark *benchmarks.Benchmark) {

	listDequeue := listdequeue.New[types.String]()
	sliceDequeue := slicedequeue.New[types.String]()

	queues := map[string]queues.Dequeue[types.String]{
		"ListDequeue":  listDequeue,
		"SliceDequeue": sliceDequeue,
	}

	for key, queue := range queues {
		data.Initialize(queue, benchmark.Setup().Size())
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				queue.RemoveBack()
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "RemoveBack", r.Duration)
	}
}

func BenchmarkAddFront(f *os.File, benchmark *benchmarks.Benchmark) {

	listDequeue := listdequeue.New[types.String]()
	sliceDequeue := slicedequeue.New[types.String]()

	queues := map[string]queues.Dequeue[types.String]{
		"ListDequeue":  listDequeue,
		"SliceDequeue": sliceDequeue,
	}

	for key, queue := range queues {
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				queue.AddFront(types.String(fmt.Sprint(i)))
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "AddFront", r.Duration)
	}
}

func BenchmarkAdd(f *os.File, benchmark *benchmarks.Benchmark) {

	listQueue := listqueue.New[types.String]()
	sliceQueue := slicequeue.New[types.String]()
	listDequeue := listdequeue.New[types.String]()
	sliceDequeue := slicedequeue.New[types.String]()

	queues := map[string]queues.Queue[types.String]{
		"ListQueue":    listQueue,
		"SliceQueue":   sliceQueue,
		"ListDequeue":  listDequeue,
		"SliceDequeue": sliceDequeue,
	}

	for key, queue := range queues {
		r := benchmark.Benchmark(key, func() {
			for i := 0; i < benchmark.Setup().Operations(); i++ {
				queue.Add(types.String(fmt.Sprint(i)))
			}
		})
		fmt.Fprintf(f, "%s,%s,%v\n", r.Name, "Add", r.Duration)
	}
}

func Benchmark(f *os.File, benchmark *benchmarks.Benchmark) {
	BenchmarkFront(f, benchmark)
	BenchmarkRemoveFront(f, benchmark)
	BenchmarkAddFront(f, benchmark)
	BenchmarkAdd(f, benchmark)
	BenchmarkRemoveBack(f, benchmark)
}
