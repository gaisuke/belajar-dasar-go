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

	var usia, tinggiBadan, beratBadan uint8

	usia, tinggiBadan, beratBadan = 25, 159, 60 // Deklarasi multi variabel bisa menggunakan koma, multi variabel bisa juga dengan cara Type Inference

	_ = "belajar Golang"
	_ = "Golang itu mudah"
	// Deklarasi "reserved variable" yang bisa dimanfaatkan untuk menampung nilai tidak terpakai / variabel keranjang sampah

	// contohNew := new(string)
	// Deklarasi variabel dengan keyword new, menampung data tipe pointer string, nnti ada penjelasan ttg pointer

	/* Deklarasi variabel dengan keyword make
	   hanya untuk bbrp jenis variabel yaitu channel, slice, map
	   nnti ada penjelasannya
	*/

	fmt.Println("\n=================")
	fmt.Printf("Halo %s %s!\nKamu sekarang berusia %d\nDengan tinggi badan %d\nDan berat badan %d\n", firstName, lastName, usia, tinggiBadan, beratBadan) // %s untuk memanggil variable dengan tipe data string yg berurutan, sedangkan %d untuk tipe data numerik non-desimal

	/* =================== */
	/* Belajar Tipe Data */

	/*
		Tipe data	Cakupan bilangan
		uint8		0 ↔ 255
		uint16		0 ↔ 65535
		uint32		0 ↔ 4294967295
		uint64		0 ↔ 18446744073709551615
		uint		sama dengan uint32 atau uint64 (tergantung nilai)
		byte		sama dengan uint8
		int8		-128 ↔ 127
		int16		-32768 ↔ 32767
		int32		-2147483648 ↔ 2147483647
		int64		-9223372036854775808 ↔ 9223372036854775807
		int			sama dengan int32 atau int64 (tergantung nilai)
		rune		sama dengan int32
	*/

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	var decimalNumber = 2.62

	var exist bool = true

	var message string = "Halo gaisss!"
	var backtickMessage = `Nama saya "Dani Munif".
	Yoroshiku Onegaishimasu.
	Mari belajar "Golang".`

	fmt.Println("\n=================")
	fmt.Printf("bilangan positif: %d\n", positiveNumber) // %d digunakan untuk memanggil variable dengan tipe data numerik non-desimal
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	fmt.Printf("bilangan desimal: %f\n", decimalNumber) // %f digunakan untuk memanggil variable dengan tipe data numerik desimal
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber) // %f digunakan untuk memanggil variable dengan tipe data numerik desimal dengan tambahan format 3 angka dibelakang koma

	fmt.Printf("exist? %t\n", exist) // %t digunakan untuk memanggil variable dengan tipe data bool

	fmt.Printf("message: %s\n", message) // %s digunakan untuk memanggil variable dengan tipe data string
	fmt.Printf("backtick message: %s\n", backtickMessage) // %s digunakan untuk memanggil variable dengan tipe data string
}
