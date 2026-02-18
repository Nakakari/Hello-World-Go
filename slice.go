package main

import "fmt"

/*
	Slice merupakan reference elemen array.
*/

func main() {
	// inisialisasi slice
	var fruits = []string{"apple"}
	fmt.Println(fruits[0]) // apple

	/*
		Salah satu perbedaan slice dan array bisa diketahui pada saat deklarasi variabel-nya, jika jumlah elemen
		tidak dituliskan, maka variabel tersebut adalah slice.
	*/
	var fruitsA = []string{"apple", "grape"}     // slice
	var fruitsB = [2]string{"banana", "melon"}   // array
	var fruitsC = [...]string{"papaya", "grape"} // array
	fmt.Println(fruitsA, "\t", fruitsB, "\t", fruitsC, "\t")

	// hubungan slice dengan array dan operasi slice
	var fruitsOne = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruitsOne[0:2] // dimulai dari indeks ke-0, hingga elemen sebelum indeks ke-2.
	fmt.Println(newFruits)

	/*
		Kode					|		Output					|		Penjelasan
		------------------------|-------------------------------|-------------------------------------------------------------------------------------
		fruits[0:2]  			| [apple, grape]				|	Semua elemen mulai indeks ke-0, hingga sebelum indeks ke-2
		fruits[0:4]				| [apple, grape, banana, melon]	|	semua elemen mulai dari indeks ke-0, hingga sebelum indeks ke-4
		fruits[0:0]				| []							|	menghasilkan slice kosong, karena tidak ada elemen sebelum indeks ke-0
		fruits[4:4]				| []							| 	menghasilkan slice kosong, karena tidak ada elemen yang dimulai dari indeks ke-4
		fruits[4:0]				| []							|	error, pada penulisan fruits[a:b] nilai a harus lebih kecil atau sama dengan b
		fruits[:]				| [apple, grape, banana, melon]	|	semua elemen
		fruits[2:]				| [banana, melon]				| 	semua elemen mulai dari indeks ke-2
		fruits[:2]				| [apple, grape]				|	semua elemen hingga sebelum indeks ke-2
	*/

	// Slice merupakan tipe data reference
	var fruitsForReference = []string{"apple", "grape", "banana", "melon"}

	var aFruitsForReference = fruitsForReference[0:3]
	var bFruitsForReference = fruitsForReference[1:4]

	var aaFruitsForReference = aFruitsForReference[1:2]
	var baFruitsForReference = bFruitsForReference[0:1]

	fmt.Println(fruitsForReference)   // [apple grape banana melon]
	fmt.Println(aFruitsForReference)  // [apple grape banana]
	fmt.Println(bFruitsForReference)  // [grape banana melon]
	fmt.Println(aaFruitsForReference) // [grape]
	fmt.Println(baFruitsForReference) // [grape]

	// buah "grape" diubah menjadi "pinnaple"
	baFruitsForReference[0] = "pinaple"

	fmt.Println(fruitsForReference)   // [apple pinaple banana melon]
	fmt.Println(aFruitsForReference)  // [apple pinaple banana]
	fmt.Println(bFruitsForReference)  // [pinaple banana melon]
	fmt.Println(aaFruitsForReference) // [pinaple]
	fmt.Println(baFruitsForReference) // [pinaple]

	// fungsi len() => menghitung jumlah elemen slice
	var fruitsLen = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruitsLen))

	// cap => menghitung lebar atau kapasitas maksimum slice.
	// nilai kembalian fungsi ini untuk slice yang baru dibuat pasti sama dengan len,
	// tapi bisa berubah seiring operasi slice yang dilakukan.
	var fruitsCap = []string{"apple", "banana", "manggo", "manggis"}
	fmt.Println(len(fruitsCap)) // len : 4
	fmt.Println(cap(fruitsCap)) // cap : 4

	var aFruitsCap = fruitsCap[0:3]
	fmt.Println(len(aFruitsCap)) // len : 3
	fmt.Println(cap(aFruitsCap)) // cap : 4

	var bFruitsCap = fruitsCap[1:4]
	fmt.Println(len(bFruitsCap)) // len : 3
	fmt.Println(cap(bFruitsCap)) // cap : 3

	// append
	// digunakan untuk menambahkan elemen pada slice.
	var fruitsAppend = []string{"apple", "grape", "banana"}
	var cFruitsAppend = append(fruitsAppend, "papaya")

	fmt.Println(fruitsAppend)  // [apple, grape, banana]
	fmt.Println(cFruitsAppend) // [apple, grape, banana, papaya]

	var bFruitsAppend = fruitsAppend[0:2]

	fmt.Println(cap(bFruitsAppend)) // 3
	fmt.Println(len(bFruitsAppend)) // 2

	fmt.Println(fruitsAppend)  // [apple, grape, banana]
	fmt.Println(bFruitsAppend) // [apple, grape]

	var dFruitsAppend = append(bFruitsAppend, "papaya")

	fmt.Println(fruitsAppend)  // [apple, grape, banana]
	fmt.Println(bFruitsAppend) // [apple, grape]
	fmt.Println(dFruitsAppend) // [apple, grape, papaya]

	// copy
	dst := make([]string, 3)
	src := []string{"watermelon", "pinapple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // [watermelon, pinapple, apple]
	fmt.Println(src) // [watermelon, pinapple, apple, orange]
	fmt.Println(n)   // elemen yang berhasil dicopy dst = 3

	dst2 := []string{"potato", "potato", "potato"}
	src2 := []string{"watermelon", "pinapple"}
	n2 := copy(dst, src)

	fmt.Println(dst2) // [watermelon, pinapple, potato]
	fmt.Println(src2) // [watermelon, pinapple]
	fmt.Println(n2)   // 2

	// pengaksesan elemen slice dengan 3 indeks
	// ex fruits[0:1:1]
	// angka kapasitas tidak boleh melebihi kapasitas slice yang akan dislicing.

	var fruitsCapIndex = []string{"apple", "grape", "banana"}
	var aFruitsCapIndex = fruitsCapIndex[0:2]
	var bFruitsCapIndex = fruitsCapIndex[0:2:2]

	fmt.Println(fruitsCapIndex)      // [apple, grape, banana]
	fmt.Println(len(fruitsCapIndex)) // 3
	fmt.Println(cap(fruitsCapIndex)) // 3

	fmt.Println(aFruitsCapIndex)      // [apple, grape]
	fmt.Println(len(aFruitsCapIndex)) // 2
	fmt.Println(cap(aFruitsCapIndex)) // 3

	fmt.Println(bFruitsCapIndex)      // [apple, grape]
	fmt.Println(len(bFruitsCapIndex)) // 2
	fmt.Println(cap(bFruitsCapIndex)) // 2
}
