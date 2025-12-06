package main

import (
	"fmt"

	"github.com/Tommy-jeff/study_go/math"
	"github.com/Tommy-jeff/study_go/requests"
)

func main() {
	// a := "Oi"
	// b := 24
	// c := false
	// d := 3.14
	// e := `Olá, 
	
	
	// Mundo!`
	// fmt.Printf("\n%T\n%T\n%T\n%T\n%T\n", a, b, c, d, e)
	// fmt.Printf("\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e)

	resultado := math.Add(10, 15)
	fmt.Printf("\nResultado da soma: %v\n", resultado)

	resultado2 := math.Subtract(30, 12)
	fmt.Printf("Resultado da subtração: %v\n", resultado2)

	eu := math.Eu
	fmt.Printf("Valor da variável eu: %v\n", eu)

	res, err := requests.GoogleRequest()
	if err != nil {
		fmt.Printf("Error during googleRequest: %v\n", err)
		return
	}

	fmt.Printf("Response status code: %v\n %v\n", res.StatusCode, res.Header)

}