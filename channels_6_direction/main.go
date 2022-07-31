package main

import (
	"fmt"
	"strings"
)

func encode(str string) string {

	submitter := func(str string) <-chan string {
		out := make(chan string)
		go func() {
			words := strings.Fields(str)
			for _, word := range words {
				out <- word
			}
			close(out)
		}()
		return out
	}

	encoder := func(in <-chan string) <-chan string {
		out := make(chan string)
		go func() {
			for word := range in {
				out <- encodeWord(word)
			}
			close(out)
		}()
		return out
	}

	receiver := func(in <-chan string) []string {
		words := []string{}
		for word := range in {
			words = append(words, word)
		}
		return words
	}

	pending := submitter(str)
	encoded := encoder(pending)
	words := receiver(encoded)
	return strings.Join(words, " ")
}

func encodeWord(word string) string {
	const shift = 13
	const char_a byte = 'a'
	encoded := make([]byte, len(word))
	for idx, char := range []byte(word) {
		delta := (char - char_a + shift) % 26
		encoded[idx] = char_a + delta
	}
	return string(encoded)
}

func main() {
	src := "go is awesome"
	res := encode(src)
	fmt.Println(res)
}
