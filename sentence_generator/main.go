package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Creator: Will Hegedus\nSentence Builder\n\n")
	sentences := generateSetences(20)
	for i := 0; i < len(sentences); i++ {
		fmt.Println(sentences[i])
	}
}
