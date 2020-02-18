package main

import (
	"fmt"
	"strings"
)

func Palindrom(word string) string {
	var str string
	str = "weeeeeeeeee"
	for i := 0; i <= len(word) / 2; i++ {
		if strings.ToUpper(string(word[i])) == strings.ToUpper(string(word[len(word) - 1 - i])) {
			str = "This word polindrom"
		}else {
			str = "This word no polindrom"
			break
		}

	}
	return str
}

func main ()  {

	fmt.Println(Palindrom("Hello"))
	fmt.Println(Palindrom("Level"))
}

