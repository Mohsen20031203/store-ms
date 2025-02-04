package main

import "hello/config"

func main() {
	_, err := config.LoadConfig(".")
	if err != nil {
		return
	}
}
