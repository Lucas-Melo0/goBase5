package main

import "fmt"

type error interface {
	Error() string
}
type ErrorGenerator struct{}

func (e ErrorGenerator) Error(salary int) string {
	err := ""
	if salary <= 15000 {
		err = fmt.Sprintln("O salário digitado não está dentro do valor mínimo")
	} else {
		err = fmt.Sprintln("Necessário pagamento de imposto")
	}

	return err

}

func main() {
	salary := 150000
	salaryError := ErrorGenerator{}
	msg := salaryError.Error(salary)
	fmt.Println(msg)
}
