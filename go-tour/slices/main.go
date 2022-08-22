package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	var x, y int = 0, 0
	b := make([][]uint8, dy, dy)
	for y < dy {

		b[y] = make([]uint8, dx, dx)
		for x < dx {
			b[y][x] = uint8(y ^ x)
			x = x + 1
		}
		y = y + 1
		x = 0
	}

	return b
}

func main() {
	pic.Show(Pic)
}
