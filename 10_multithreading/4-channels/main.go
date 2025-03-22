package main

import (
	"fmt"
	"sync"
	"time"
)

// func main() {
// 	canal := make(chan int)
// 	go publish(canal)
// 	reader(canal)
// }

// // this causes a deadlock
// func deadlock() {
// 	forever := make(chan bool)
// 	<-forever
// }

// func publish(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }

// func reader(ch chan int) {
// 	for x := range ch {
// 		println(x)
// 	}
// }

// DnD Analogy

func player(id int, turnOrder chan int, wg *sync.WaitGroup, rounds *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Player %d started and is waiting for a turn...\n", id)

	for turn := range turnOrder {
		fmt.Printf("Player %d: Received turn value %d\n", id, turn)

		if turn == id {
			fmt.Printf("Player %d is taking their turn!\n", id)
			time.Sleep(time.Second) // Simulating action
			fmt.Printf("Player %d ends their turn.\n", id)

			nextPlayer := (id % 3) + 1
			fmt.Printf("Player %d: Passing turn to Player %d\n", id, nextPlayer)
			turnOrder <- nextPlayer // Pass the turn forward

			rounds.Done() // Mark one round as done
		}
	}

	fmt.Printf("Player %d: Exiting since channel is closed.\n", id)
}

func main() {
	turnOrder := make(chan int, 1) // Buffered so we can start with 1 value
	var wg sync.WaitGroup
	var rounds sync.WaitGroup
	numPlayers := 3
	numRounds := 10

	// Start players
	for i := 1; i <= numPlayers; i++ {
		wg.Add(1)
		go player(i, turnOrder, &wg, &rounds)
	}

	// Set the number of rounds
	rounds.Add(numRounds)

	// Start the game with Player 1's turn
	turnOrder <- 1

	go func() {
		rounds.Wait()
		close(turnOrder) // Close the channel after the rounds are done
	}()

	wg.Wait() // Wait for all players to finish

	fmt.Println("Combat Ends!")
}
