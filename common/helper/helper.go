package helper

import (
	"fmt"
	"time"
)

//GetStrDateTime Function
func GetStrDateTime() string {
	dt := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("%-19s", dt)
	return msg
}

//GetStrtimestamp Function
func GetStrtimestamp() string {
	dt := time.Now().Format("20060102150405.9999")
	msg := fmt.Sprintf("%s", dt)
	return msg
}

func BytesToString(data []byte) string {
	return string(data[:])
}

func GetStrDate() string {
	dt := time.Now().Format("2006-01-02")
	msg := fmt.Sprintf("%-24s", dt)
	return msg
}
