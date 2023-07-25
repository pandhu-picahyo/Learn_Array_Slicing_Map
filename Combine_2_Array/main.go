package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {

	// menggunakan map untuk menyimpan data
	dataUnik := make(map[string]bool)

	// mengisi map dengan data dari Array pertama
	for _, nilai := range arrayA {
		dataUnik[nilai] = true
	}

	// mengisi map dengan data dari Array kedua jika data belum masuk
	for _, nilai := range arrayB {
		if _, ada := dataUnik[nilai]; !ada {
			dataUnik[nilai] = true
		}
	}

	// membuat slice baru untuk menyimpan data unik dari map
	sliceUnik := make([]string, 0, len(dataUnik))
	for kunci := range dataUnik {
		sliceUnik = append(sliceUnik, kunci)
	}

	return sliceUnik
}

func main() {

	// Test cases

	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))

	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))

	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))

	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))

	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))

	fmt.Println(ArrayMerge([]string{}, []string{}))

}
