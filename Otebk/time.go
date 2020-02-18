package main

import (
	"fmt"
	"time"
)

func main () {
	today := time.Now()
	fmt.Println(today)
	day := time.Now().Day()
	fmt.Print(day)
	month := time.Now().Month()
	fmt.Print(month)
	hours, _ := time.ParseDuration("1hour")
	fmt.Println(hours)
	fmt.Println(hours.Seconds(), hours.Nanoseconds())
	var temp string
	var date time.Time
	var err error
	temp = "2020/01/01 10:32:32"
	layout := "2020/02/02 10:10:10"
	date, err = time.Parse(layout, temp)
	fmt.Print(date, err)
}
