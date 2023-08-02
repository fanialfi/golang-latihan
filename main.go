package main

import (
	"fmt"
	"latihan/lib"
)

func main() {
	var employe1 = lib.Employee{
		Id:       1,
		Name:     "employe 1",
		Position: "back-end developer",
		Gaji:     5000000.00,
	}
	var employe2 = lib.Employee{
		Id:       2,
		Name:     "employe 2",
		Position: "back-end enginer",
		Gaji:     12000000.00,
	}
	fmt.Printf("deskripsi employe :\n")
	lib.DisplayEmploye(&employe1)
	fmt.Println()
	lib.DisplayEmploye(&employe2)

	lib.ReturnBonus(&employe1, 1000000.0)
	lib.ReturnBonus(&employe2, 2000000.0)

	fmt.Printf("\ndeskripsi employee setelah menerima bonus :\n")
	lib.DisplayEmploye(&employe1)
	fmt.Println()
	lib.DisplayEmploye(&employe2)

	fmt.Printf("\ncalculate gaji berdasarkan position :\n")
	lib.CalculateGaji(&employe1)
	lib.CalculateGaji(&employe2)

	lib.DisplayEmploye(&employe1)
	fmt.Println()
	lib.DisplayEmploye(&employe2)
}
