package main

import (
	"fmt"
	"container/ring"
	"strings"
	"strconv"
)

var (
	skipSize, pos int
	loop          *ring.Ring
)

func makeLoop(len int) {
	loop = ring.New(len)
	for i := 0; i < len; i++ {
		loop.Value = i
		loop = loop.Next()
	}
}

func printRing(r *ring.Ring) {
	r.Do(func(x interface{}) {
		fmt.Print(x," ")
	})
	fmt.Println()
}

func reverseInts(input []int) []int {
	if len(input) == 0 {
		return input
	}
	return append(reverseInts(input[1:]), input[0])
}

func knot(length int) {
	var firstHalf *ring.Ring

	loop = loop.Prev()
	firstHalf = loop.Unlink(length)

	fmt.Print("First:")
	printRing(firstHalf)

	revRing := ring.New(length)
	firstHalf.Do(func(x interface{}) {
		revRing = revRing.Move(-1)
		revRing.Value = x
	})

	fmt.Print("First reversed:")
	printRing(revRing)

	fmt.Print("Loop after split:")
	printRing(loop)

	if(length == loop.Len()) {
		loop = revRing
	} else {
		loop = loop.Link(revRing)
	}
	fmt.Print("Loop after merge:")
	printRing(loop)

	loop = loop.Move(skipSize)

	pos += length+skipSize
	skipSize++
}

func main() {
	makeLoop(256)
	input := "189,1,111,246,254,2,0,120,215,93,255,50,84,15,94,62"

	for _,v := range strings.Split(input, ",") {
		i,_ := strconv.Atoi(v)
		knot(i)

		fmt.Print("Loop from start: ")
		printRing(loop.Move(-pos))
	}

	fmt.Println("Multiple of",loop.Move(-pos).Value,"+",loop.Move(-pos+1).Value,"=",(loop.Move(-pos).Value.(int) * loop.Move(-pos+1).Value.(int)))
}
