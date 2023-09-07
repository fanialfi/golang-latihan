package pomodoro

import (
	"fmt"
	"time"
)

func CreateDuration(message string) time.Duration {
	var strDuration string
	fmt.Printf("%s ", message)
	fmt.Scan(&strDuration)

	duration, err := time.ParseDuration(strDuration)

	if err != nil {
		fmt.Println("format anda tidak sesuai")
		CreateDuration(message)
	}
	return duration
}

func CreatePomodoro(duration time.Duration) {
	var zero time.Duration = 0
	step := time.Second

	if duration < zero {
		return
	}

	for i := duration; i > zero; i = i - step {
		fmt.Println(i, "remaning")
		time.Sleep(step)
	}
}
