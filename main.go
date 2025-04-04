package main

import "markov-chain/config"

func main() {
	config.ParseFlags()
	err := config.HandleConfig(&config.Cfg)
}