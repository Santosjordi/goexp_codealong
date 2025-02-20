package main

import "github.com/santosjordi/posgoexp/9_apis/configs"

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config.DBDriver)
}
