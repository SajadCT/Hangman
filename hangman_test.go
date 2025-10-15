package main

import (
	"strings"
	"testing"
)

func TestSecretWordNoCapitals(t *testing.T) {
	wordList := "/usr/share/dict/words"
	secretWord := getSecretWord(wordList)
	if secretWord != strings.ToLower(secretWord) {
		t.Errorf("Should not get words with capital letters. Got %s", secretWord)
	}

}

func TestSecretWordLength(t *testing.T) {
	wordList := "/usr/share/dict/words"
	secretWord := getSecretWord(wordList)
	if len(secretWord) < 6 {
		t.Errorf("Should have minimum 6 characters. Got %d", len(secretWord))
	}
}

func TestSecretWordPunctuations(t *testing.T) {
	wordList := "/usr/share/dict/words"
	secretWord := getSecretWord(wordList)
	for _, ch := range secretWord {
		if !(ch >= 'a' && ch <= 'z') {
			t.Errorf("Should not have punctuation in the word. Got %s", secretWord)
		}
	}
}
