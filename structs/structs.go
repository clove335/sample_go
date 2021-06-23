package main

import (
  "fmt"
  "math"
)

type Vertex struct {
    X int
    Y int
}

type BodyInfo struct {
    height int
    weight int
    bmi float64
}

func main() {
    fmt.Println(Vertex{30, 40})

    v := Vertex{50, 50}
    v.X = 90
    v.Y = 100
    fmt.Println(v.X*v.Y)

    // a pointer to structs
    p := &v
    fmt.Println(p.X+(*p).Y)

    b := BodyInfo{160, 50, .0}
    b.bmi = float64(b.weight)/math.Pow(float64(b.height)/100,2)
    fmt.Printf("Your height: %v, Your weight: %v,\nYour BMI is %v! However, be careful because BMI does not consider the proportion of FAT and PROTEIN.\n", b.height, b.weight, b.bmi)
}
