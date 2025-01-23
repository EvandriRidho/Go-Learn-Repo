package main

import "fmt"

func main() {
	type noKTP string 
	var ktp noKTP = "123456"
	fmt.Println(ktp)
	fmt.Println(noKTP("654321"))
}