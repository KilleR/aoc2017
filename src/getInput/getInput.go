package getInput

import (
	"io/ioutil"
	"log"
	"strings"
)

func GetInputLines(fileName string) []string {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Failed to open input:", err)
	}

	return strings.Split(string(file), "\r\n")
}

func GetInputString(fileName string) string {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Failed to open input:", err)
	}

	return string(file)
}