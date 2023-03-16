package main

import (
	"fmt"
	"os"
)

//=============================================start init type data===========================
type Teman struct {
	Nama     string
	Alamat   string
	Pekerjaan string
	Alasan   string
}
//==============================================END init type data==========================
//================================================start init data============================
var temanKelas = []Teman{
	{"Muhammad", "Jl.Gunung ", "Developer", "Mau jadi Developer"},
	{"Ariyanda", "Jl. Lembah", "Designer", "Mau jadi Designer"},
	{"Zulyadiansyah", "JL. Tebing", "Data Scientist", "Mau jadi Data Scientist"},
}
//================================================END init data ===========================

func main() {
	arg := os.Args[1]
	idx := argToInt(arg)
//===================================================PESAN ERROR============================
	if idx < 1 || idx > len(temanKelas) {
		fmt.Println("Nomor absen tidak valid")
		return
	}
//=================================================END PESAN ERROR===========================
	teman := temanKelas[idx-1]
	fmt.Printf("Data teman dengan nomor absen %d:\n", idx)
	fmt.Printf("Nama     : %s\n", teman.Nama)
	fmt.Printf("Alamat   : %s\n", teman.Alamat)
	fmt.Printf("Pekerjaan : %s\n", teman.Pekerjaan)
	fmt.Printf("Alasan   : %s\n", teman.Alasan)
}
//========================================== STARAT ERROR DATA=================================
func argToInt(arg string) int {
	var idx int
	_, err := fmt.Sscanf(arg, "%d", &idx)
	if err != nil {
		return 0
	}
	return idx
}

//=========================================== END ERROR DATA==================================