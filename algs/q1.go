package algs

import (
	"fmt"
	"log"
)

// use the list structure
var q1 = ListQuestion{
	"Find the second greatest number",
	[]WantIntList{
		WantIntList{"t1", []int{2, 3, 1, 0, 9}, 3},
		WantIntList{"t2", []int{3, 2, 1, 4, 9}, 4},
	},
	nil,
}

func Find2ndGreatest(l []int) int {
	maxIdx := 0
	var max2, max int
	var maxs [2]int
	fmt.Println(l)
	if len(l) > 0 {
		max2, max = l[0], l[0]
		max2 = 0
		maxs[0], maxs[1] = max, max2
	} else {
		log.Fatal("list should len > 0")
	}
	for _, i := range l {
		log.Println(i)
		if i > max {
			max = i
			maxs[1-maxIdx] = maxs[maxIdx]
			maxs[maxIdx] = i
			//maxIdx = 1 - maxIdx // !!! error line
			log.Println(maxIdx, maxs)
		}
	}
	log.Println(max)
	// ! for the last one
	//return maxs[maxIdx]
	return maxs[1-maxIdx]
}
