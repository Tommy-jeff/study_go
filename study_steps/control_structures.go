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

func ControlStructuresComplex() {

	cliente := objects.NewCliente(
		"Thomas", 24, "email@email.com", "123.456.789-00",
	)

	fmt.Println(cliente.GetInfo())


	clienteInter := objects.NewClienteInter(
		"Anna", 30, "email@email.com", "123.456.789-00", "Brazil",
	)

	fmt.Println(clienteInter.GetInfo())
	/// --------------------------------------------------------------------------------------------------------

	clienteJson := cliente.ToJson()
	fmt.Println("\nCliente in JSON format:", string(clienteJson))

	clienteInterJson := clienteInter.ToJson()
	fmt.Println("\nClienteInter in JSON format:", string(clienteInterJson))
	
	/// --------------------------------------------------------------------------------------------------------
	simulatedReceivedJson := `{"nome":"Jeff","idade":35,"email":"email@email.com","cpf":"123.456.789-00"}`

	clienteFromJson := objects.NewEmptyCliente()
	clienteFromJson.FromJson([]byte(simulatedReceivedJson))
	fmt.Println("\nCliente from JSON:\n", clienteFromJson.GetInfo())
	

}