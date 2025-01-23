package main

import "fmt"

func main() {
	nilaiUjian := 90
	absensi := 80

	lulusUjian := nilaiUjian > 80
	lulusAbsensi := absensi > 80

	lulus := lulusUjian && lulusAbsensi // AND operator
	fmt.Println(lulus)
}