package main

import "app/internal/config"

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
}
