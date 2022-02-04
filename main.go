package main

import (
	"fmt"
	"sort"
)

type people []string

func (x people) Len() int { return len(x) }

func (x people) Less(i, j int) bool { return x[i] < x[j] }

func (x people) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func main() {
	// saeed2 := []string{"ahmad", "rajab", "sasee", "sfaga", "sdasdkahwuqwqwfq", "Qqqqq", "afakafa", "wfksdf"}
	saeed3 := []int{3, 5, 3, 2, 4, 5, 7, 5, 4, 1, 9}
	var gam interface{}
	gam = saeed3

	gam = 32

	gam = "hamaid"
	fmt.Println(gam)
	// sort.StringSlice.Sort(saeed2)
	// fmt.Println(saeed2)

	sort.Sort(sort.Reverse(sort.IntSlice(saeed3)))
	// sort.Ints(saeed3)

	fmt.Println(saeed3)
}
