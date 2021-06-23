package main

import "fmt"

const (
	n    = 200
	size = n + 1
)

var memo [size]uint64

func fibo(n uint64) uint64 {
	if memo[n] != 0 {
		return memo[n]
	}
	switch n {
	case 0:
		return 1
	case 1:
		return 1
	default:
		memo[n-1] = fibo(n - 1)
		memo[n-2] = fibo(n - 2)
		return memo[n-1] + memo[n-2]
	}
}

func main() {
	var i uint64
	for i = 0; i < n; i++ {
		fmt.Println(fibo(i))
	}
}
