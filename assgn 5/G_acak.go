package main

import (
	"fmt"
	"time"
)

//====================================================Start Groutine 1 =============================
func process1(data interface{}) {
	for i := 1; i < 5; i++ {
		fmt.Println(data,i)
		
	}
}
//=====================================================End Groutine 1======================
//=====================================================Start Groutine 2===================
func process2(data interface{}) {
	
	for i := 1; i < 5; i++ {
		fmt.Println(data,i)
		
	}
}
//====================================================End Groutine 2 ========================

func main() {
	data1 := "[bisa1 bisa2 bisa3]" // data groutine 
	data2 := "[coba1 coba2 coba3]" // data groutine

	for i := 0; i < 1; i++ {
		go process1(data1)
		go process2(data2)
	}

	time.Sleep(5 * time.Second) // jeda waktu agar groutine kelihatan
}