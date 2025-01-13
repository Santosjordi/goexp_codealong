package main

import (
	"io"
	"net/http"
)

// Como customizar uma request antes de executar

func main() {
	c := http.Client{}
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accept", "application/json")

	// os passos anteriores são necessários para configurar o objeto request, uma vez que
	// todas os dados e configurações são definidor, o objeto é passado para o client
	resp, err := c.Do(req)
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
