package main

import (
	"fmt"
	"strings"
	"unicode"
)

// counter stores the number of digits in each word.
// each key is a word and value is the number of digits in the word.
type counter map[string]int

// countDigitsInWords counts digits in phrase words
func countDigitsInWords(phrase string) counter {
	words := strings.Fields(phrase)
	counted := make(chan int)

	go func() {
		for _, word := range words {
			count := countDigits(word)
			counted <- count
		}
	}()

	stats := counter{}
	for _, word := range words {
		count := <-counted
		stats[word] = count
	}

	return stats
}

// countDigits returns the number of digits in a string
func countDigits(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsDigit(char) {
			count++
		}
	}
	return count
}

// printStats prints words and their digit stats
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func main() {
	phrase := "0ne 1wo thr33 4068"
	stats := countDigitsInWords(phrase)
	printStats(stats)
}
