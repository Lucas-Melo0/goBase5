package main

import (
	"errors"
	"fmt"
)

func SalaryCalculator(hoursWorked int, valuePerHour int) (float64, error) {

	monthlyWage := float64(hoursWorked * valuePerHour)
	if hoursWorked < 80 {
		err := errors.New("erro: o trabalhador não pode ter trabalhado menos de 80 horas por mês")
		return 0, err
	}

	if monthlyWage >= 20000 {
		monthlyWage = monthlyWage * 0.90
	}
	return monthlyWage, nil

}

func main() {
	salary, err := SalaryCalculator(60, 150)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The calculated salary is %.2f\n", salary)
	}

}
