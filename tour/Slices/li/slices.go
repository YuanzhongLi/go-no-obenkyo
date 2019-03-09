package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ret := make([][]uint8, dx)
	for ix := range ret {
		ret[ix] = make([]uint8, dy)

		for iy := range ret[ix] {
			ret[ix][iy] = uint8(ix * iy)
		}
	}
	return ret
}

func main() {
	pic.Show(Pic)
}
