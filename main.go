package main

import (
	"fmt"
	"math"
)

func main() {
	// mencari bilangan prima
	fmt.Printf("masukkan bilangan yang akan di cek : ")
	var bilangan int
	fmt.Scan(&bilangan)
	var result = cekBilanganPrima(bilangan)
	fmt.Println("hasil", result)

	// hitung isi bola
	var r float64

	fmt.Printf("masukkan jari jari : ")
	fmt.Scan(&r)

	var isiBola = hitungIsiBola(r)

	fmt.Printf("isi bola adalah %.3f\n", isiBola)

	fmt.Printf("masukkan suhu dalam satuan celcius : ")
	var celcius float64
	fmt.Scan(&celcius)

	var fahrenheit = konversiSuhu(celcius, false)
	fmt.Printf("konversi suhu dari %.2f celcius ke fahrenheit : %.2f\n", celcius, fahrenheit)
}

func konversiSuhu(s float64, bol bool) float64 {
	var a float64 = (s * (9.0 / 5.0)) + float64(32)

	return a
}

func hitungIsiBola(r float64) float64 {
	var x = 4 / 3

	pi := float64(math.Pi)
	return float64(x) * pi * float64(math.Pow(r, 3))
}

func cekBilanganPrima(number int) int {
	var prima int
	for i := 2; i <= number; i++ {
		if isOne := float64(number) / float64(i); isOne == 1.0 {
			fmt.Printf("tipe %T value %f\n", isOne, isOne)
			prima = number
		}
	}
	return prima
}
