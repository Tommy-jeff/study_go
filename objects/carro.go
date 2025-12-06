package objects

import "fmt"

// Carro struct represents a car with brand, model, and year.
type Carro struct {
	Marca  string
	Modelo string
	Ano    int
}

// NewCar returns a new Carro struct with the specified brand, model, and year.
func NewCar(marca string, modelo string, ano int) Carro {
	return Carro {
		Marca:  marca,
		Modelo: modelo,
		Ano:    ano,
	}
}

// GetInfo returns a formatted string containing the car's information.
func (c *Carro) GetInfo() string {
	c.Ano = 2050

	return fmt.Sprintf("Carro Marca: %s, Modelo: %s, Ano: %d", c.Marca, c.Modelo, c.Ano)
}