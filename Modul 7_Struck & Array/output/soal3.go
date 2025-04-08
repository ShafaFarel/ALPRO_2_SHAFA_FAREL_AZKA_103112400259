package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		// Jika skor negatif maka program akan berakhir
		if skorA < 0 || skorB < 0 {
			break
		}

		// Menentukan pemenang dan menyimpannya dalam array
		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorA < skorB {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}

		pertandingan++
	}

	// Menampilkan semua hasil setelah input selesai
	for i, pemenang := range hasil {
		fmt.Printf("Hasil %d: %s\n", i+1, pemenang)
	}

	fmt.Println("Pertandingan selesai")
}