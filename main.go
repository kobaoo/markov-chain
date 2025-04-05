package main

import (
	"fmt"
	"log"
	"markov-chain/config"
	"markov-chain/logic"
)

func main() {
	config.ParseFlags()
	err := config.HandleConfig(&config.Cfg)
	if err != nil {
		log.Fatal("Error: ", err)
	}
	err = logic.GetChain()
	if err != nil {
		log.Fatal("Error: ", err)
	}

	logic.PrintMarkovChainText()
	fmt.Println()
}