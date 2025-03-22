// Starting to learn go routines
package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 50; i++ {
		fmt.Printf("Task: %s, Iteration: %d\n", name, i)
		time.Sleep(50 * time.Millisecond)
		wg.Done()
	}
}

// Thread 1
func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(150)
	// Thread 2
	go task("000", &waitGroup)
	// Thread 3
	go task("XXX", &waitGroup)
	// Thread 4
	go func() {
		for i := 0; i < 50; i++ {
			fmt.Printf("Task: %s, Iteration: %d\n", "Anonymous", i)
			time.Sleep(50 * time.Millisecond)
			waitGroup.Done()
		}
	}()
	// Se rodar apenas o main, o programa termina antes das go routines
	// Para esperar as go routines terminarem, podemos usar time.Sleep
	waitGroup.Wait()
}
