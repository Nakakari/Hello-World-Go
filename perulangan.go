package main

import "fmt"

func main() {
	// di Go, perulangan hanya for saja.
	for i := 0; i < 5; i++ {
		fmt.Println("Angka: ", i)
	}

	// penggunaan keyword for dengan argumen hanya kondisi
	var i = 0
	for i < 5 {
		fmt.Println("Angka: ", i)
		i++
	}

	// penggunaan keyword for tanpa argumen
	var j = 0
	for {
		fmt.Println("Angka: ", j)
		j++
		if j == 5 {
			break
		}
	}

	// penggunaan keyword for - range
	var xs = "123" // string
	for i, v := range xs {
		fmt.Println("Index=", i, "Value", v)
	}

	var ys = [5]int{10, 20, 30, 40, 50} // array
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	var zs = ys[0:2] // slice
	for _, v := range zs {
		fmt.Println("Value=", v)
	}
}
