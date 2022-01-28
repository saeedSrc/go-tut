package main

import "fmt"

func main() {
	words := make(map[string]string)

	words["saeeed"] = ""
	words["hamid"] = ""
	words["samajideeed"] = ""
	words["haw"] = ""
	words["asdads"] = ""
	words["saasdasdeeed"] = ""
	words["ljasld"] = ""
	words["wquey"] = ""
	words["kweqg"] = ""
	words["sqqweeeed"] = ""
	words["wqqq"] = ""
	words["qqqq"] = ""

	repeats := make(map[byte]int)
	for k, _ := range words {

		repeats[k[0]]]++
		// if _, ok := repeats[k[0]]; !ok {
		// }
	}

	for k, v := range repeats {
		fmt.Println(string(k), v)
	}
}
