package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var sl [][]uint8
	
	for i := 0; i < dy ; i++ {
		tmp := []uint8{}
		sl = append(sl, tmp)
		for j := 0; j < dx; j++ {
		    sl[i] = append(sl[i], uint8(i^j))
		}
	}

	return sl
}

func main() {
	pic.Show(Pic)
}
