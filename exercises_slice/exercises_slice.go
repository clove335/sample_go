package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	dy_slices := make([][]uint8, dy)
	dx_slices := make([]uint8, dx)
	for i := range dy_slices {
		for j:= range dx_slices {
			dx_slices[j] = uint8((i^j))
		}
		dy_slices[i] = dx_slices
	}
	return dy_slices
}

func main() {
	pic.Show(Pic)
}
