package fibonacci

// Fib returns fibonacci sequence
func Fibs(n int) []int {

	if n < 2 {
		return []int{n}
	}

	fibs := []int{0, 1}
	for i := 2; i < n; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2])

	}
	return fibs
}

// Fib returns the Nth Fibonacci number
func Fib(n int) int {

	if n < 2 {
		return n
	}

	fibs := []int{0, 1}
	for i := 2; i < n; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2])

	}
	return fibs[len(fibs)-1]
}
