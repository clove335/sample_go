package main

import (
	"fmt"
	"math"
)

// Go's basic types
// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr
// byte == alias for uint8
// rune == alias for int32
// float32 float64
// complex64 complex128
//
// Zero values in Go
// 1. 0 for numerics
// 2. false for bool
// 3. "" == empty stirng for strings
//fail_short_assign := "failed!"

//func mul(x int, y int) int {
func mul(x, y int) int {
	l := 8
	return x * y * l
}

func swap(x, y string) (string, string) {
	return y, x
}

var c, python, java bool

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(math.Pi)
	fmt.Println(mul(3, 9))

	a, b := swap("hello", "my world,")
	fmt.Println(a, b)

	var i int
	fmt.Println(i, c, python, java)

	var ccc, pypy, javajava = true, false, "Java!"
	fmt.Println(ccc, pypy, javajava)
	k := 3.3
	fmt.Printf("k(%f) is now type %T\n", k, k)

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("finished!")

	defer fmt.Println("defer after loop finished!")
	fmt.Println("end of function")
}
