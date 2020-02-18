package main

import "fmt"

func Summ(arr[] int) (string, int, string, int) {

	odd_sum, even_sum := 0, 0

	for _, num := range arr {
		if num % 2 == 0 {
			even_sum += num
		} else {
			odd_sum +=num
		}
	}
	return "Odd numbers summ: ", odd_sum, "\nEven numbers sum: ", even_sum
}


func main ()  {

	arr := [] int {21, 32, 43, 54, 65, 17, 38}

	fmt.Print(Summ(arr))

}