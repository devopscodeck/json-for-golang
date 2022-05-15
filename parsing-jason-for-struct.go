// Programa em Golang para ilustrar
// o conceito de parsing json para struct
package main

import (
	"encoding/json"
	"fmt"
)

// Definindo a estrutura
type Pais struct {

	// definindo as variaveis da estrutura
	Nome       string
	Capital    string
	Continente string
}

// main function
func main() {

	// Definindo a instancia da estrutura
	var pais1 Pais

	// dado em formato JSON que devem
	//ser decodificados
	Dados := []byte(`{
		"Nome": "India",
		"Capital": "New Delhi",
		"Continente": "Asia"
	}`)

	// decodificando estrutura pais1
	// a partir de formato json
	err := json.Unmarshal(Dados, &pais1)

	if err != nil {

		// Se erro não for nil
		// print error
		fmt.Println(err)
	}

	// Printando detalhes do
	// dado decodificado
	fmt.Println("A estrutura eh:", pais1)
	fmt.Printf("O país é a %s a capital é %s e o continente %s.\n", pais1.Nome,
		pais1.Capital, pais1.Continente)
}
