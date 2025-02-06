package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("This is the system module:", uuid.New().String())
}
