package saying

import (
	"fmt"
	"testing"
)

func GreetTest(t *testing.T) {
	v := Greet("saeed")
	if v != "heloo dear saeed" {
		t.Error("err")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("james"))
}
