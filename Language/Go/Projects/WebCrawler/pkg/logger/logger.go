package logger

import (
	"fmt"
	"time"
)

func Info(msg string) {
	var str string = fmt.Sprintf("%s - [INFO] - %s", time.Now().Format(time.RFC3339), msg)
	fmt.Println(str)
}

func Error(msg string) {
	var str string = fmt.Sprintf("%s - [ERROR] - %s", time.Now().Format(time.RFC3339), msg)
	fmt.Println(str)
}

func Warn(msg string) {
	var str string = fmt.Sprintf("%s - [WARNING] - %s", time.Now().Format(time.RFC3339), msg)
	fmt.Println(str)
}
