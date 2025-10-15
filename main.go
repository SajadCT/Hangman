package main

import "fmt"

func getSecretWord(wordFileName string) string {
	return "elephant"
}

func main() {
	fmt.Println(getSecretWord("/usr/share/dict/words"))

}
