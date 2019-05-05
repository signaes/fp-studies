package main

type memoized func(int) int

func main() {
	println("fib(5) ", fib(5))
	println("fib(3) ", fib(3))
	println("fib(1) ", fib(1))
	println("fib(2) ", fib(2))
	println("fib(4) ", fib(4))
	println("fib(2) + fib(3) = fib(4) | ", fib(2), " + ", fib(3), " = ", fib(4))
	println("fib(10) ", fib(10))
	println("fib(20) ", fib(20))
	println("fib(30) ", fib(30))
	println("fib(40) ", fib(40))
	println("fib(41) ", fib(41))
	println("fib(42) ", fib(42))
	println("memoizedFib(5) ", memoizedFib(5))
	println("memoizedFib(10) ", memoizedFib(10))
	println("memoizedFib(20) ", memoizedFib(20))
	println("memoizedFib(30) ", memoizedFib(30))
	println("memoizedFib(40) ", memoizedFib(40))
	println("memoizedFib(41) ", memoizedFib(41))
	println("memoizedFib(42) ", memoizedFib(42))
	println("memoizedFib(43) ", memoizedFib(43))
	println("memoizedFib(40) ", memoizedFib(40))
	println("memoizedFib(41) ", memoizedFib(41))
	println("memoizedFib(42) ", memoizedFib(42))
	println("memoizedFib(43) ", memoizedFib(43))
}

func memoize(mf memoized) memoized {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		result := mf(key)
		cache[key] = result
		return result
	}
}

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return fib(x-2) + fib(x-1)
	}
}

var fibMem = memoize(memoized(fib))

func memoizedFib(x int) int {
	return fibMem(x)
}
