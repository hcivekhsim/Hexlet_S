package main

import (
	"fmt"
	fc"gmod/utils"
)



func main() {
	fmt.Println("Hello World")
	fmt.Println(fc.BuildGreeting("Alice", 5))
	d := fc.GetHiddenCard("2342345345435", 6)
	fmt.Println(d)
}
