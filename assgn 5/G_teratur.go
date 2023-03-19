package main

import (
	"fmt"
	"sync"
	"time"
)

//Catatan ariyanda
//program untuk menampilkan data 1,2 secara berurutan di groutine belum bisa
//19 03 2023 ==> waktu pengerjaan
// tindakan :
// sudah pakai Mutex

//==============================================Groutine 1 ===========================
func process1(data interface{}, mutex *sync.Mutex) {
	for i := 0; i < 4; i++ {
		mutex.Lock()
		fmt.Println(data,i)

		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
//=============================================END Groutine 1==============================
//=============================================Groutine 2 ===============================
func process2(data interface{}, mutex *sync.Mutex) {
	for i := 0; i < 4; i++ {
		mutex.Lock()
		fmt.Println(data,i)
		
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond)
	}
}
//=====================================================end G 2============================

func main() {
	data1 := "[bisa1 bisa2 bisa3]" // data groutine 
	data2 := "[coba1 coba2 coba3]" // data groutine

	mutex := &sync.Mutex{}

	for i := 0; i < 2; i++ {
		go process1(data1, mutex)
		go process2(data2, mutex)
	}

	time.Sleep(5 * time.Second)
}


// belum selesai