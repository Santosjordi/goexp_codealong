package main

import (
	"fmt"
	"time"
)

func worker(wokerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d: Received %d\n", wokerId, x)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	data := make(chan int)

	amountOfWorkers := 300

	for i := 0; i < amountOfWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 1000; i++ {
		data <- i
	}
}

// // this processes serially
// func worker(wokerId int, data chan int) {
// 	for x := range data {
// 		fmt.Printf("Worker %d: Received %d\n", wokerId, x)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	data := make(chan int)
// 	go worker(1, data)
// 	for i := 0; i < 100; i++ {
// 		data <- i
// 	}
// }
