package fibonacci

// Fib returns fibonacci sequence
func Fibs(n uint64) []uint64 {

	if n < 2 {
		return []uint64{n}
	}

	fibs := []uint64{0, 1}
	var i uint64
	for i = 2; i < n; i++ {
		fibs = append(fibs, fibs[i-1]+fibs[i-2])

	}
	return fibs
}

// Fib returns the Nth Fibonacci number
// O(n) -- time complexity
// O(n) -- space complexity (array)
func Fib(n uint64) uint64 {

	if n < 2 {
		return n
	}

	fibs := []uint64{0, 1} // <-- Memoization
	var i uint64
	for i = 2; i < n; i++ { //<-- time complexity is O(n)
		fibs = append(fibs, fibs[i-1]+fibs[i-2])

	}
	return fibs[len(fibs)-1]
}

// FibSmaller returns the Nth Fibonacci number
// O(n) -- time complexity
// O(1) -- space complexity single vars
func FibSmaller(n uint64) uint64 {

	if n < 2 {
		return n
	}

	var i, twice, current uint64
	var last uint64 = 1

	for i = 2; i < n; i++ {
		current = twice + last
		twice = last
		last = current
	}
	return current
}

// FibRecurse
// O(2^n) -- time complexity
// O(2^n) -- space complexity
// This is the worst way to solve... huge space and complexity
func FibRecurse(n uint64) uint64 {

	if n < 2 {
		return n
	}

	return FibRecurse(n-1) + FibRecurse(n-2)
}

// O(n) -- time
// O(1) -- space (no stack)
//FibTail Recursive
func FibTail(n int) int {
	return fibT(n, 0, 1)
}

func fibT(n, first, second int) int {
	if n == 0 {
		return first
	}
	return fibT(n-1, second, first+second)
}
