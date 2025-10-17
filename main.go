package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Hangman struct {
	secretWord       string
	guessLetter      []byte
	correctGuesses   []byte
	remainingChanses uint
}

func newHangman(secretWord string) Hangman {
	return Hangman{
		secretWord:       secretWord,
		guessLetter:      []byte{},
		correctGuesses:   []byte{},
		remainingChanses: 7,
	}

}

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
func getNewState(state Hangman, userInput string) Hangman {

	return state
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))

}
