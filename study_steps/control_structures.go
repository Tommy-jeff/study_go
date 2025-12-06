package studysteps

import (
	"fmt"

	"github.com/Tommy-jeff/study_go/objects"
)

func ControlStructures() {
	// Instances of structs
	carro1 := objects.NewCar("Ferrari", "F40", 1992)
	fmt.Println(carro1.GetInfo())
}