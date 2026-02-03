package main

import "fmt"

func main(){
  const firstName string = "john"
  const lastName = "wick" // teknik interface
  fmt.Print("halo ", firstName, "!\n")
  fmt.Print("nice to meet your ", lastName, "!\n")

  /*
    perbedaan print vs println
    - println menghasilkan baris baru serta tidak memerlukan spasi lagi sebagai pemisah
    - print tidak menghasilkan baris baru
  */
  fmt.Println("jhon wick")
  fmt.Println("jhon", "wick")

  fmt.Print("jhon wick\n")
  fmt.Print("john", "wick\n")
  fmt.Print("john", " ", "wick\n")

  // Deklarasi multi konstanta
  const (
    square = "kotak"
    isToday bool = true
    numeric uint8 = 1
    floatNum = 2.2
  )

  const (
    a = "konstanta"
    b // memiliki value sama
  )

  const (
    today string = "Senin"
    sekarang
    isToday2 = true
  )

  // deklarasi multi konstanta dalam satu baris
  const satu, dua = 1, 2
  const three, four string = "tiga", "empat"
}
