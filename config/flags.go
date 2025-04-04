package config

import (
	"flag"
	"fmt"
)

type Config struct {
	MaxWords int
	StartPrefix string
	PrefixLength int
}

var Cfg Config

func ParseFlags() {
	flag.IntVar(&Cfg.MaxWords, "w", 100, "Maximum words to print")
	flag.StringVar(&Cfg.StartPrefix, "p", "", "Starting prefix")
	flag.IntVar(&Cfg.PrefixLength, "l", 2, "Starting prefix length")

	flag.Usage = func() {
		PrintHelp()
	}
	
	flag.Parse()
}

func PrintHelp() {
	fmt.Println(`Markov Chain text generator.

Usage:
  markovchain [-w <N>] [-p <S>] [-l <N>]
  markovchain --help

Options:
  --help  Show this screen.
  -w N    Number of maximum words
  -p S    Starting prefix
  -l N    Prefix length`)
}