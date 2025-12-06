package studysteps

import (
	"fmt"

	"github.com/Tommy-jeff/study_go/objects"
)

func Pointers() {
	// var example
	a:= 10

	// pointer example: pointerA stores the memory address of variable a
	var pointerA *int = &a

	// Address is accessed using & and value at address using *
	fmt.Println("Memory address of a:", pointerA)
	fmt.Println("Value in address of a:", *pointerA)

	// Ok, let's change the value of a using the pointer
	*pointerA = 20
	fmt.Println("New value of a:", a)

	// See that both a and *pointerA reflect the new value, because they point to the same memory location
	fmt.Println("Value in address of a:", *pointerA)

	// Now, let's use a function that modifies the value of a using its pointer
	PointersExample(&a)
	fmt.Println("Value of a after PointersExample function:", a)

	ferrari := objects.NewCar("Ferrari", "F40", 1992)
	fmt.Println(ferrari.GetInfo())
	// Note that the year will be changed to 2050 because GetInfo has a pointer receiver and modifies the Ano field directly

}

// PointersExample adds 10 to the value at the memory address passed, see i don't use return because i'm modifying the value directly via pointer
func PointersExample(a *int) {
	*a = *a + 10
}