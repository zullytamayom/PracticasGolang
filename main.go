package main

import (
	"errors"
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
	nombre := "Zully"
	fmt.Printf("mi nombre es %s y tengo %d anios", nombre, mathutil.Edad)
	fmt.Println("-----------------------------")

	fmt.Println(mathutil.Multiplicar(5))
	fmt.Println("-----------------------------")
	//fmt.Println(sumar(5, 8))
	total := sumar(2, 5)
	fmt.Println(total)

	resultadoDivision, err := dividir(5, 2)
	if err != nil {
		fmt.Println("Hubo un error")
	}
	fmt.Println(resultadoDivision)

}

func sumar(x, y int) int {
	return x + y
}

func dividir(x, y int) (int, error) {
	if y <= 0 {
		return 0, errors.New("No se puede dividir entre cero.")
	}
	return x / y, nil
}
