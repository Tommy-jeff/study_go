package studysteps

import "fmt"

func TestsPointers() {

	t := 10

	var pointerT *int = &t

	*pointerT = 20

	fmt.Println("Value of t:", t)


}