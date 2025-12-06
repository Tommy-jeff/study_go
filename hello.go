package main

import (
	"fmt"

	"github.com/Tommy-jeff/study_go/math"
	"github.com/Tommy-jeff/study_go/requests"
	"github.com/Tommy-jeff/study_go/objects"
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

	resultadoSoma := math.Add(10, 15)
	fmt.Printf("\nResultado da soma: %v\n", resultadoSoma)

	resultadoSubtracao := math.Subtract(30, 12)
	fmt.Printf("Resultado da subtração: %v\n", resultadoSubtracao)

	resultadoSomaAll := math.AddAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("Resultado da soma de todos: %v\n", resultadoSomaAll)

	varEu := math.Eu
	fmt.Printf("Valor da variável eu: %v\n", varEu)


	// HTTP Request with error handling
	res1, err := requests.GoogleRequest()
	if err != nil {
		fmt.Printf("Error during googleRequest: %v\n", err)
		return
	}
	fmt.Printf("Response status code: %v\n %v\n", res1.StatusCode, res1.Header)

	// HTTP Request without error handling
	res2, _ := requests.GoogleRequest()
	
	fmt.Printf("Response status code: %v\n %v\n", res2.StatusCode, res2.Header)


	// Anonimous function
	anonimousFunc := func(name string) string{
		return fmt.Sprintf("Hello %v, from anonimous function!", name)
	}
	fmt.Println(anonimousFunc("Tommy"))
	
	// Anonimous add all funcition
	anonimousAddAll := func(numbers ...int) (result int){
		for _, v := range numbers {
			result += v
		}
		return
	}
	fmt.Printf("Resultado da soma de todos (anonimous): %v\n", anonimousAddAll(1,2,3,4,5,6,7,8,9,10))
	
	// Instances of structs
	carro1 := objects.NewCar("Ferrari", "F40", 1992)
	fmt.Println(carro1.GetInfo())
}