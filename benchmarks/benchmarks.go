package benchmarks

import "time"

type Timer struct {
	start time.Time
	end   time.Time
}

func (timer *Timer) Start() {
	timer.start = time.Now()
}

func (timer *Timer) Stop() {
	timer.end = time.Now()
}

func (timer *Timer) Duration() float32 {
	return float32(timer.end.Sub(timer.start).Nanoseconds())
}

type Setup struct {
	count  int // number of measurements to aggregate.
	trials int // number of samples needed to produce a single measurement.
}

func NewSetup(trials int, count int) *Setup {
	return &Setup{trials: trials, count: count}
}

type Benchmark struct {
	setup *Setup
	timer *Timer
}

func NewBenchmark(setup *Setup, timer *Timer) *Benchmark {
	return &Benchmark{setup: setup, timer: timer}
}

func (benchmark *Benchmark) Benchmark(name string, f func()) Result {

	trial := func() float32 {
		sample := float32(0.0)
		for j := 0; j < benchmark.setup.trials; j++ {
			benchmark.timer.Start()
			f()
			benchmark.timer.Stop()
			sample += (benchmark.timer.Duration())
		}
		return sample / float32(benchmark.setup.trials)
	}

	duration := float32(0.0)
	for i := 0; i < benchmark.setup.count; i++ {
		duration += trial()
	}
	return Result{Name: name, Duration: duration / float32(benchmark.setup.count)}
}

type Result struct {
	Name     string
	Duration float32
}
