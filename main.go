package main

import "fmt"

type Employee struct {
	id       int
	name     string
	position string
	gaji     float64
	bonus    float64
}

func displayEmploye(employe Employee) {
	fmt.Printf("nama\t\t: %s\nposition\t: %s\ngaji\t\t: %.3f\nbonus\t\t: %.3f\n", employe.name, employe.position, employe.gaji, employe.bonus)
}

func returnBonus(employee *Employee, bonus float64) {
	employee.bonus = bonus
}

func main() {
	var employe1 = Employee{
		id:       1,
		name:     "employe 1",
		position: "back-end developer",
		gaji:     5000000.00,
		bonus:    0.0,
	}
	var employe2 = Employee{
		id:       2,
		name:     "employe 2",
		position: "back-end enginer",
		gaji:     12000000.00,
		bonus:    0.0,
	}
	fmt.Printf("deskripsi employe :\n\n")
	displayEmploye(employe1)
	fmt.Println()
	displayEmploye(employe2)

	returnBonus(&employe1, 1000000.0)
	returnBonus(&employe2, 2000000.0)

	fmt.Printf("\ndeskripsi employee setelah menerima bonus :\n")
	displayEmploye(employe1)
	fmt.Println()
	displayEmploye(employe2)
}
