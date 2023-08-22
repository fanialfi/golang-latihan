package lib

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	Id       time.Time
	Nama     string
	Email    string
	Hp       int
	Password string
	Data     chan string
}

func (u *User) ChangeEmail(email string) {
	const cek = "@"
	data := strings.TrimSpace(email)

	if checked := strings.Contains(data, cek); checked {
		u.Email = data
		fmt.Printf("email berhasil diubah ke \"%s\"\n", data)
	} else {
		fmt.Println(fmt.Errorf("email tidak masuk dalam kriteria"))
	}
}

func (u *User) ChangePassword(password string) {
	if valid, err := validateLengthPassword(password); valid {
		u.Password = password
	} else {
		fmt.Println(err.Error())
	}
}

func (u *User) SendTelphone(r *User) {
	for i := 0; i < 10; i++ {
		now := time.Now()
		h, m, s := now.Clock()
		r.Data <- fmt.Sprintf("jam %d:%d:%d:%d - sending message %d", h, m, s, now.Nanosecond(), i)
		time.Sleep(time.Second)
	}
	close(r.Data)
}

func (u *User) ReceiveTelphone() {
	for data := range u.Data {
		fmt.Printf("%s - receive message : %s\n", u.Nama, data)
	}
}

// function untuk mencetak detail isi variabel yang dibuat dari struct User
func CetakDetail(data *User) {
	fmt.Printf("\nNama\t\t: %s\nEmail\t\t: %s\nHp\t\t: %d\nPassword\t: %s\n\n", data.Nama, data.Email, data.Hp, data.Password)
}

// jika panjang password kurang dari 8, maka akan mengembalikan error
func validateLengthPassword(data string) (bool, error) {
	const panjang = 8
	password := strings.TrimSpace(data)

	if len(password) < panjang {
		return false, fmt.Errorf("panjang password minimal %v, panjang password kamu %d", panjang, len(password))
	}
	return true, nil
}
