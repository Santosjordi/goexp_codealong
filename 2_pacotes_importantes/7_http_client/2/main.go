package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1 * time.Second)}
	jsonVar := bytes.NewBuffer([]byte(`{"name":"John"}`))
	resp, err := c.Post("http://santosfoundry.duckdns.org", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
