package main

import (
	"fmt"

	"github.com/thejaswa/puppy"
)

func main() {

	s1 := puppy.Bark()
	s2 := puppy.Braks()
	s5 := puppy.BigBark()
	s6 := puppy.BigBarks()

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s5)
	fmt.Println(s6)

	fmt.Println(puppy.Bark())
	fmt.Println(puppy.Braks())
	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.BigBarks())

}
