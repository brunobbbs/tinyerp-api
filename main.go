package main

import (
	"fmt"
	"log"
)

func main() {
	t := NewTinyERP()
	c, err := t.getContact("448064541")
	if err != nil {
		log.Printf("Erro ao obter o contato. %v", err)
	}
	fmt.Println("Dados do contato: ")
	fmt.Println("ID: ", c.Response.Contact.ID)
	fmt.Println("Nome: ", c.Response.Contact.Name)
	fmt.Println("Tipo: ", c.Response.Contact.Type)
	fmt.Println("CPF/CNPJ: ", c.Response.Contact.CPFCNPJ)
}
