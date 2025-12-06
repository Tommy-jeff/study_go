package objects

import (
	"encoding/json"
	"fmt"
	"log"
)


// Cliente struct represents a client with name, age, email, and CPF.
// the struct fields have JSON tags for serialization to change the field names in JSON output
type Cliente struct {
	Nome     string `json:"nome"`
	Idade    int	`json:"idade"`
	Email    string `json:"email"`
	Cpf 	 string `json:"cpf"`
}

type ClienteInter struct {
	Cliente
	Pais string `json:"pais"`
}

func NewCliente(nome string, idade int, email string, cpf string) Cliente {
	return Cliente{
		Nome:  nome,
		Idade: idade,
		Email: email,
		Cpf:   cpf,
	}
}

func NewEmptyCliente() Cliente {
	return Cliente{}
}

func NewClienteInter(nome string, idade int, email string, cpf string, pais string) ClienteInter {
	return ClienteInter{
		Cliente: NewCliente(nome, idade, email, cpf),
		Pais:    pais,
	}
}

func NewEmptyClienteInter() ClienteInter {
	return ClienteInter{}
}

// this method can be called on Cliente struct instances and on ClienteInter struct instances, because ClienteInter embeds Cliente
func (c *Cliente) GetInfo() string {
	return fmt.Sprintf("Nome: %s\nIdade: %d\nEmail: %s\nCPF: %s", c.Nome, c.Idade, c.Email, c.Cpf)
}

func (c *Cliente) ToJson() (ClienteJson []byte) {
	ClienteJson, err := json.Marshal(c)
	if err != nil {
		log.Fatal("Error converting to JSON:", err)
		return
	}
	return
}

func (c *Cliente) FromJson(data []byte) {
	err := json.Unmarshal(data, c)
	if err != nil {
		log.Fatal("Error converting from JSON:", err)
		return
	}
}


