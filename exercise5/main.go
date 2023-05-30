package main

import (
	"fmt"
	"os"
)

func main() {

	customers, err := os.Open("customers.txt")
	if err != nil {
		panic("o arquivo indicado não foi encontrado ou está danificado")
	}

	data, err := os.ReadFile(customers.Name())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

}
