package main

import "fmt"

func FizzBuzz(number int) string {
	var str string
	if number % 5 ==0 && number % 3 != 0{
		str = "This number us Fiz"
	}else if number % 5 != 0 && number % 3 == 0 {
		str = "This number is Buzz"
	}else if number % 5 ==0 && number % 3 == 0 {
		str = "This number is FizBuzz"
	}else {
		str = "This number isn't Fiz or Buzz"
	}
	return str
}

func main ()  {

	fmt.Println(FizzBuzz(15))

}