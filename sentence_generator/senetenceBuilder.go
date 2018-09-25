package main

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	// slices
	articles     = []string{"the", "a", "one", "some", "any"}
	nouns        = []string{"boy", "girl", "dog", "town", "car"}
	verbs        = []string{"drove", "jumped", "ran", "walked", "skipped"}
	prepositions = []string{"to", "from", "over", "under", "on"}
)

func getRandomInt() int {
	// need to sleep or it runs too fast & you get the same words
	time.Sleep(3)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(5)
}

func getArticle() string {
	index := getRandomInt()
	return articles[index]
}

func getNoun() string {
	index := getRandomInt()
	return nouns[index]
}

func getVerb() string {
	index := getRandomInt()
	return verbs[index]
}

func getPreposition() string {
	index := getRandomInt()
	return prepositions[index]
}

func compileSentence() string {
	return strings.Title(getArticle()) + " " +
		getNoun() + " " +
		getVerb() + " " +
		getPreposition() + " " +
		getArticle() + " " +
		getNoun() + "."
}

func generateSetences(sentences int) []string {
	compiledSentences := make([]string, sentences)
	for i := 0; i < sentences; i++ {
		compiledSentences[i] = strconv.Itoa(i+1) + ": " + compileSentence()
	}
	return compiledSentences
}
