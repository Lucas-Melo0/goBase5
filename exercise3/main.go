package main

import (
	"errors"
	"fmt"
)

func CheckTax(salary int) error {
	if salary <= 15000 {
		return fmt.Errorf("O salário digitado: %v não está dentro do valor mínimo", salary)
	} else {
		return errors.New("Necessário pagamento de imposto")
	}
}

func main() {
	salary := 15000
	err := CheckTax(salary)
	if err != nil {
		fmt.Println(err)
	}
}
