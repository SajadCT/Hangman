package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func containsPunctuation(s string) bool {
	for _, ch := range s {
		if ch < 'a' || ch > 'z' {
			return true
		}
	}
	return false
}

func getSecretWord(wordFileName string) string {
	var allowedWords []string

	file, err := os.Open(wordFileName)
	if err != nil {
		panic(fmt.Sprintf("Error in %v cause of %v", wordFileName, err))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := scanner.Text()
		if word == strings.ToLower(word) && len(word) >= 6 && !containsPunctuation(word) {
			allowedWords = append(allowedWords, word)
		}
	}
	randomNum := rand.Intn(len(allowedWords))
	return allowedWords[randomNum]

}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))

}
