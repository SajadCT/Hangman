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
	if state.remainingChances > 0 && !bytes.Contains(state.guessLetter, []byte{guess}) {

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

	if state.remainingChances == 0 {
		return true
	}

	return false
}

func hasWon(state Hangman) bool {

	correctLetters := make(map[byte]bool)
	for i := 0; i < len(state.secretWord); i++ {
		correctLetters[state.secretWord[i]] = true
	}

	for letter := range correctLetters {
		if !bytes.Contains(state.correctGuesses, []byte{letter}) {
			return false
		}
	}
	return true

}

func getUserInput() byte {
	fmt.Print("Enter a Character:")
	reader := bufio.NewReader(os.Stdin)
	char, _ := reader.ReadByte()
	reader.ReadByte()

	if char >= 'A' && char <= 'Z' {
		char = char + 32
	}
	return char
}

func displaySecretWord(state Hangman) string {
	var word strings.Builder

	for _, ch := range state.secretWord {
		letter := byte(ch)
		found := false
		if bytes.Contains(state.correctGuesses, []byte{letter}) {
			found = true
		}
		if found {
			word.WriteByte(letter)
		} else {
			word.WriteByte('-')
		}
	}

	return word.String()

}

func main() {

	secretWord := getSecretWord("/usr/share/dict/words")
	game := newHangman(secretWord)
	fmt.Println("WELCOME TO HANGMAN!")
	fmt.Println(secretWord)

	for !isGameOver((game)) {
		fmt.Println("Hidden Word:", displaySecretWord(game))
		fmt.Println("Chancess Remaining:", game.remainingChances)
		fmt.Println("Guessed Letters : ", string(game.guessLetter))
		guess := getUserInput()
		if guess < 'a' || guess > 'z' {
			fmt.Println("Please enter a valid character....!!!")
			fmt.Println()
			continue
		}

		game = checkGuess(game, guess)

		if isGameOver(game) {
			if hasWon(game) {
				fmt.Println("You have won the game! The word is:", secretWord)
			} else {
				fmt.Println("Sorry! You loose. the secret word is:", secretWord)
			}
		}

	}
	fmt.Println()
}
