package logger

import "fmt"

func Info(action string, value string) {
	fmt.Println(action, value)
}
