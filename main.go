package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)

	convert(&a)
	fmt.Println(a)

}

func convert(b *int) {
	*b = 10
}
