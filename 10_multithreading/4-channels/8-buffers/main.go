package main

// main demonstrates the use of buffered channels in Go.
// Imagine a parking lot with 2 parking spaces (buffer size of 2).
// Cars (messages) can enter the parking lot (channel) if there is space available.
// Here, we park two cars ("hello" and "world") in the parking lot.
// Then, we retrieve and print the cars in the order they were parked.
func main() {
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"

	println(<-ch)
	println(<-ch)
}
