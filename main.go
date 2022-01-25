
package main

// import some
import (
	"fmt"
)

var costumizeErr = fmt.Errorf("this is formatted err : %v", "format")

type devide struct {
	first int
	last  int
	text  string
}


func main() {
  saeed := devide {
	  1,2, "salam",
  }

  saeed.Devider(1, 2)
}

//Devider This is devider package
func (d devide) Devider(f, l int) (result float64,  err error) {
	if l == 0 {
		return 0, costumizeErr
	}

	return 1, nil
}

func (d devide) Error() string {
	return d.text
}
