package main

import "fmt"

func main() {
	// Pendeklarasian Variabel 
	var name string
	name = "Evandri Ridho Hasmono"
	fmt.Println(name)

	// tanpa tipe data
	var nama = "Shafa Davina"
	fmt.Println(nama)

	// Tanpa tipe data dan var
	hobby := "Coding"
	fmt.Println(hobby)

	// Multiple variabel
	var (
		firstName = "Evandri"
		middleName = "Ridho"
		lastName = "Hasmono"
	)
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}