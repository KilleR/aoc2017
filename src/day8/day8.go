package main

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"regexp"
	"strconv"
)

var registers map[string]int
var cmdRex *regexp.Regexp

func getInput(fileName string) []string {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Failed to open input:", err)
	}

	return strings.Split(string(file), "\r\n")
}

func doCmd(cmd string) {
	//fmt.Println("Processing",cmd)

	res := cmdRex.FindStringSubmatch(cmd)

	incrementor, err := strconv.Atoi(res[3])
	if err != nil {
		log.Fatalln("Error getting incrementor", err)
	}
	//fmt.Println("Incrementor:", incrementor)

	comparator, err := strconv.Atoi(res[6])
	if err != nil {
		log.Fatalln("Error getting comparator", err)
	}
	//fmt.Println("Comparator:", comparator)

	sourceRegister := res[4]
	//fmt.Println("Source Register:", sourceRegister)

	destRegister := res[1]
	//fmt.Println("Destination register:", destRegister)

	incDec := res[2]
	//fmt.Println("Increment:", incDec)

	op := res[5]
	//fmt.Println("Operator", op)

	// check both the destination and source registers exist
	source, ok := registers[sourceRegister]
	if !ok {
		//fmt.Println("Source register",sourceRegister,"does not exist, creating...")
		source = 0
		registers[sourceRegister] = source
	}

	dest, ok := registers[destRegister]
	if !ok {
		//fmt.Println("Destination register",destRegister,"does not exist, creating...")
		dest = 0
		registers[destRegister] = dest
	}

	truthiness := false
	// do compare
	switch op {
	case "<":
		if source < comparator {
			truthiness = true
		}
	case "<=":
		if source <= comparator {
			truthiness = true
		}
	case ">":
		if source > comparator {
			truthiness = true
		}
	case ">=":
		if source >= comparator {
			truthiness = true
		}
	case "==":
		if source == comparator {
			truthiness = true
		}
	case "!=":
		if source != comparator {
			truthiness = true
		}
	}
	//fmt.Println("Truthiness?",truthiness)

	if truthiness {
		switch incDec {
		case "inc":
			dest += incrementor
		case "dec":
			dest -= incrementor
		}
	}

	registers[destRegister] = dest
}

func getMaxReg() int {
	maxReg := 0
	for _,v := range registers {
		if v> maxReg {
			maxReg = v
		}
	}
	return maxReg
}

func init() {
	registers = make(map[string]int)
	cmdRex = regexp.MustCompile(`([a-z]+) (inc|dec) (-?[0-9]+) if ([a-z]+) (<|>|<=|>=|==|!=) (-?[0-9]+)`)
}

func main() {
	input := getInput("src/day8/input")

	maxRegDuring := 0
	for _,c := range input {
		doCmd(c)
		max := getMaxReg()
		if max > maxRegDuring {
			maxRegDuring = max
		}
	}

	// get largest register
	maxReg := getMaxReg()

	fmt.Println("Largest register:",maxReg)
	fmt.Println("Largest register at any time:",maxRegDuring)
}
