package main

import "fmt"
func main(){
//============================================= START PERULANGAN I =========================================
i := 0
    for i <= 4 {
        fmt.Println("Nilai i =", i)
        i++

       
    }
//============================================END PERULANGAN I================================
//============================================START PERULANGAN J ============================ 
j := 0
    for j <= 10 {
        if j == 5 { // menghilangkan angka 5 pada perulangan J
           //====================================== PERULANGAN KARAKTER =======================
            
          var w = []rune{}
          w = append(w, 'C','A','w','A','N',' ','S','U','C','I')
          nn := 0
          for _, value := range w {
              fmt.Printf("character %U \"%c\" starts at byte position %d\n", value, value, nn)
              nn += 2
      
          }
          //===============================================END=============================
            j++
           continue
        }

        
        fmt.Println("Nilai j =", j)
        j++
    }
input := "САШАРВО"
for index, value := range input{
    fmt.Printf("character %U \"%c\" starts at byte position %d\n", value, value, index)
    

}
    
}







