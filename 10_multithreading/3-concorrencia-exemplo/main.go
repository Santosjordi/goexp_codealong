package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var count int64

// func main() {
// 	// this fuction is suposed to print a knock knock joke about race conditions
// 	// when the endppoint is called it should check the count variable,
// 	// if its value is equal to 0 it should print "knock knock" and increment the count by 1
// 	// if its value is equal to 1 it should print "who's there?" and increment the count by 1
// 	// if its value is equal to 2 it should print "race condition" and reset the count to 0
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(100 * time.Millisecond)
// 		if count == 0 {
// 			fmt.Println("knock knock")
// 			count++
// 		} else if count == 1 {
// 			fmt.Println("who's there?")
// 			count++
// 		} else {
// 			fmt.Println("race condition")
// 			count = 0
// 		}
// 	})
// 	http.ListenAndServe(":3000", nil)
// }

// run with -race flag
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(100 * time.Millisecond)
		if count == 0 {
			fmt.Println("knock knock")
			atomic.AddInt64(&count, 1)
		} else if count == 1 {
			fmt.Println("who's there?")
			atomic.AddInt64(&count, 1)
		} else {
			fmt.Println("race condition")
			atomic.StoreInt64(&count, 0)
		}
	})
	http.ListenAndServe(":3000", nil)
}
