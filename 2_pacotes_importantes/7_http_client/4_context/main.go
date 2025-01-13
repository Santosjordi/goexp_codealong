package main

import (
	"context"
	"io"
	"net/http"
	"time"
)

// O pacote "context" é utilizado para cancelar requests

func main() {
	ctx := context.Background()
	// o pacote permite cancelar uma execução de acordo com um critério definido
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
