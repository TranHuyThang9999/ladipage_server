package main

import (
	"log"
	"time"
)

func FormatTime(req time.Time) string {
	return req.Format("2006-01-02 15:04:05")
}

// 2025-01-19 07:06:59.534307 +00:00
func main() {
	timeNow := time.Now()
	log.Println("time now : ", timeNow)
	log.Println("time : ", FormatTime(timeNow))
}
