package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {

	// membuat map untuk menghitung jumlah masing-masing angka
	hitungAngka := make(map[int]int)

	// menghitung jumlah angka dan menyimpannya dalam map
	for _, isiData := range angka {
		jumAngka, _ := strconv.Atoi(string(isiData))
		hitungAngka[jumAngka]++
	}

	// membuat slice untuk menyimpan angka yang hanya muncul 1 kali
	var nilaiUnik []int
	for nilai, jumlah := range hitungAngka {
		if jumlah == 1 {
			nilaiUnik = append(nilaiUnik, nilai)
		}
	}

	return nilaiUnik

}

func main() {

	// test case
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("1122"))
}
