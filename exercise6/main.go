package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Customer struct {
	File      string
	Name      string
	Surname   string
	CPF       string
	Cellphone string
	Address   string
}

func Register(name, surname, cpf, cellphone, address string) Customer {
	newCustomer := Customer{Name: name, Surname: surname, CPF: cpf, Cellphone: cellphone, Address: address}
	return newCustomer
}

func IdGeneration(cpf string) string {
	rand.Seed(time.Now().UnixNano())
	newId := fmt.Sprintf("%s%d%d", cpf, time.Now().UnixNano(), rand.Intn(1000))
	if newId == "" {
		panic("invalid value for id")
	}
	return newId
}
func (c *Customer) SetFileName(id string) {
	c.File = fmt.Sprintf("%s.txt", id)
}

func CheckCustomerExistence(cpf string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("erro: o arquivo indicado não foi encontrado ou está danificado")
		}
	}()

	_, err := os.ReadFile("customers.txt")
	if err != nil {
		panic(err)
	}
}

func main() {
	CheckCustomerExistence("123")
	lucas := Register("Lucas", "Melo", "123", "8399210201", "Rua golfo")
	id := IdGeneration(lucas.CPF)
	lucas.SetFileName(id)
	fmt.Println(lucas)

}
