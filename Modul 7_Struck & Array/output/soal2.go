package main

import (
	"fmt"
	"math"
)

// Mengisi array sesuai dengan jumlah yang dimasukkan
func isiArray(jumlah int) []int {
	array := make([]int, jumlah)
	for i := 0; i < jumlah; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&array[i])
	}
	return array
}

// Mengelompokkan array menjadi genap dan ganjil berdasarkan index bukan nilai
func ganjilGenap(array []int) ([]int, []int) {
	genap, ganjil := []int{}, []int{}
	for i := 0; i < len(array); i++ {
		if i%2 == 0 {
			genap = append(genap, array[i])
		} else {
			ganjil = append(ganjil, array[i])
		}
	}
	return genap, ganjil
}

// Menampilkan elemen dengan indeks kelipatan
func kelipatan(array []int, bilangan int) []int {
	hasil := []int{}
	for i := bilangan; i < len(array); i += bilangan {
		hasil = append(hasil, array[i])
	}
	return hasil
}

// Menghapus elemen berdasarkan indeks
func hapusIndeks(array []int, indeks int) []int {
	if indeks < 0 || indeks >= len(array) {
		fmt.Println("Indeks tidak ditemukan")
		return array
	}
	return append(array[:indeks], array[indeks+1:]...)
}

// Menghitung rata-rata
func rataRata(array []int) float64 {
	total := 0
	for i := 0; i < len(array); i++ {
		total += array[i]
	}
	return float64(total) / float64(len(array))
}

// Menghitung simpangan baku
func simpanganBaku(array []int) float64 {
	rata := rataRata(array)
	var jumlah float64
	for i := 0; i < len(array); i++ {
		jumlah += math.Pow(float64(array[i])-rata, 2)
	}
	return math.Sqrt(jumlah / float64(len(array)))
}

// Menghitung frekuensi
func frekuensi(array []int, angka int) int {
	jumlah := 0
	for i := 0; i < len(array); i++ {
		if array[i] == angka {
			jumlah++
		}
	}
	return jumlah
}

func main() {
	var jumlah, bilangan, indeks, angka int

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&jumlah)

	if jumlah <= 0 {
		fmt.Println("Program berhenti karena jumlah elemen adalah <=0.")
		return
	}

	// Menampilkan isi, rata-rata, dan simpangan baku
	array := isiArray(jumlah)
	fmt.Println("\nIsi array yang dimasukkan:", array)
	fmt.Printf("Rata-rata: %.2f\n", rataRata(array))
	fmt.Printf("Simpangan baku: %.2f\n", simpanganBaku(array))

	// Menampilkan indeks genap dan ganjil
	genap, ganjil := ganjilGenap(array)
	fmt.Println("\nElemen dengan indeks genap:", genap)
	fmt.Println("Elemen dengan indeks ganjil:", ganjil)

	// Menampilkan elemen dengan indeks kelipatan
	fmt.Print("\nMasukkan nilai kelipatan x: ")
	fmt.Scan(&bilangan)
	hasilKelipatan := kelipatan(array, bilangan)
	fmt.Println("Elemen dengan indeks kelipatan", bilangan, ":", hasilKelipatan)

	// Menghapus elemen
	fmt.Print("\nMasukkan indeks yang ingin dihapus: ")
	fmt.Scan(&indeks)
	array = hapusIndeks(array, indeks)
	fmt.Println("Array setelah penghapusan:", array)

	// Menampilkan rata-rata dan simpangan setelah penghapusan
	fmt.Printf("\nRata-rata setelah penghapusan: %.2f\n", rataRata(array))
	fmt.Printf("Simpangan setelah penghapusan: %.2f\n", simpanganBaku(array))

	// Mencari frekuensi dari bilangan yang ada di array
	fmt.Print("\nMasukkan bilangan untuk mencari frekuensi: ")
	fmt.Scan(&angka)
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", angka, frekuensi(array, angka))
}