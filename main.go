// first go file yay
package main

import (
	"fmt"
	t "time"
)

func main() {
	fmt.Println("hello world")
	name := "evan"

	num := 19

	fact := false

	//fmt.Printf("My name is %q", num)
	fmt.Printf("My name is %s and I am %d years old. Fact or cap? %t", name, num, fact)

	fmt.Println("The time right now is", t.Now())
	var number uint16 = 2
	fmt.Println(number)
}
