package main

import "fmt"

func main()  {
	/*
		komentar kode
		menampilkan pesan hello world
	*/

	fmt.Println("=================")
	fmt.Printf("Hello World!\n") // karena tidak otomatis line baru maka ditambah \n
	fmt.Println("testing") // tidak perlu \n karena sudah otomatis line baru

	// fmt.Println("baris ini tidak akan dieksekusi")

	/* ======================== */
	// Belajar Variabel
	var firstName string = "Dani" // Deklarasi variabel dengan menuliskan var dan tipe data atau disebut manifest typing

	lastName := "Munif" // Deklarasi variabel tanpa tipe data, otomatis mengikuti tipe data nilainya. harus pakai titik dua samadengan (atau disebut Type Inference)

	var usia, tinggiBadan, beratBadan string

	usia, tinggiBadan, beratBadan = "25", "159", "60" // Deklarasi multi variabel bisa menggunakan koma, multi variabel bisa juga dengan cara Type Inference

	_ = "belajar Golang"
	_ = "Golang itu mudah"
	// Deklarasi "reserved variable" yang bisa dimanfaatkan untuk menampung nilai tidak terpakai / variabel keranjang sampah

	// contohNew := new(string)
	// Deklarasi variabel dengan keyword new, menampung data tipe pointer string, nnti ada penjelasan ttg pointer

	/* Deklarasi variabel dengan keyword make
	   hanya untuk bbrp jenis variabel yaitu channel, slice, map
	   nnti ada penjelasannya
	*/

	fmt.Println("=================")
	fmt.Printf("Halo %s %s!\nKamu sekarang berusia %s\nDengan tinggi badan %s\nDan berat badan %s", firstName, lastName, usia, tinggiBadan, beratBadan)

	/* =================== */
	/* Belajar Tipe Data */



	fmt.Println("=================")
}
