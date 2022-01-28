package main

import "fmt"

func main() {
	words := make(map[string]string)

	words["saeeed"] = ""

	repeats := make(map[byte]int)
	for k, _ := range words {

		repeats[k[0]]++
		fmt.Println(hashWords(k, 12))
		// if _, ok := repeats[k[0]]; !ok {
		// }
	}

}

func hashWords(word string, hash int) int {

	sum := 0
	for _, v := range word {

		sum += int(v)
	}

	return sum % 12
}
