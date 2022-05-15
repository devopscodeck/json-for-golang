// Programa em Golang para ilustrar
// o conceito de formatar o texto para JSON
package main

import (
	"encoding/json"
	"fmt"
)

// declarando a estrutura
type Humano struct {

	// definindo as variaveis da estrutura
	Nome     string
	Idade    int
	Endereco string
}

// main function
func main() {

	// definindo a estrutura da instancia
	humano1 := Humano{"Julio", 23, "Coronel Leitao"}

	// encoding estrutura humano1
	// para o formato json
	humano_enc, err := json.Marshal(humano1)

	if err != nil {

		// se o errro não for nil
		// print error
		fmt.Println(err)
	}
	// como o humano_enc eh um byte de formato array
	// é preciso ser convertido para uma string, conforme abaixo
	fmt.Println(string(humano_enc))
	//
	//
	//
	// convertendo slices de
	// golang para formato JSON

	// definindo um array
	// de estrutura da instancia
	humano2 := []Humano{
		{Nome: "Rahul", Idade: 23, Endereco: "New Delhi"},
		{Nome: "Priyanshi", Idade: 20, Endereco: "Pune"},
		{Nome: "Shivam", Idade: 24, Endereco: "Bangalore"},
	}

	// codificando para o formato JSON
	humano2_enc, err := json.Marshal(humano2)

	if err != nil {

		// Se erro não nulo
		// print error
		fmt.Println(err)
	}

	// printando o array codificado em json
	fmt.Println()
	fmt.Println(string(humano2_enc))
}
