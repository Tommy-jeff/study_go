package studysteps

import (
	"fmt"

	"github.com/Tommy-jeff/study_go/math"
)

func MathOperations() {
	resultadoSoma := math.Add(10, 15)
	fmt.Printf("\nResultado da soma: %v\n", resultadoSoma)

	resultadoSubtracao := math.Subtract(30, 12)
	fmt.Printf("Resultado da subtração: %v\n", resultadoSubtracao)

	resultadoSomaAll := math.AddAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Printf("Resultado da soma de todos: %v\n", resultadoSomaAll)

	varEu := math.Eu
	fmt.Printf("Valor da variável eu: %v\n", varEu)
}