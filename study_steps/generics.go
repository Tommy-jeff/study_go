package studysteps

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	~int64 | ~float64
}

type MyNumber int64

func SomaGenerica[T Number] (m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Soma[T Number] (number1 T, number2 T) T {
	return number1 + number2
}

// o tipo comparable permite comparar os valores mesmo sendo genericos, o que nao e possivel com tipos genericos normais
// porém, ele suportas apenas operações de comparação de igualdade (==, !=)
func Compare[T comparable] (number1 T, number2 T) T {
	if number1 == number2 {
		return number1
	}
	return number2
}


// usando a biblioteca constraints para permitir comparações maiores e menores
// nesse caso, o tipo O deve implementar a interface constraints.Ordered que permite comparações de ordem (>, <, >=, <=)
// então ela aceita tipos como int, float, string, etc.
func compareOrdered[T constraints.Ordered] (item1 T, item2 T) T {
	if item1 > item2 {
		return item1
	}
	return item2
}

// Uma interface chamada stringer, que exige apenas um método
type stringer interface {
	String() string
}

// Um tipo próprio chamado MyString, baseado em string
type MyString string
// Esse tipo implementa o método String(), logo satisfaz a interface
func (s MyString) String() string {
	return string(s)
}

// Uma função genérica que recebe um slice cujo tipo T é qualquer tipo que implemente a interface stringer
func Concat[T stringer] (values []T) string {
	result := ""
	for _, v := range values {
		result += v.String()
	}
	return result
} 

func GenericsExample() {

	fmt.Println(Concat([]MyString{"Hello, ", "world", "!"}))

	fmt.Println(compareOrdered(10, 20))
	fmt.Println(compareOrdered("a", "b"))

	fmt.Println(Soma[int64](10, 20))

	fmt.Println(SomaGenerica(map[string]MyNumber{
		"A": 10,
		"B": 20,
		"C": 30,
	}))
	
	fmt.Println(SomaGenerica(map[string]float64{
		"A": 14.5,
		"B": 22.3,
		"C": 55.2,
	}))

}