package studysteps

import "fmt"

func AnonimousFunctions() {
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
}