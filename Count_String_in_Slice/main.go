package main

import "fmt"

func Mapping(slice []string) map[string]int {

	// Membuat map untuk menghitung jumlah masing-masing string
	hitungString := make(map[string]int)

	// Menghitung jumlah string dan menyimpan dalam map
	for _, kata := range slice {
		hitungString[kata]++
	}

	return hitungString
}

func main() {

	// test case
	fmt.Println(Mapping([]string{"asd", "qwe", "asd", "adi", "qwe", "qwe"}))

	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))

	fmt.Println(Mapping([]string{}))
}
