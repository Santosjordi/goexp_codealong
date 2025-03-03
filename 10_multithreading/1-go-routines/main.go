// Starting to learn go routines
package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Task: %s, Iteration: %d\n", name, i)
		time.Sleep(500 * time.Millisecond)
	}
}

// Thread 1
func main() {
	// Thread 2
	go task("Learning Go Routines")
	// Thread 3
	go task("This task is running simultaneously")
	// Se rodar apenas o main, o programa termina antes das go routines
	// Para esperar as go routines terminarem, podemos usar time.Sleep
	time.Sleep(5 * time.Second)
}
