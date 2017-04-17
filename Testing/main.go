package main

import (
	"fmt"

	"github.com/dantrinh/cis193stats"
)

func main() {

	a, _ := cis193stats.Mean([]float64{1, 2, 3, 4, 5})
	fmt.Println(a) // 3

}
