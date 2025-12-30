package main

import (
	"fmt"
	"math"
)

func luasPersegiPanjang(panjang, lebar float64) float64 {
	return panjang * lebar
}
func luasKelilingPersegiPanjang(panjang, lebar float64) float64 {
	return (panjang + lebar) * 2
}
func kelilingSegitiga(panjang float64) float64 {
	return panjang * 3
}

func luasSegitiga(alas, tinggi float64) float64 {
	return (alas * tinggi) / 2
}

func luasLingkaran(r float64) float64 {
	return math.Pi * r * r //r = jari jari
}

func kelilinLingkaran(r float64) float64 {
	return 2 * math.Pi * r
}

func main() {
	var jenis int

	fmt.Println("Pilih operasi:")
	fmt.Println("1. Hitung luas persegi")
	fmt.Println("2. Hitung keliling persegi")
	fmt.Println("3. Hitung keliling segitiga")
	fmt.Println("4. Hitung luas segitiga")
	fmt.Println("5. Hitung luas lingkaran")
	fmt.Println("6. Hitung keliling lingkaran")
	fmt.Print("Masukkan pilihan angka: ")
	fmt.Scan(&jenis)

	switch jenis {
	case 1:
		var panjang, lebar float64
		fmt.Print("\nMasukkan panjang: ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar: ")
		fmt.Scan(&lebar)
		fmt.Println("Luas persegi panjang =", luasPersegiPanjang(panjang, lebar), "cm")

	case 2:
		var panjang, lebar float64
		fmt.Println("\nMasukkan Panjang: ")
		fmt.Scan(&panjang)
		fmt.Println("Masukkan lebar: ")
		fmt.Scan(&lebar)
		fmt.Println("Keliling persegi panjang =", luasKelilingPersegiPanjang(panjang, lebar), "cm")

	case 3:
		var panjang float64
		fmt.Println("\nMasukkan panjang: ")
		fmt.Scan(&panjang)
		fmt.Println("Keliling segitiga", kelilingSegitiga(panjang), "cm")

	case 4:
		var alas, tinggi float64
		fmt.Println("\nMasukkan alas: ")
		fmt.Scan(&alas)
		fmt.Println("Masukkan tinggi: ")
		fmt.Scan(&tinggi)
		fmt.Println("Luas segitiga", luasSegitiga(alas, tinggi), "cm")

	case 5:
		var r float64
		fmt.Println("\nMasukkan jari-jari lingkaran: ")
		fmt.Scan(&r)
		fmt.Println("Luas lingkaran", luasLingkaran(r), "cm")

	case 6:
		var r float64
		fmt.Println("\nMasukkan jari-jari lingkaran: ")
		fmt.Scan(&r)
		fmt.Println("Keliling lingaran", kelilinLingkaran(r), "cm")

	default:
		fmt.Println("\nPilihan tidak valid!")

	}

}
