package main

import (
	"bufio"
	"fmt"
	"os"
)

func main ()  {
	w := bufio.NewWriter(os.Stdin)
	fmt.Println(w, "Hello, ", "world")

}
