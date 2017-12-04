package main

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"sort"
)

func getInput() []string {
	file, err := ioutil.ReadFile("src/day4/input")
	if err != nil {
		log.Fatalln("Failed to open input:", err)
	}

	return strings.Split(string(file), "\r\n")
}

func checkValidPassphrase(password string) bool {

	words := strings.Split(password, " ")

	for i,word := range words {
		for _, word2 := range words[i+1:] {
			if sortString(word) == sortString(word2) {
				return false
			}
		}
	}

	return true
}

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main() {
	inSlice := getInput()
	validPassphrases := 0
	for _, r := range inSlice {
		isValid := checkValidPassphrase(r)
		fmt.Println(r, isValid)
		if isValid {
			validPassphrases++
		}
	}

	fmt.Println(validPassphrases)
}
