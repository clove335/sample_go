package main

import "fmt"
import "math"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	n := 10
	for i := 0; i < n; i++ {
		fmt.Print("this is Print function. ", i, "\n")
	}
	fmt.Print("use some functions from math library.\n")
	for i := -10; i < n; i++ {
		fmt.Println("Before abs:", i, " After abs:", math.Abs(float64(i)))
	}
}
