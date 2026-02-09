package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "Setyo"
	names[1] = "Tuhu"
	names[2] = "Bagio"
	names[3] = "Mulyo"

	fmt.Println(names[0], names[1], names[2], names[3])

	// inisialisasi nilai awal array
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	// inisialisasi nilai array dengan gaya vertikal
	var fruitsTwo [4]string
	// cara horizontal
	fruitsTwo = [4]string{"cherry", "strowberi", "markisa", "rambutan"}
	// cara vertikal
	fruitsTwo = [4]string{
		"cherry",
		"strowberi",
		"markisa",
		"semangka",
	}

	fmt.Println(fruitsTwo)

	// inisialisasi nilai awal array tanpa jumlah emelemn
	var numbers = [...]int{2, 3, 4, 2, 3, 4}
	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen: \t:", len(numbers))

	// array multidimensi
	var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 1}, {6, 5, 4}}
	fmt.Println("numbers1: ", numbers1)
	fmt.Println("numbers2: ", numbers2)

	// perulangan elemen array menggunakan keyword for
	var fruitsFour = [4]string{"apple", "mango", "melon", "grape"}
	for i := 0; i < len(fruitsFour); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruitsFour[i])
	}

	// perulangan elemen array menggunakan keyword for - range
	var fruitsFives = [4]string{"apple", "banana", "pinaple", "cherry"}
	for i, fruitsFive := range fruitsFives {
		fmt.Printf("elemen %d : %s\n", i, fruitsFive)
	}

	// penggunaan _ untuk mengskip variable index yang tidak digunakan
	var fruitsSixs = [4]string{"banana", "markisa", "manggo", "apple"}
	for _, frufruitsSix := range fruitsSixs {
		fmt.Printf("nama buah : %s\n", frufruitsSix)
	}
	// jika yang tidak digunakan nilainya
	// for i, _ := range fruitsSixs {
	// }
	// atau
	// 	for i := range fruitsSixs {
	// 	}

	// alokasi elemen array menggunakan keyword make
	var fruitsSeven = make([]string, 2)
	fruitsSeven[0] = "apple"
	fruitsSeven[1] = "manggo"
	fmt.Println(fruitsSeven)
}
