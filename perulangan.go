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

	var kvs = map[byte]int{'a': 0, 'b': 1, 'c': 2} //map
	for k, v := range kvs {
		fmt.Println("key=", k, "value=", v)
	}

	// boleh juga baik k dan atau v-nya diabaikan
	for range kvs {
		fmt.Println("done")
	}

	// selain itu bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
		fmt.Println(i)
	}

	/*
		Penggunaan keyword break & continue
	*/
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}

	/*
		Perulangan bersarang.
	*/
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Println(j, " ")
		}

		fmt.Println()
	}

	// Pemanfaatan label dalam perulangan
outerloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerloop
			}

			fmt.Println("matriks [", i, "][", j, "]", "\n")
		}
	}
}
