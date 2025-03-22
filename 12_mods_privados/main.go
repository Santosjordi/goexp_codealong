package main

import (
	"fmt"

	"github.com/Santosjordi/goexp_codealong/11_eventos/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
