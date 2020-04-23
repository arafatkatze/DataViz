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
	var max2, max int
	fmt.Println(l)
	if len(l) > 0 {
		max2, max = l[0], l[0]
	} else {
		log.Fatal("list should len > 0")
	}
	for _, i := range l {
		log.Println(i)
		if i > max {
			max = i
		} else if i > max2 {
			fmt.Println(i, ">", max2)
			max2 = i
		}
	}
	return max2
}
