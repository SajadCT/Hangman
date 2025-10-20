package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Hangman struct {
	secretWord       string
	guessLetter      []byte
	correctGuesses   []byte
	remainingChances uint
}

func newHangman(secretWord string) Hangman {
	return Hangman{
		secretWord:       secretWord,
		guessLetter:      []byte{},
		correctGuesses:   []byte{},
		remainingChances: 7,
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
func checkGuess(state Hangman, guess byte) Hangman {
	if state.remainingChances > 1 && !bytes.Contains(state.guessLetter, []byte{guess}) {

		if strings.ContainsRune(state.secretWord, rune(guess)) { //if guess is correct
			state.correctGuesses = append(state.correctGuesses, guess)
			state.guessLetter = append(state.guessLetter, guess)

		} else { //if guess is wrong
			state.guessLetter = append(state.guessLetter, guess)
			state.remainingChances--
		}

	}
	return state
}

func isGameOver(state Hangman) bool {
	if hasWon(state) {
		return true
	}

	if state.remainingChances == 0 && len(state.guessLetter) == 7 {
		return true
	}

	return false
}

func hasWon(state Hangman) bool {
	return false

}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))

}
