package main

// Importar esta libreria con go get
import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	res := [][]uint8{}
	for i := 0; i < dy; i++ {
		ren := []uint8{}
		for j := 0; j < dx; j++ {
			ren = append(ren, uint8((j ^ i)))
		}
		res = append(res, ren)
	}
	return res
}

func slides() {
	pic.Show(Pic)

}
