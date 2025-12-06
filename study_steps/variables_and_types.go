package studysteps

import "fmt"

func VariablesAndTypes() {
	a := "Oi"
	b := 24
	c := false
	d := 3.14
	e := `Olá, 
	
	
	Mundo!`
	
	fmt.Printf("\n%T\n%T\n%T\n%T\n%T\n", a, b, c, d, e)
	fmt.Printf("\n%v\n%v\n%v\n%v\n%v\n", a, b, c, d, e)
}