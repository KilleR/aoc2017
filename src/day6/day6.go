package main

import "fmt"

func redistribute(inSlice []int) []int {
	// find the largest int
	max := 0;
	largestIndex := 0
	for i,v := range inSlice {
		if v > max {
			max = v
			largestIndex = i
		}
	}

	toRedistribute := inSlice[largestIndex]
	placementIndex := largestIndex
	inSlice[largestIndex] = 0
	for i:= 0; i < toRedistribute; i++ {
		placementIndex++
		// wrap around on pointer overflow
		if placementIndex >= len(inSlice) {
			placementIndex = 0
		}
		inSlice[placementIndex]++
	}

	return inSlice
}

func checkForDuplicates(o [][]int, n []int) *int {
	var isEqual *int

	eL:
	for x,r := range o {
		// if they are not the same size, they are automatically not equal
		if len(r) != len(n) {
			continue
		}
		for i,v := range r {
			// if they are different, then this set is not a match, move on to the next
			if v != n[i] {
				continue eL
			}
		}
		// only if all are equal do we have a a match#
		isEqual = &x
		fmt.Println("Equality at",x)
		break eL
	}

	return isEqual
}

func main() {
	input := []int{4,10,4,1,8,4,9,14,5,1,14,15,0,15,3,5}
	//input := []int{0, 2, 7, 0}

	var (
		pastValues [][]int
		dupIndex *int
	)
	i := 0;
	// reslice
	n := []int{}
	n = append(n, input...)
	pastValues = append(pastValues, n)

	for {
		i++
		input = redistribute(input)
		dupIndex = checkForDuplicates(pastValues, input)
		if dupIndex != nil {
			fmt.Println("Duplicate",input,"in",pastValues)
			fmt.Println("Duplicate index:", *dupIndex)
			break
		}
		// reslice
		n := []int{}
		n = append(n, input...)
		pastValues = append(pastValues, n)
	}
	fmt.Println("Loop after:",i)
	fmt.Println("Loop length:",i-*dupIndex)

	// DO IT AGAIN!
	//i = 0;
	//pastValues = [][]int{}
	//dupIndex = 0
	//for {
	//	i++
	//	input = redistribute(input)
	//	dupIndex = checkForDuplicates(pastValues, input)
	//	if dupIndex != 0 {
	//		fmt.Println("Duplicate index:", dupIndex)
	//		break
	//	}
	//	// reslice
	//	n := []int{}
	//	n = append(n, input...)
	//	pastValues = append(pastValues, n)
	//}
	//fmt.Println("Loop after:",i)
}
