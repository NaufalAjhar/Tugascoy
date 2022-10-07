//Soal Tiga Digit

package main

import (
	"fmt"
)

func main() {
	var x int64
	var d1, d2, d3 int64

	fmt.Scan(&x)
	if x <= 999 {
		d1 = (x / 100 % 10)
		d2 = (x / 10 % 10)
		d3 = (x / 1 % 10)
		println(d1, d2, d3)
	}
}
