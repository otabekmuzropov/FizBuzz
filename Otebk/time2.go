package main

import (
	"fmt"
	"time"
)

func main () {

	var temp string

	var date time.Time

	x := time.Now()

	layout := "01/02/2006 15-04-05"
	temp = "01/23/2020 10-10-10"
	date, _ = time.Parse(layout, temp)

	date1 := time.Date(2020, time.February, 1, 9, 45, 3, 32, time.UTC)
	date2 := time.Date(2020, time.March, 1, 9, 45, 3, 32, time.UTC)

	fmt.Println(date2.Sub(date1).Hours() / 24)

	fmt.Println(time.ParseDuration("s"))

	fmt.Println(date, "\n", x)

	fmt.Println(int(time.Now().Month()))
}