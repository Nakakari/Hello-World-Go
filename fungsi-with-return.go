package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var randomValue int
	// var => deklarasi variabel, randomValue => nama variabel, int => tipe data variabel

	randomValue = randomWithRange(2, 10)
	fmt.Println("random number: ", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number: ", randomValue)
	randomValue = randomWithRange(2, 10)
	fmt.Println("random number: ", randomValue)
}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min+1) + min
	return value
}
