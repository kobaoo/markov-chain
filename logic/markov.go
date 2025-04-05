package logic

import (
	"fmt"
	"io"
	"markov-chain/config"
	"math/rand"
	"os"
	"strings"
)

var MarkovChain = make(map[string][]string)

func GetChain() error {
	fileInfo, err := os.Stdin.Stat()
	if err != nil {
		return fmt.Errorf("failed to access stdin: %s", err)
	}
	if fileInfo.Mode()&os.ModeNamedPipe == 0 {
		return fmt.Errorf("no input text")
	}

	buf := new(strings.Builder)
	io.Copy(buf, os.Stdin)
	input := buf.String()
	text := strings.Fields(input)

	prefixLength := config.Cfg.PrefixLength

	if len(text) < 2 || len(text) < prefixLength{
		return fmt.Errorf("not enough data to use")
	}

	if strings.TrimSpace(config.Cfg.StartPrefix) == "" {
		config.Cfg.StartPrefix = strings.Join(text[:prefixLength], " ")
	} else {

	}
	
	i := 0
	for ; i < len(text)-prefixLength; i++ {
		prefix := JoinWithSpace(text[i:i+prefixLength])
		MarkovChain[prefix] = append(MarkovChain[prefix], text[i+prefixLength])
	}
	prefix := JoinWithSpace(text[i:i+prefixLength])
	MarkovChain[prefix] = append(MarkovChain[prefix], "!@#@!#") //end
	return nil
}

func PrintMarkovChainText() {
	maxWords := config.Cfg.MaxWords
	prefix := config.Cfg.StartPrefix
	prefixLength := config.Cfg.PrefixLength
	fmt.Print(prefix)
	
	for i := 0; i < maxWords - prefixLength; i++ {
		next, ok := MarkovChain[prefix]
		if !ok {
			return
		}
		word := next[rand.Intn(len(next))]
		updatePrefix(&prefix, word)
		fmt.Print(" ", word)
	}
}