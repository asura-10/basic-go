package testbenchmark

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Benchmark() {
	sum := 0
	for i := 1; i < 20; i++ {
		sum += i
	}
}
