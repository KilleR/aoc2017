package main

import (
	"io/ioutil"
	"log"
	"strings"
	"regexp"
	"fmt"
	"strconv"
)

var depth int

type holder struct {
	Name    string
	Weight  int
	TowerWeight int
	Held    []*holder
	Carried bool
}

func getInput() []string {
	file, err := ioutil.ReadFile("src/day7/input")
	if err != nil {
		log.Fatalln("Failed to open input:", err)
	}

	return strings.Split(string(file), "\r\n")
}

func parseInput() {
	holders := make(map[string]*holder)

	input := getInput()

	rex := regexp.MustCompile(`([a-z]+) \(([0-9]+)\)( -> )?(([a-z]+,? ?)+)?`)

	for _,v := range input {
		res := rex.FindStringSubmatch(v)
		//fmt.Println(v)
		//fmt.Println(res)
		//fmt.Println("Name:",res[1])
		//fmt.Println("Weight:",res[2])
		//fmt.Println("Held:",res[4])
		w, err := strconv.Atoi(res[2])
		if err != nil {
			log.Fatalln("Failed to convert weight:",err)
		}
		// if there's no holder at this location, make one
		thisHolder,ok := holders[res[1]]
		if !ok {
			thisHolder = &holder{}
		}
		thisHolder.Name = res[1]
		thisHolder.Weight = w


		if res[4] != "" {
			held := strings.Split(res[4], ", ")
			for _,h := range held {
				// if there's no holder for this held holder, make one
				heldHolder,ok := holders[h]
				if !ok {
					heldHolder = &holder{}
				}
				heldHolder.Carried = true
				thisHolder.Held = append(thisHolder.Held, heldHolder)
				holders[h] = heldHolder
			}
		}

		holders[res[1]] = thisHolder
	}

	// second pass : find those which are not carried, and work out sum weights
	for _,h := range holders {
		if !h.Carried {
			sumWeight(h)
			fmt.Println("Not carried!:",h)
		}
	}

	//for _,h := range holders {
	//	fmt.Println(h)
	//}

	fmt.Println("Have ",len(holders),"holders")
}

func sumWeight(hol *holder) int {
	depth++
	fmt.Print(strings.Repeat("> ", depth))
	fmt.Println("Get",hol.Name,"weight")
	totalWeight := hol.Weight;
	var weights []int
	for _,h := range hol.Held {
		w := sumWeight(h)
		weights = append(weights, w)
		totalWeight += w
	}
	hol.TowerWeight = totalWeight

	// check for matching weights
	for i,h := range hol.Held {
		for _,h2 := range hol.Held[i+1:] {
			if h.TowerWeight!=h2.TowerWeight {
				fmt.Println("Weight mismatch on",hol,"between:",h.Name,h.TowerWeight,"and",h2.Name,h2.TowerWeight)
			}
		}
	}

	depth--
	return totalWeight
}

func main() {
	parseInput()
}
