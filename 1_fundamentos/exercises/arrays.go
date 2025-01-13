package main

import "fmt"

var myArray = [5]int{1, 2, 3, 4, 5}

func main() {
	for i, v := range myArray {
		fmt.Printf("Index %d, value %d\n", i, v)
	}
}
