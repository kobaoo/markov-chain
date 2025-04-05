package main

import (
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
	
	// for key, value := range logic.MarkovChain {
	// 	fmt.Println("key:", key, " | Value: ", value)
	// }

	logic.PrintMarkovChainText()
}