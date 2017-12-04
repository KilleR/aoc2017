package main

import (
	"math"
	"fmt"
)

func makeSpiralMemory(in int, spiralSum bool) [][]int {
	size := int(math.Ceil(math.Sqrt(float64(in))))+2
	if size%2 == 0 {
		size++
	}
	memS := make([][]int, size)
	for i := range memS {
		memS[i] = make([]int, size)
	}

	centre := int(math.Ceil(float64(size/2)))
	fmt.Println("centre:",centre)
	fmt.Println("size:",size)
	locX := 0
	locY := 0
	direction := 0
	//ring := 1
	spiralLoop:
	for i := 1; i<=in; i++ {
		//fmt.Println(i, locX, locY, centre+locX, centre+locY)
		if spiralSum {
			localSum := 0
			if i == 1 {
				localSum = 1
			} else {
				for x := centre + locX - 1; x <= centre+locX+1; x++ {
					for y := centre + locY - 1; y <= centre+locY+1; y++ {
						//fmt.Println("Getting",x,y)
						localSum += memS[y][x]
					}
				}
			}
			memS[centre+locY][centre+locX] = localSum
			if in < localSum {
				fmt.Println("Val:", localSum)
				break spiralLoop
			}
		} else {
			memS[centre+locY][centre+locX] = i
		}

		if i == in {
			break
		}
		switch direction {
		case 0:
			locX++
			direction = 1
		case 1:
			locY--
			if locX == -locY {
				direction = 2
			}
		case 2:
			locX--
			if -locX == -locY {
				direction = 3
			}
		case 3:
			locY++
			if -locX == locY {
				direction = 4
			}
		case 4:
			locX++
			if locX == locY {
				direction = 0
			}
		}
		//fmt.Println("GRID")
		//for _,r := range memS {
		//	fmt.Println(r)
		//}
	}

	//for _,r := range memS {
	//	fmt.Println(r)
	//}

	fmt.Println(locX, locY)


	return memS
}

func main() {
	_= makeSpiralMemory(312051, true)

}
