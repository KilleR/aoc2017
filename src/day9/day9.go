package main

import (
	"getInput"
	"fmt"
)

func readStream(s string) {
	inBytes := []byte(s)

	var (
		inGarbage, skipNext bool
		nestLevel, totalScore, totalGarbage int
		scores []int
	)

	totalGarbage = 0
	for _,v := range inBytes {
		if inGarbage {
			if skipNext {
				skipNext = false
				continue
			}
			switch v {
			case '>':
				inGarbage = false
			case '!':
				skipNext = true
			default:
				totalGarbage++
			}
			continue
		}

		switch v {
		case '{':
			nestLevel++
			scores = append(scores, nestLevel)
			totalScore += nestLevel
		case '}':
			nestLevel--
		case '<':
			inGarbage = true
		}
	}

	fmt.Println("Scores:",scores)
	fmt.Println("Total:",totalScore)
	fmt.Println("Total garbage:",totalGarbage)
}

func main() {
	input := getInput.GetInputString("src/day9/input")
	//input = "{{<a!>},{<a!>},{<a!>},{<ab>}}"

	fmt.Println(input)
	readStream(input)
}
