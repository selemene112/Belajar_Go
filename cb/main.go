package main

import (
	"fmt"
	"strings"
)

func main() {
    sentence := "selamat malam"

	kata := strings.Split(sentence, " ")

 
    
    // Memecah kalimat menjadi kata-kata
    words := make([]string, 0)
    word := ""
    for _, ch := range sentence {
        if ch == ' ' {
            words = append(words, word)
            word = ""
        } else {
            word += string(ch)
        }
    }
    words = append(words, word) // Menambahkan kata terakhir
    
    
    // Menghitung munculnya setiap karakter
    count := make(map[rune]int)
    for _, ch := range sentence {
        if ch != ' ' {
            count[ch]++
        }
    }

	for _, k := range kata {
		// Looping setiap karakter dalam setiap kata
		for _, h := range k {
			// Tampilkan setiap karakter dalam setiap kata pada baris baru
			fmt.Println(string(h))
		}
	}
    
    // Menampilkan jumlah munculnya setiap karakter
    fmt.Println("Jumlah karakter:")
    for ch, cnt := range count {
        fmt.Printf("%c : %d\n", ch, cnt)
    }
}