//Soal Matematika B

package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var hasil float64

	fmt.Scan(&x)

	hasil = (math.Pow(x, 3) + 3*x) / (math.Pow(x, 4) - 3*math.Pow(x, 2) + 4)

	fmt.Println(hasil)
}
