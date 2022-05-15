// Programa em Golang, para ilustrar o conceito
// de parsing json para um array.

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

// funcao principal
func main() {

	//definindo a estrutura da instancia
	var pais []Pais

	//Array em json para ser decoded
	// para um array in golang
	Dados := []byte(`
	[
		{"Nome": "Japao", "Capital": "Tokio", "Continente": "Asia"},
		{"Nome": "Brasil", "Capital": "Brasilia", "Continente": "AmericaSul"},
		{"Nome": "Canada", "Capital": "Ottawa", "Continente": "AmericaNorte"}

	]`)

	//Decoding Array Json para
	// o Array Pais
	err := json.Unmarshal(Dados, &pais)

	if err != nil {

		//Se error não for nulo
		// print error
		fmt.Println(err)

	}

	//print o array decoded
	//valores um por um
	for i := range pais {
		fmt.Println(pais[i].Nome + " - " + pais[i].Capital + " - " + pais[i].Continente)
	}
}
