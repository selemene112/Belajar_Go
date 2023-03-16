package main

import (
	"fmt"
	"strings"
)

func main() {
	
	// Variabel string kalimat
	kalimat := "selamat malam"

	// Pecah kalimat menjadi kata-kata yang terpisah
	kata := strings.Split(kalimat, " ")

	//============================================ penjumlahan kata=======================


	//=============================================end penjumlahan kata=================

	// Looping setiap kata pada variabel kata
	for _, k := range kata {
		// Looping setiap karakter dalam setiap kata
		for _, h := range k {
			// Tampilkan setiap karakter dalam setiap kata pada baris baru
			fmt.Println(string(h))
		}
	}
}