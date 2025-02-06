package main

import (
	"fmt"

	"github.com/santosjordi/posgoexp/7-packaging/math"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Result", math.Math{A: 2, B: 3}.Add())
}
