package main

import "fmt"

/*
	Pada satu package yang sama tidak baik untuk redeclare func main. Meskipun berhasil running jika pada file ini terdeklarsi main().
*/
func main(){
	var firstName string = "John"

	var lastName string
	lastName = "Wick"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// deklarasi tanpa type data
	var catFirstName string = "Fufu"
	catLastName := "Fafa"

	fmt.Printf("halo %s %s!\n", catFirstName, catLastName)

	// menggunakan var, tanpa type data, menggunakan perantara "="
	var companyFirstName = "Suka"
	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	companyLastName := "Belajar"
	fmt.Printf("Halo PT %s %s!\n", companyFirstName, companyLastName)

	// Deklari multi variable
	var first, second, thrid string
	first, second, thrid = "one", "two", "three"

	var fourth, fifth, sixth string = "empat", "lima", "enam"
	// seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	_ = "unkown"

	// dengan menggunakan teknik type interface, deklarasi multi variabel digunakan untuk variable-variabel yang tipe data satu sama lainnya berbeda
	// one, isFriday, twoPointTwo, say := 1, true, 2.2, "helo"

	fmt.Printf(first, second, thrid, fourth, fifth, sixth)


	name := new(string)
	fmt.Println(name) // 0x20818a220
	fmt.Println(*name) // ""

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negativ: %d\n", negativeNumber)

	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n",)

	var exist bool =true
	fmt.Printf("exist? %t \n", exist)

	var message string = "halo"
	fmt.Printf("message: %s \n", message)

	var messageBackticks = `Nama saya "Jhon Wick"
	Salam kenal.
	Mari belajar Goalng!
	`
	fmt.Println(messageBackticks)
}