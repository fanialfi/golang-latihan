package lib

import "fmt"

type Employee struct {
	Id       int
	Name     string
	Position string
	Gaji     float64
	Bonus    float64
}

func DisplayEmploye(employe *Employee) {
	fmt.Printf("nama\t\t: %s\nposition\t: %s\ngaji\t\t: %.3f\nbonus\t\t: %.3f\n", employe.Name, employe.Position, employe.Gaji, employe.Bonus)
}

func ReturnBonus(employee *Employee, bonus float64) {
	employee.Bonus = bonus
}

func CalculateGaji(employe *Employee) {
	switch employe.Position {
	case "back-end enginer":
		employe.Gaji = employe.Gaji + (employe.Gaji / 25.0)
	case "back-end developer":
		employe.Gaji = employe.Gaji + (employe.Gaji / 10.0)
	}
}
