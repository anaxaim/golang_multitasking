package main

import (
	"fmt"
	"strings"
	"unicode"
)

// nextFunc returns the next word from the generator
type nextFunc func() string

// counter stores the number of digits in each word.
// each key is a word and value is the number of digits in the word.
type counter map[string]int

// pair stores a word and the number of digits in it
type pair struct {
	word  string
	count int
}

// {
// 	out := make(chan int)
// 	done := make(chan struct{})

// 	func work(done chan struct{}, out chan int) {
// 		out <- 42
// 		done <- struct{}{}
// 	}

// 	go work(done, out) // (1)

// 	go func() { // (2)
// 		<-done
// 		fmt.Println("work done")
// 		done <- struct{}{}
// 	}()

// 	fmt.Println(<-out) // (3)
// 	<-done
// 	fmt.Println("all goroutines done")
// }

// countDigitsInWords counts digits in words,
// fetching each word with the next() function
func countDigitsInWords(next nextFunc) counter {
	pending := make(chan string)
	go submitWords(next, pending)

	done := make(chan struct{})
	counted := make(chan pair)

	for i := 0; i < 4; i++ {
		go countWords(done, pending, counted)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-done
		}
		close(counted)
	}()

	return fillStats(counted)
}

// submitWords отправляет слова на подсчет
func submitWords(next nextFunc, out chan<- string) {
	for {
		word := next()
		if word == "" {
			break
		}
		out <- word
	}
	close(out)
}

// countWords считает цифры в словах
func countWords(done chan<- struct{}, in <-chan string, out chan<- pair) {
	// реализуйте логику подсчета цифр
	// с использованием каналов done, in и out

	for word := range in {
		count := countDigits(word)
		out <- pair{word, count}
	}
	done <- struct{}{}
}

// fillStats готовит итоговую статистику
func fillStats(in <-chan pair) counter {
	stats := counter{}
	for p := range in {
		stats[p.word] = p.count
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

// printStats prints words and their digit counts
func printStats(stats counter) {
	for word, count := range stats {
		fmt.Printf("%s: %d\n", word, count)
	}
}

// wordGenerator returns a generator,
// which emits words from a phrase.
func wordGenerator(phrase string) nextFunc {
	words := strings.Fields(phrase)
	idx := 0
	return func() string {
		if idx == len(words) {
			return ""
		}
		word := words[idx]
		idx++
		return word
	}
}

func main() {
	phrase := "1 22 333 4444 55555 666666 7777777 88888888"
	next := wordGenerator(phrase)
	stats := countDigitsInWords(next)
	printStats(stats)
}
