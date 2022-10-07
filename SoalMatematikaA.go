// Soal Matematika A

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64
	var hasil float64

	fmt.Scan(&x)
	fmt.Scan(&y)

	hasil = 1/((3*(math.Pow(x, 2)))+10) + 10*y + 7

	fmt.Println(hasil)
}
