package main

import "fmt"

func main() {
	var ujian = 80
	var absen = 80

	var lulusUjian = ujian >= 80
	var luluAbsen = absen >= 80

	var lulus = lulusUjian && luluAbsen

	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absen >= 80)
}
