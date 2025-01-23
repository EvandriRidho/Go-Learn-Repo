package main

import "fmt"

func main() {
	name1 := "Evandri"
	name2 := "Davina"

	result1 := name1 == name2 // samadengan
	result2 := name1 != name2 // tidak samadengan

	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Println("=====================")

	a := 10
	b := 20
	fmt.Println(a < b) // kurang dari
	fmt.Println(a > b) // lebih dari
	fmt.Println(a >= b) // lebih dari sama dengan
	fmt.Println(a <= b) // kurang dari sama dengan
}