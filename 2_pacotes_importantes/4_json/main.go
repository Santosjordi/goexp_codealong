package main

import (
	"encoding/json"
	"os"
)

type Account struct {
	Number  int     `json:"-"`
	Balance float64 `json:"balance"`
}

func main() {
	acc := Account{Number: 123, Balance: 100.50}
	res, err := json.Marshal(acc)
	if err != nil {
		println("Error:", err)
	}
	println(string(res))

	err = json.NewEncoder(os.Stdout).Encode(acc)
	if err != nil {
		println("Error:", err)
	}

	var contaX Account
	err = json.Unmarshal(res, &contaX)
	if err != nil {
		println("Error:", err)
	}
	println(contaX.Number, contaX.Balance)
}
