package main

import (
	"fmt"
	"latihan/pomodoro"
)

func main() {
	// latihan pomodoro
	message := []string{"masukkan lama pomodoro", "masukkan lama istirahat"}
	focus := pomodoro.CreateDuration(message[0])
	istirahat := pomodoro.CreateDuration(message[1])
	pomodoro.CreatePomodoro(focus)
	fmt.Print("tekan enter untuk melanjutkan ")
	fmt.Scanln()
	pomodoro.CreatePomodoro(istirahat)

	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("panic occured, message :", r)
	// 	} else {
	// 		fmt.Println("application running perfect")
	// 	}
	// }()

	// var employe1 = lib.Employee{
	// 	Id:       1,
	// 	Name:     "employe 1",
	// 	Position: "back-end developer",
	// 	Gaji:     5000000.00,
	// }
	// var employe2 = lib.Employee{
	// 	Id:       2,
	// 	Name:     "employe 2",
	// 	Position: "back-end enginer",
	// 	Gaji:     12000000.00,
	// }
	// fmt.Printf("deskripsi employe :\n")
	// lib.DisplayEmploye(&employe1)
	// fmt.Println()
	// lib.DisplayEmploye(&employe2)

	// lib.ReturnBonus(&employe1, 1000000.0)
	// lib.ReturnBonus(&employe2, 2000000.0)

	// fmt.Printf("\ndeskripsi employee setelah menerima bonus :\n")
	// lib.DisplayEmploye(&employe1)
	// fmt.Println()
	// lib.DisplayEmploye(&employe2)

	// fmt.Printf("\ncalculate gaji berdasarkan position :\n")
	// lib.CalculateGaji(&employe1)
	// lib.CalculateGaji(&employe2)

	// lib.DisplayEmploye(&employe1)
	// fmt.Println()
	// lib.DisplayEmploye(&employe2)

	// fmt.Println()

	// var fani = lib.User{
	// 	Id:       time.Now(),
	// 	Nama:     "fani",
	// 	Email:    "fani@fani.com",
	// 	Hp:       1234567890,
	// 	Password: "passwordFani12345678",
	// 	Data:     make(chan string, 3),
	// }

	// var alfi = lib.User{
	// 	Id:       time.Now(),
	// 	Nama:     "alfi",
	// 	Email:    "alfi@alfi.com",
	// 	Hp:       9876543210,
	// 	Password: "passwordAlfi87654321",
	// 	Data:     make(chan string, 3),
	// }

	// lib.CetakDetail(&fani)
	// lib.CetakDetail(&alfi)

	// fani.ChangeEmail("fani@fanialfi.space")
	// alfi.ChangeEmail("alfi@fanialfi.space")

	// lib.CetakDetail(&fani)
	// lib.CetakDetail(&alfi)

	// go fani.SendTelphone(&alfi)
	// alfi.ReceiveTelphone()

	// go alfi.SendTelphone(&fani)
	// fani.ReceiveTelphone()

	// alfi.ChangeEmail("email baru")
	// fani.ChangeEmail("email baru")

	// lib.CetakDetail(&fani)
	// lib.CetakDetail(&alfi)

	// alfi.ChangePassword("1234567890")
	// fani.ChangePassword("12344567890")

	// lib.CetakDetail(&fani)
	// lib.CetakDetail(&alfi)

	// alfi.ChangePassword("123")
	// fani.ChangePassword("123")

	// lib.CetakDetail(&fani)
	// lib.CetakDetail(&alfi)
}
