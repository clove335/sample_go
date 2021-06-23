package main

import (
	"fmt"
	. "lukechampine.com/uint128"
)

// fibonacci is a function that returns
// a function that returns an uint128.
func fibonacci() func() Uint128 {
	var i Uint128
	var a Uint128
	var b Uint128
	var c Uint128
	b = New(1, 0)
	c = a.Add(b)
	//c := a + b
	return func() Uint128 {
		var num Uint128
		if i.IsZero() {
			i = i.Add64(uint64(1))
		} else if i.Equals(New(1, 0)) {
			i = i.Add64(uint64(1))
			num = num.Add64(uint64(1))
		} else {
			num = c
			a = b
			b = c
			c = a.Add(b)
		}
		//fmt.Println(i, " ", a, " ", b, " ", c)
		return num
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(i, " ", f())
	}
}
