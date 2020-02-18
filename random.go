package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main ()  {
	x := rand.NewSource(time.Now().UnixNano())
	fmt.Println(rand.New(x).Intn(5))
}
