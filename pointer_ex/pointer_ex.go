package main

import "fmt"

func main() {
    i, j := 42, 2432
    p := &i
    fmt.Println(*p)
    *p /= 2
    fmt.Println(i)

    p = &j
    *p /= 16
    fmt.Println(j)
}
