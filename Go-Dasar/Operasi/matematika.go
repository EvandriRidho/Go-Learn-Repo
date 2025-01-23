package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := a + b
	fmt.Println(c)


	i := 10
	i += 10 // i = i + 10
	fmt.Println(i)


	j := 1
	j++ // j = j + 1
	j++ // j = j + 1
	fmt.Println(j)
}