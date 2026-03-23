package main

import (
	"fmt"
	"golang/mathutil"
	"math"
)

func main() {
	fmt.Println(mathutil.Add(5, 5))
	fmt.Println(mathutil.Sub(5.3, 5.1))
	fmt.Println(math.Pow(2, 3))
	fmt.Println(math.Sqrt(144))
	var numero int = 50
	fmt.Println(numero)
	var a, b, c = 3, 4, 5
	fmt.Println(a, b, c)
	nombre, edad := "Zully", 37
	fmt.Printf("mi nombre es %s y tengo %d anios", nombre, edad)
}
