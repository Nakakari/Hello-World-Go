package main

import "fmt"

func main() {
	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"])
	fmt.Println("mei", chicken["mei"])

	// inisialisasi nilai map
	/**
	var data map[string]int
	data["one"] = 1
	// akan muncul error karena ketiadaan {}
	**/
	var data map[string]int
	data = map[string]int{}
	data["one"] = 1
	fmt.Println("one ", data["one"])

	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}
	fmt.Println(chicken1)

	// cara vertical
	var chicken2 = map[string]int{
		"januari":  50,
		"februari": 40,
	}
	fmt.Println(chicken2)

	/*
			   Variabel map bisa diinisialisasikan dengan tanpa nilai awal, caranya menggunakan tanda kurung kurawal, contoh: map[string]int{}.
			   Atau bisa juga dengan menggunakan keyword make dan new.
		     Khusus keyword new, yang dihasilkan merupakan pointer.
	*/
	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)
	fmt.Println(chicken3, chicken4, chicken5)

	// iterasi item map menggunakan for - range
	// yang dikembalikan ialah key dan value, bukan indeks dan elemen
	var chicken6 = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
		"april":    67,
	}
	for key, val := range chicken6 {
		fmt.Println(key, " \t:", val)
	}

	// menghapus item map
	var chicken7 = map[string]int{"januari": 50, "februari": 40}
	fmt.Println(len(chicken7)) // 2
	fmt.Println(chicken7)
	delete(chicken7, "januari")
	fmt.Println(len(chicken7))
	fmt.Println(chicken7)

	// deteksi keberadaan item dengan key tertentu
	var chicken8 = map[string]int{"januari": 50, "februari": 40}
	var value, isExist = chicken8["mei"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item is not exist")
	}

	// kombinasi slice dan map
	var chickens = []map[string]string{
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}
	for _, chichickenSlice := range chickens {
		fmt.Println(chichickenSlice["gender"], chichickenSlice["name"])
	}

	// penulisan tipe data pada item adalah opsional
	var chickensOpsional = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	var data2 = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
	fmt.Println(chickensOpsional, data2)
}
