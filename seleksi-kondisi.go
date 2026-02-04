package main

import (
	"fmt"
)

func main() {
	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Println("tidak lulus. nilai anda %d\n", point)
	}

	var pointTwo = 8840.0

	/*
	   Deklarasi variabel temporary hanya bisa dilakukan lewat metode type interface yang mengguankan tanda :=.
	   Penggunaan var untuk kode di bawah hanya akan membuat error.
	*/
	if percent := pointTwo / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

	// Swift - case
	var pointThree = 6
	switch pointThree {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// case untuk banyak kondisi
	switch pointThree {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// kurung kurawal pada keyword case & default
	switch pointThree {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		{
			fmt.Println("awesome")
			fmt.Println("kulukulu")
		}
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you can be better")
		}
	}

	// switch dengan gaya if-else
	switch {
	case pointThree == 8:
		fmt.Println("perfect")
	case (pointThree < 8) && (pointThree > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// penggunaan fallthrough pada switch
	/*
	   keyword fallthrough digunakan untuk memaksa proses pengecekan tetap diteruskan ke case selanjutnya dengan tanpa menhiraukan nilai kondisinya,
	   efeknya case selanjutnya diaggap true meskipun nilainya tidak memenuhi.
	*/
	switch {
	case pointThree == 8:
		fmt.Println("perfect")
	case (pointThree < 8) && (point > 3):
		fmt.Println("awesome") // tetap akan tampil
		fallthrough
	case pointThree < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}

	// seleksi kondisi bersarang
	var pointFour = 10
	if pointFour > 7 {
		switch pointFour {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if pointFour == 5 {
			fmt.Println("not bad")
		} else if pointFour == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it!")
			if pointFour == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
