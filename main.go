package main

import (
	"encoding/json"
	"fmt"
)

// import "fmt"

type person struct {
	Firstname string
	Lastname  string
}

func main() {

	// hello := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// fmt.Println(sum(hello...))

	// fmt.Println(odd(sum, hello...))

	saeed := person{
		"saeed",
		"rasooli",
	}

	fmt.Println(saeed)

	saeedmarshall, err := json.Marshal(saeed)
	if err != nil {
		fmt.Printf("%V", err)
	}
	fmt.Println(string(saeedmarshall))

	newsaeed := person{
		"saeed",
		"rasooli",
	}
	json.Unmarshal(saeedmarshall, &newsaeed)

	fmt.Println(newsaeed)

}

func sum(h ...int) int {
	total := 0
	for _, v := range h {
		total += v
	}
	return total
}

func odd(f func(...int) int, h ...int) int {
	oddarray := []int{}
	for _, v := range h {
		if v%2 != 0 {
			oddarray = append(oddarray, v)
		}
	}

	return f(oddarray...)
}
