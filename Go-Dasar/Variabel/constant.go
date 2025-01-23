package main

import "fmt"

func main() {
	// assign 1 variabel
	const name = "Ridho"
	fmt.Println(name)

	// assign multi variabel
	const (
		firstName = "Evandri"
		lastName = "Hasmono"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}