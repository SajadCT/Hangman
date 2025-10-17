package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

func createDictFile(words []string) (string, error) {
	f, err := os.CreateTemp("/tmp", "Hangman-dict")
	if err != nil {
		fmt.Println("can't create temporary file.")
	}
	data := strings.Join(words, "\n")
	_, err = f.Write([]byte(data))
	if err != nil {
		return "", err
	}
	return f.Name(), nil

}

func TestSecretWordNoCapitals(t *testing.T) {
	wordList, err := createDictFile([]string{"Lion", "Elephant", "monkey"})
	defer os.Remove(wordList)
	if err != nil {

		t.Errorf("Couldn't create word list. Can't proceed with test : %v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should get 'monkey' but Got %s", secretWord)
	}
}

func TestSecretWordLength(t *testing.T) {
	wordList, err := createDictFile([]string{"lion", "pen", "monkey"})
	defer os.Remove(wordList)
	if err != nil {

		t.Errorf("Couldn't create word list. Can't proceed with test : %v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should get 'monkey' but Got %s", secretWord)
	}
}

func TestSecretWordPunctuations(t *testing.T) {
	wordList, err := createDictFile([]string{"Lion's", "Elephant's", "monkey"})
	defer os.Remove(wordList)
	if err != nil {

		t.Errorf("Couldn't create word list. Can't proceed with test : %v", err)
	}
	secretWord := getSecretWord(wordList)
	if secretWord != "monkey" {
		t.Errorf("Should get 'monkey' but Got %s", secretWord)
	}
}

func TestCorrectGuess(t *testing.T) {
	secretWord := "soldier"
	guess := 's'
	currentState := newHangman(secretWord)
	newState := checkGuess(currentState, byte(guess))
	expected := Hangman{
		secretWord:       currentState.secretWord,
		guessLetter:      append(currentState.guessLetter, byte(guess)),
		correctGuesses:   append(currentState.correctGuesses, byte(guess)),
		remainingChances: 7,
	}
	if newState.secretWord != expected.secretWord {
		t.Errorf("Secret word is modified")
	}
	if !bytes.Equal(newState.guessLetter, expected.guessLetter) {
		t.Errorf("Guess should be %q but got %q", expected.guessLetter, newState.guessLetter)
	}
	if !bytes.Equal(newState.correctGuesses, expected.correctGuesses) {
		t.Errorf("Correct Guess should be %q but got %q", expected.correctGuesses, newState.correctGuesses)
	}
	if !(newState.remainingChances == expected.remainingChances) {
		t.Errorf("Remaining chances is modified")
	}
}

func TestCorrectGuess2(t *testing.T) {
	secretWord := "soldier"
	guess := 'o'
	currentState := Hangman{
		secretWord:       secretWord,
		guessLetter:      []byte{'a', 'b', 's'},
		correctGuesses:   []byte{'s'},
		remainingChances: 5,
	}
	newState := checkGuess(currentState, byte(guess))

	expected := Hangman{
		secretWord:       currentState.secretWord,
		guessLetter:      append(currentState.guessLetter, byte(guess)),
		correctGuesses:   append(currentState.correctGuesses, byte(guess)),
		remainingChances: currentState.remainingChances,
	}

	if newState.secretWord != expected.secretWord {
		t.Errorf("Secret word is modified")
	}
	if !bytes.Equal(newState.guessLetter, expected.guessLetter) {
		t.Errorf("Guess should be %q but got %q", expected.guessLetter, newState.guessLetter)
	}

	if !bytes.Equal(newState.correctGuesses, expected.correctGuesses) {
		t.Errorf("Correct Guess should be %q but got %q", expected.correctGuesses, newState.correctGuesses)
	}

	if !(newState.remainingChances == expected.remainingChances) {
		t.Errorf("Remaining chances is modified")
	}
}
